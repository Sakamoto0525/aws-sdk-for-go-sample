# CognitoでClientCredentialsGrantを利用する

参考：https://qiita.com/poruruba/items/3681f1e1f8e72fcaadfa#api-gateway%E3%81%A7%E5%91%BC%E3%81%B3%E5%87%BA%E3%81%97%E5%85%83%E8%AA%8D%E8%A8%BC%E3%81%AB%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B%E5%AE%9F%E8%A1%8C%E7%B7%A8


## 動作確認

### 1. アクセストークンを取得

Cognito アプリクライアントIDとアプリクライアントシークレットを組み合わせてbase64エンコードした文字列を使用します。

文字列は下記コマンドで生成できます。

```bash
echo -n "アプリクライアントID:アプリクライアントのシークレット" | base64
```

次にCognitoコンソール画面のCognitoドメインを見ます

POST `Cognitoドメイン名`/oauth2/token に対してアクセストークンをリクエストします

下記が設定するキーと値です。

> Header

|キー|値|
| --- | --- |
|Authorization|Basic 文字列|

> Body

|キー|値|
| --- | --- |
|grant_type|client_credentials|
|scope|client-credentials-grant-sample-resource-server/users.read|

### 2. Lambda APIを叩く

1で取得したアクセストークンを使いリクエストを送信する

> Header

|キー|値|
| --- | --- |
|Authorization|アクセストークン|

`GET https://xxxxxxxx.execute-api.ap-northeast-1.amazonaws.com/dev/hello`
