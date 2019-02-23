package handlers

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type EnvSingleton struct {
	Db		*dynamodb.DynamoDB
	Kms		*kms.KMS
	tableName	string
}
