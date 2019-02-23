package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ApiDb struct {
	*dynamodb.DynamoDB
}

type EnvSingleton struct {
	Db		*ApiDb
	tableName	string
}
