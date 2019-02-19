package api

import (
	"fmt"
	"encoding/json"
	"net/http"

	//"github.com/ttacon/libphonenumber"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Product struct {
        Name  string  `json:"Name"`
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
        fmt.Println("healthcheck api")

        prod := &Product{Name: "chirag"}

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(prod); err != nil {

                fmt.Printf("/HostRegistrationStatus Internal Error: %s", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

type Address struct {
	City		string `json:"city"`
	State		string `json:"state"`
	Country		string `json:"country"`
	PinCode		string `json:"pincode"`
	PoBox		string `json:"pobox, omitempty"`
	StreetAddress	string `json:"streetAddress"`
}

type User struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Phone		string		`json:"phone"`
	Address		Address		`json:"address"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Got error marshalling map1:")
		fmt.Println(err.Error())
	}

	fmt.Printf("User: %v", user)

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item: av ,
		TableName: aws.String("Users"),
	}

	fmt.Printf("\nUser3: %v", input)

	fmt.Println(GetEnvInstance().Db)

	_, err = GetEnvInstance().Db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully added 'The Big New Movie' (2015) to Movies table")
}
