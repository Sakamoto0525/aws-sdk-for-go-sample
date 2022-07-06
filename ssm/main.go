package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	svc := ssm.New(sess)

	res, err := svc.GetParameters(&ssm.GetParametersInput{
		Names:          []*string{aws.String("/name")},
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		panic(err)
	}

	result := map[string]string{}
	for _, p := range res.Parameters {
		switch *p.Name {
		case "/name":
			result["Name"] = *p.Value
		}
	}

	fmt.Println(result)
}
