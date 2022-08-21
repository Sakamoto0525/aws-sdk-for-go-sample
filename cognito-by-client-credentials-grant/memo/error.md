# error


## When configuring an authorizer with type: "COGNITO_USER_POOLS", property "id" or "name" has to be specified.

##### エラー内容

```bash
22-08-21 15:44 ~/Projects/Study/aws-sdk-for-go-sample/cognito-by-client-credentials-grant $ make deploy
rm -rf ./bin
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
sls deploy --verbose

Warning: Invalid configuration encountered
  at 'functions.hello.events.0.httpApi.authorizer': unrecognized property 'authorizerId'
  at 'functions.hello.events.0.httpApi.authorizer.type': must be equal to one of the allowed values [request, jwt, aws_iam]

Learn more about configuration validation here: http://slss.io/configuration-validation

Deploying cognito-by-client-credentials-grant to stage dev (ap-northeast-1)

Packaging
Excluding development dependencies for service package

✖ Stack cognito-by-client-credentials-grant-dev failed to deploy (0s)
Environment: darwin, node 14.17.0, framework 3.21.0, plugin 6.2.2, SDK 4.3.2
Credentials: Local, "default" profile
Docs:        docs.serverless.com
Support:     forum.serverless.com
Bugs:        github.com/serverless/serverless/issues

Error:
When configuring an authorizer with type: "COGNITO_USER_POOLS", property "id" or "name" has to be specified.

1 deprecation found: run 'serverless doctor' for more details
make: *** [deploy] Error 1
```

##### 原因

##### 対処

serverles.ymlのfunctinos.関数名.eventsで定義しているhttpApiをhttpに変更したらいけた

```yml
functions:
  hello:
    handler: bin/hello
    events:
      # - httpApi:
      - http: # - http: に変更
          path: /hello
          method: get
          authorizer:
            type: COGNITO_USER_POOLS
            authorizerId: !Ref ApiGatewayWithAuthorizationAuthorizer
            scopes:
              - client-credentials-grant-sample-resource-server/users.read
```

## 1 validation error detected: Value 'cognito-by-client-credentials-grant-dev-ap-northeast-1-lambdaRole' at 'roleName' failed to satisfy constraint: Member must have length less than or equal to 64

##### エラー内容

```bash
Error:
CREATE_FAILED: IamRoleLambdaExecution (AWS::IAM::Role)
1 validation error detected: Value 'cognito-by-client-credentials-grant-dev-ap-northeast-1-lambdaRole' at 'roleName' failed to satisfy constraint: Member must have length less than or equal to 64 (Service: AmazonIdentityManagement; Status Code: 400; Error Code: ValidationError; Request ID: 704dc0a9-9bd4-4a01-8aa3-da28bc6d7158; Proxy: null)
```

##### 原因

'roleName'の名前の長さが64文字以内でないといけない

どうやら、IAMロール名は`{プロジェクト名}-{ステージ名}-{リージョン名}-lambdaRole`で決まるぽい
単純にプロジェクト名が長すぎたせい

##### 対処

プロジェクト名を略称にしました

## エラー

##### エラー内容
##### 原因
##### 対処
