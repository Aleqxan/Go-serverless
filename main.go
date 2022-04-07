package main

import (
	"os"

	"github.com/Aleqxan/go-serverless-yt/pkg/handlers"
	"github.com/aws-sdk-go/aws"
	"github.com/aws-sdk-go/service/dynamodb/dynamobiface"
	"github.com/aws/aws-lambda-go"
	"github.com/aws/aws-lamda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)
var(
	dynaClient dynamodbiface.DynamoDBAPI
)
func main(){
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&awsSession{
		Region: aws.String(region)},)
	
	if err!= nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyRequest, error){
		switch req.HTTPMethod{
		case "GET":
			return handlers.GetUser(req, tableName, dynaClient)
		case "POST":
			return handlers.CreateUser(req, tableName, dynaClient)
		case "PUT":
			return handlers.UpdateUser(req, tableName, dynaClient)
		case "DELETE":
			return handlers.DeleteUser(req, tableName, dynaClient)
		}
		default:
			return handlers.UnhandleMethod()
}