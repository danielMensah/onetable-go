package onetable

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/danielMensah/onetable-go/internal/database"
)

type Client interface {
	// CreateItem creates a new item in the table.
	CreateItem(ctx context.Context, in interface{}) (*dynamodb.PutItemOutput, error)
	// Find can retrieve a single item or a list of items from the table depending on the out parameter.
	// If out is a slice, Find will retrieve multiple items. If out is not a slice, Find will retrieve a single item.
	Find(ctx context.Context, out interface{}, query string, args ...interface{}) error
}

// New creates a new instance of OneTable Client.
func New(dynamoDBClient *dynamodb.Client, tableName string) Client {
	return database.NewClient(dynamoDBClient, aws.String(tableName))
}