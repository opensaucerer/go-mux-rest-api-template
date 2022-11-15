package service

import (
	"log"

	config "github.com/opensaucerer/gotemplate/config"

	aws "github.com/aws/aws-sdk-go/aws"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	dynamo "github.com/guregu/dynamo"
)

// exposing as global connection variables
var (
	Dynamo    *dynamodb.DynamoDB
	Session   *awssession.Session
	DB        *dynamo.DB
	Txtable   dynamo.Table
	CompTable dynamo.Table
)

func MustInitiateDynamo() {

	// config := &aws.Config{
	// 	Region:   aws.String("ap-south-1"),
	// 	Endpoint: aws.String("http://localhost:" + config.MustGet("PORT", "8000")),
	// }

	// session = awssession.Must(awssession.NewSession(config))
	Session = awssession.Must(awssession.NewSessionWithOptions(awssession.Options{SharedConfigState: awssession.SharedConfigEnable}))

	Dynamo = dynamodb.New(Session)

	DB = dynamo.New(Session, &aws.Config{Region: aws.String(config.MustGet("AWS_REGION", "ap-south-1"))})
	Txtable = DB.Table(config.COOL_TABLE_NAME)

	log.Println("initiated connection to DynamoDB...")
}
