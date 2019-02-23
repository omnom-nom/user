package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/ttacon/libphonenumber"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Address struct {
	City		string `json:"city"`
	State		string `json:"state"`
	Country		string `json:"country"`
	PinCode		string `json:"pincode"`
	PoBox		string `json:"pobox, omitempty"`
	StreetAddress	string `json:"streetAddress"`
}

type UserRequest struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Phone		string		`json:"phone"`
	Address		Address		`json:"address"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

	user := UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorf("Error decoding user input: %s", err.Error())
		writeHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if ok, err := validatePhone(user.Phone, user.Address.Country); !ok {
		writeHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Errorf("Error marshaling to dynamoDb attribute: %s", err.Error())
		writeHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(GetEnvInstance().tableName),
	}

	_, err = GetEnvInstance().Db.PutItem(input)
	if err != nil {
		log.Errorf("Error putting item into table: %s", err.Error())
		writeHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	log.Infof("Success inserting item into db table")
}


func validatePhone(phone, country string) (bool, error) {
	log.Infof("Validating phone: %s", phone)

	num, err := libphonenumber.Parse(phone, country)
	if err != nil {
		log.Errorf("Error: %s parsing phone: %s for country: %s", err.Error(), phone, country)
		return false, err
	}
	if !libphonenumber.IsValidNumber(num) {
		err := fmt.Errorf("Error: %s validating phone: %s for country: %s", err.Error(), phone, country)
		log.Error(err.Error())
		return false, err
	}

	return true, nil
}
