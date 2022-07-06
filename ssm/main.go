package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	// AWSサービスと接続するための共通設定。利用するリージョンやAWSでの認証方法を設定可能。
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		// Profile: "default",
	}))

	// ssm (AWS System Manager)を利用するための設定
	svc := ssm.New(sess)

	// ストアパラメータからNamesに指定したキーと一致した値を全て取得する
	res, err := svc.GetParameters(&ssm.GetParametersInput{
		Names: []*string{aws.String("/name")},

		// trueの場合はパラメータが暗号化している場合に複合化して取得する
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
