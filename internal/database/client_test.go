package database

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/danielMensah/onetable-go/internal/database/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sampleItem struct {
	ID   string `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
}

func TestClient_CreateItem(t *testing.T) {
	tests := []struct {
		name string
		item sampleItem
		err  error
	}{
		{
			name: "successfully create item",
			item: sampleItem{
				ID:   "123",
				Name: "abigail",
			},
			err: nil,
		},
		{
			name: "failed to create item",
			item: sampleItem{
				ID:   "123",
				Name: "abigail",
			},
			err: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			m := mocks.NewDynamoDBClient(t)
			c := &Client{
				db:        m,
				tableName: aws.String("onetable"),
			}

			m.EXPECT().PutItem(ctx, &dynamodb.PutItemInput{
				TableName: aws.String("onetable"),
				Item: map[string]types.AttributeValue{
					"id":   &types.AttributeValueMemberS{Value: tt.item.ID},
					"name": &types.AttributeValueMemberS{Value: tt.item.Name},
				},
			}).Return(nil, tt.err)

			_, err := c.CreateItem(ctx, tt.item)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClient_Find(t *testing.T) {
	m := mocks.NewDynamoDBClient(t)
	ctx := context.TODO()

	tests := []struct {
		name         string
		query        string
		args         []interface{}
		out          interface{}
		mockSetup    func()
		err          error
		expectedOut  interface{}
		expectedCall bool
	}{
		{
			name:  "successfully finds a list of items",
			query: "id = $1",
			args:  []interface{}{"123"},
			out:   &[]sampleItem{},
			expectedOut: []sampleItem{
				{ID: "123", Name: "John Doe"},
				{ID: "123", Name: "Jane Doe"},
			},
			mockSetup: func() {
				m.On("Query", ctx, &dynamodb.QueryInput{
					TableName:              aws.String("testTable"),
					KeyConditionExpression: aws.String("id = :v1"),
					ExpressionAttributeValues: map[string]types.AttributeValue{
						":v1": &types.AttributeValueMemberS{Value: "123"},
					},
				}).Return(&dynamodb.QueryOutput{
					Items: []map[string]types.AttributeValue{
						{"id": &types.AttributeValueMemberS{Value: "123"}, "name": &types.AttributeValueMemberS{Value: "John Doe"}},
						{"id": &types.AttributeValueMemberS{Value: "123"}, "name": &types.AttributeValueMemberS{Value: "Jane Doe"}},
					},
				}, nil).Once()
			},
			expectedCall: true,
		},
		{
			name:        "successfully finds an item",
			query:       "id = $1",
			args:        []interface{}{"123"},
			out:         &sampleItem{},
			expectedOut: sampleItem{ID: "123", Name: "John Doe"},
			mockSetup: func() {
				m.On("GetItem", ctx, &dynamodb.GetItemInput{
					TableName: aws.String("testTable"),
					Key:       map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: "123"}},
				}).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"id":   &types.AttributeValueMemberS{Value: "123"},
						"name": &types.AttributeValueMemberS{Value: "John Doe"},
					},
				}, nil).Once()
			},
			expectedCall: true,
		},
		{
			name:  "fails when output is not a pointer",
			query: "id = $1",
			args:  []interface{}{"123"},
			out:   []sampleItem{}, // Not a pointer
			mockSetup: func() {
				// No mock setup needed as error occurs before Query/GetItem is called
			},
			err:          ErrParameterNotPointer,
			expectedCall: false,
		},
		{
			name:  "fails when query returns an error",
			query: "id = $1",
			args:  []interface{}{"123"},
			out:   &[]sampleItem{},
			mockSetup: func() {
				m.On("Query", ctx, &dynamodb.QueryInput{
					TableName:              aws.String("testTable"),
					KeyConditionExpression: aws.String("id = :v1"),
					ExpressionAttributeValues: map[string]types.AttributeValue{
						":v1": &types.AttributeValueMemberS{Value: "123"},
					},
				}).Return(nil, errors.New("query failed")).Once()
			},
			err:          ErrQuery,
			expectedCall: true,
		},
		{
			name:        "handles empty query results gracefully",
			query:       "id = $1",
			args:        []interface{}{"999"}, // No matching items
			out:         &[]sampleItem{},
			expectedOut: []sampleItem{}, // Expect empty slice
			mockSetup: func() {
				m.On("Query", ctx, &dynamodb.QueryInput{
					TableName:              aws.String("testTable"),
					KeyConditionExpression: aws.String("id = :v1"),
					ExpressionAttributeValues: map[string]types.AttributeValue{
						":v1": &types.AttributeValueMemberS{Value: "999"},
					},
				}).Return(&dynamodb.QueryOutput{Items: []map[string]types.AttributeValue{}}, nil).Once()
			},
			expectedCall: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			c := &Client{
				db:        m,
				tableName: aws.String("testTable"),
			}

			err := c.Find(ctx, tt.out, tt.query, tt.args...)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)

				outType := reflect.TypeOf(tt.out).Elem().Kind()
				if outType == reflect.Slice {
					actual, ok := tt.out.(*[]sampleItem)
					require.True(t, ok)

					assert.Equal(t, tt.expectedOut, *actual)
				} else {
					actual, ok := tt.out.(*sampleItem)
					require.True(t, ok)

					assert.Equal(t, tt.expectedOut, *actual)
				}
			}

			if tt.expectedCall {
				m.AssertExpectations(t)
			}
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	tests := []struct {
		name         string
		key          map[string]string
		updates      map[string]interface{}
		mockSetup    func()
		err          error
		expectedCall bool
	}{
		{
			name: "successfully updates an item",
			key:  map[string]string{"id": "123"},
			updates: map[string]interface{}{
				"name": "John Doe Updated",
			},
			mockSetup: func() {
				m := mocks.NewDynamoDBClient(t)
				updateExpression := "SET name = :name"
				expressionAttributeValues := map[string]types.AttributeValue{
					":name": &types.AttributeValueMemberS{Value: "John Doe Updated"},
				}

				m.On("UpdateItem", context.TODO(), &dynamodb.UpdateItemInput{
					TableName:                 aws.String("onetable"),
					Key:                       map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: "123"}},
					UpdateExpression:          &updateExpression,
					ExpressionAttributeValues: expressionAttributeValues,
				}).Return(&dynamodb.UpdateItemOutput{}, nil).Once()
			},
			err:          nil,
			expectedCall: true,
		},
		{
			name: "fails when UpdateItem returns an error",
			key:  map[string]string{"id": "123"},
			updates: map[string]interface{}{
				"name": "John Doe Updated",
			},
			mockSetup: func() {
				m := mocks.NewDynamoDBClient(t)
				updateExpression := "SET name = :name"
				expressionAttributeValues := map[string]types.AttributeValue{
					":name": &types.AttributeValueMemberS{Value: "John Doe Updated"},
				}

				m.On("UpdateItem", context.TODO(), &dynamodb.UpdateItemInput{
					TableName:                 aws.String("onetable"),
					Key:                       map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: "123"}},
					UpdateExpression:          &updateExpression,
					ExpressionAttributeValues: expressionAttributeValues,
				}).Return(nil, errors.New("update failed")).Once()
			},
			err:          errors.New("update failed"),
			expectedCall: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			m := mocks.NewDynamoDBClient(t)

			c := &Client{
				db:        m,
				tableName: aws.String("onetable"),
			}

			tt.mockSetup()

			err := c.UpdateItem(ctx, tt.key, tt.updates)
			if tt.err != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			if tt.expectedCall {
				m.AssertExpectations(t)
			}
		})
	}
}
