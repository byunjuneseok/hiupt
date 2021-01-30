package main

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

func Retrieve(id string) (User, error) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	user := User{}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME_USER")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return User{}, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}


//func List() ([]User, error) {
//
//	filter := expression.Name("id").Equal(expression.Value(id))
//	projection := expression.NamesList(expression.Name("id"), expression.Name("email"))
//
//	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
//
//	if err != nil {
//		return user, err
//	}
//
//}

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

