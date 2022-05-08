# invoke-async-aws-lambda-example

AWS Lambda関数から別の関数を非同期に呼び出すサンプルプログラムです。

```
rm invoke-async-aws-lambda-example.zip || true
GOOS=linux go build main.go
zip invoke-async-aws-lambda-example.zip main
rm main
```

で`invoke-async-aws-lambda-example.zip`が作成されるので、それをLambdaにアップロードすると動きます。
ただし、ハンドラを`main`に変更し、Lambda関数の環境変数に以下を設定してください。

- INVOKE_FUNCTION_NAME: 呼び出すLambda関数の名前

また、Lamda関数のIAMロールに以下のポリシーを付与してください。

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "lambda:InvokeAsync",
        "lambda:InvokeFunction"
      ],
      "Resource": "arn:aws:lambda:*:*:function:*",
      "Effect": "Allow"
    }
  ]
}
```
