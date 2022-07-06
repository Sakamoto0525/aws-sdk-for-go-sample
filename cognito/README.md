# CognitoでIDトークンを取得する

コンソール上から既にUserPoolにユーザー情報を作成している前提

> サンプルコードのレスポンス

```bash
{
  AuthenticationResult: {
    AccessToken: <sensitive>,
    ExpiresIn: 3600,
    IdToken: <sensitive>,
    RefreshToken: <sensitive>,
    TokenType: "Bearer"
  },
  ChallengeParameters: {

  }
}
```

### 参考記事

[
Go 言語と AWS SDK V2 で Amazon Cognito を操作する](https://maku.blog/p/nej9wjb/)