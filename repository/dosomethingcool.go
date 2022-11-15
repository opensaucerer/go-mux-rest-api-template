package repository

import (
	service "github.com/opensaucerer/gotemplate/service"
	types "github.com/opensaucerer/gotemplate/typing"

	// awserr "github.com/aws/aws-sdk-go/aws/awserr"

	// dynamodbattribute "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	dynamo "github.com/guregu/dynamo"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func DoSomethingCool(c types.Commitment) error {
	err := service.Txtable.Put(dynamo.AWSEncoding(c)).Run()
	if err != nil {
		return err
	}
	return nil
}
