package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	poolID       = "xxxxxxxx"
	clientID     = "xxxxxxxx"
	clientSecret = "xxxxxxxx"
)

// IDトークンを取得する
func getIDToken() {
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

var (
	username = "xxxxxxxx"
	password = "xxxxxxxx"
)

// ユーザーを作成する
func createUser(sess *session.Session) {
	input := cognitoidentityprovider.AdminCreateUserInput{
		// ユーザープールID
		UserPoolId: aws.String(poolID),
		// ユーザーネーム
		// ・必須
		// ・1文字以上、128文字以下
		// ・UTF-8文字列
		// ・ユニークである為、変更不可
		Username: aws.String(username),
		// ユーザーの一時的なパスワード
		// ・必須ではなく、未設定の場合はCognitoがよしなに生成してくれる
		// ・初回サインイン後にこの仮パスワードと新規パスワードを設定する
		TemporaryPassword: aws.String(password),
	}
	if err := input.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := cognitoidentityprovider.New(sess)
	if _, err := c.AdminCreateUser(&input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// アプリクライアントシークレットをHash化する
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(*&username + clientID))
	secretHash := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	initiateAuthOutput, err := c.AdminInitiateAuth(&cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String("ADMIN_USER_PASSWORD_AUTH"),
		UserPoolId: &poolID,
		ClientId:   &clientID,
		AuthParameters: map[string]*string{
			"USERNAME":    &username,
			"PASSWORD":    &password,
			"SECRET_HASH": &secretHash,
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respondToAuthChallengeOutput, err := c.AdminRespondToAuthChallenge(&cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		UserPoolId:    &poolID,
		ClientId:      &clientID,
		ChallengeName: aws.String("NEW_PASSWORD_REQUIRED"),
		ChallengeResponses: map[string]*string{
			"USERNAME":     &username,
			"NEW_PASSWORD": &password,
			"SECRET_HASH":  &secretHash,
		},
		Session: initiateAuthOutput.Session,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*respondToAuthChallengeOutput.AuthenticationResult.AccessToken)

}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	createUser(sess)
}
