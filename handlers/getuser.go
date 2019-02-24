package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/kms"

	"github.com/mitchellh/mapstructure"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

	result, err := GetEnvInstance().Db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(GetEnvInstance().tableName),
		Key: map[string]*dynamodb.AttributeValue{
		    "email": {
			S: aws.String("chiragtayal@gmail.com"),
		    },
		    "phone": {
			S: aws.String("7404175904"),
		    },
		},
	})

	if err != nil {
		log.Errorf("Error getting dynamoDb attribute: %s", err.Error())
                writeHTTPError(w, http.StatusInternalServerError, err)
                return
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Errorf("Error marshaling to dynamoDb attribute: %s", err.Error())
		writeHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	user := UserRequest{
		Email: item.Email,
		Phone: item.Phone,
		Name: item.Name,
		UserType: item.UserType,
	}
	addressMap := decryptData(item.Address)
	mapstructure.Decode(addressMap, &user.Address)

	log.Debugf("User: %v", user)

	writeHTTPSuccess(w, http.StatusOK, user)
}

func decryptData(blob []byte) interface{} {

	result, err := GetEnvInstance().Kms.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})
	if err != nil {
		log.Errorf("Unable to decrypt blob: %v", err)
		return nil
	}

	var data interface{}
	err = json.Unmarshal(result.Plaintext, &data)
        if err != nil {
                log.Errorf("Response data conversion to json failed: %s", err)
                return nil
        }

	return data
}
