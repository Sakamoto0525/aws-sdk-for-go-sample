package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	authInputParams := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId: aws.String("xxxxxxxxxxxxxxxxxxx"), // アプリクライアントID
		// アプリクライアント詳細における認証フロー設定で`ALLOW_ADMIN_USER_PASSWORD_AUTH`に
		// チェックが入っていないとエラーになる
		AuthFlow:   aws.String("ADMIN_USER_PASSWORD_AUTH"),
		UserPoolId: aws.String("ap-northeast-1_xxxxxxxx"), // プールID
		// コンソール上から作成したユーザープールのユーザー情報
		AuthParameters: map[string]*string{
			"USERNAME": aws.String("username"),
			"PASSWORD": aws.String("password"),
		},
	}

	mySession := session.Must(session.NewSession())
	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	req, resp := svc.AdminInitiateAuthRequest(authInputParams)
	err := req.Send()
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}
}
