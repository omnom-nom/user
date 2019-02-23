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
	"github.com/aws/aws-sdk-go/service/kms"
)

type Address struct {
	City		string `json:"city"`
	State		string `json:"state"`
	Country		string `json:"country"`
	PinCode		string `json:"pincode"`
	PoBox		string `json:"pobox, omitempty"`
	StreetAddress	string `json:"streetAddress"`
}

type Payment struct {
	CardAlias	string	`json:"cardalias, omitempty"`
	NameOnCard	string	`json:"nameOnCard"`
	AddressOnCard	Address	`json:"addressOnCard"`
	CardExpiry	string	`json:"cardExpiry"`
	CardCVV		string  `json:"cvv"`
}

type UserRequest struct {
	Email		string		`json:"email"`
	Phone		string		`json:"phone"`
	Name		string		`json:"name"`
	Address		Address		`json:"address"`
	Payment		Payment		`json:"payment, omitempty"`
}

type Item struct {
	Email		string	`json:"email"`
	Phone		string	`json:"phone"`
	Name		string	`json:"name"`
	Address		[]byte	`json:"address"`
	Payment		[]byte	`json:"payment"`
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

	item := Item{
		Email: user.Email,
		Phone: user.Phone,
		Name: user.Name,
		Address: encryptData(user.Address),
		Payment: encryptData(user.Payment),
	}

	av, err := dynamodbattribute.MarshalMap(item)
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

	log.Infof("Successfuly inserting item into db table")
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

func encryptData(d interface{}) []byte {

	data, err := json.Marshal(d)
        if err != nil {
                log.Errorf("Response data conversion to json failed: %s", err)
                return nil
        }

	ra, err := GetEnvInstance().Kms.Encrypt(&kms.EncryptInput{
		KeyId: aws.String("alias/db"),
		Plaintext: []byte(data),
	})

	return ra.CiphertextBlob
}
