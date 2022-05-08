package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdaApi "github.com/aws/aws-sdk-go-v2/service/lambda"
)

// API Gateway経由で実行されることを想定しているため、
// ハンドラの第二引数の型をAPIGatewayProxyRequestにしている。
// 戻り値の型はAPIGatewayProxyResponseになる。
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	requestJson, _ := json.Marshal(request)
	log.Println(string(requestJson))

	client := lambdaApi.NewFromConfig(cfg)

	// 別のLambda関数の呼び出し
	result, err := client.Invoke(ctx, &lambdaApi.InvokeInput{
		// 呼び出し先のLambda関数名(環境変数から取得)
		FunctionName: aws.String(os.Getenv("FUNCTION_NAME")),
		// Eventを指定すると非同期実行される
		InvocationType: "Event",
		// requestのペイロード(json)をそのまま呼び出し先のLambda関数に渡す
		Payload: requestJson,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Bodyを空文字列にして返した場合、Slackには何も投稿されない
	// 何か投稿したい場合はBodyに文字列を入れる
	return events.APIGatewayProxyResponse{Body: "", StatusCode: int(result.StatusCode)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
