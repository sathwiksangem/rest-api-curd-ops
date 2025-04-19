package db

import (
	"context"
	appTypes "rest-api-curd-ops/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	dynaClient *dynamodb.Client
	tableName  = "test-table"
)

func InitDynamo() error {
	cfg, err := config.LoadDefaultConfig(context.TODO()) // use WithSharedConfigProfile in options if default profile is not needed
	if err != nil {
		return err
	}
	dynaClient = dynamodb.NewFromConfig(cfg)
	return nil
}

func CreateItem(item appTypes.Item) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}
	_, err = dynaClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	return err
}

func GetItem(id string) (*appTypes.Item, error) {
	out, err := dynaClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil || out.Item == nil {
		return nil, err
	}
	var item appTypes.Item
	err = attributevalue.UnmarshalMap(out.Item, &item)
	return &item, err
}

func UpdateItem(item appTypes.Item) error {
	return CreateItem(item) // PutItem replaces existing item
}

func DeleteItem(id string) error {
	_, err := dynaClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
