package users

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func Create(body string) (User, error)  {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	var thisUser User

	_ = json.Unmarshal([]byte(body), &thisUser)

	av, err := dynamodbattribute.MarshalMap(thisUser)
	if err != nil {
		fmt.Println(err.Error())
		return thisUser, err
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(os.Getenv("TABLE_NAME_USER")),
	}

	_, err = svc.PutItem(input)

	return thisUser, err
}

//func List(body string) ([]User, error) {
//	sess := session.Must(session.NewSession())
//	svc := dynamodb.New(sess)
//
//	var users []User
//
//	result, err := svc.Scan()
//	if err != nil {
//		fmt.Println(err.Error())
//		return users, err
//	}
//
//	err = dynamodbattribute.UnmarshalListOfMaps()
//
//}