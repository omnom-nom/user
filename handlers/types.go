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

type Address struct {
        City            string `json:"city"`
        State           string `json:"state"`
        Country         string `json:"country"`
        PinCode         string `json:"pincode"`
        PoBox           string `json:"pobox, omitempty"`
        StreetAddress   string `json:"streetAddress"`
}

type Payment struct {
        CardAlias       string  `json:"cardalias, omitempty"`
        NameOnCard      string  `json:"nameOnCard"`
        AddressOnCard   Address `json:"addressOnCard"`
        CardExpiry      string  `json:"cardExpiry"`
        CardCVV         string  `json:"cvv"`
}

type UserRequest struct {
        Email		string  `json:"email"`
        Phone		string	`json:"phone"`
        Name		string  `json:"name"`
        Address		Address `json:"address"`
	UserType	string	`json:"usertype"` //customer, merchant. driver
}

type Item struct {
	Email		string	`json:"email"`
	Phone		string	`json:"phone"`
	Name		string	`json:"name"`
	Address		[]byte	`json:"address"`
	UserType	string	`json:"usertype"` //customer, merchant. driver
}
