// Code generated by mockery v2.50.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// CreateItem provides a mock function with given fields: ctx, in
func (_m *Client) CreateItem(ctx context.Context, in interface{}) (*dynamodb.PutItemOutput, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for CreateItem")
	}

	var r0 *dynamodb.PutItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*dynamodb.PutItemOutput, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *dynamodb.PutItemOutput); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CreateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateItem'
type Client_CreateItem_Call struct {
	*mock.Call
}

// CreateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - in interface{}
func (_e *Client_Expecter) CreateItem(ctx interface{}, in interface{}) *Client_CreateItem_Call {
	return &Client_CreateItem_Call{Call: _e.mock.On("CreateItem", ctx, in)}
}

func (_c *Client_CreateItem_Call) Run(run func(ctx context.Context, in interface{})) *Client_CreateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Client_CreateItem_Call) Return(_a0 *dynamodb.PutItemOutput, _a1 error) *Client_CreateItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CreateItem_Call) RunAndReturn(run func(context.Context, interface{}) (*dynamodb.PutItemOutput, error)) *Client_CreateItem_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx, out, query, args
func (_m *Client) Find(ctx context.Context, out interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, out, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, out, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type Client_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
//   - out interface{}
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Find(ctx interface{}, out interface{}, query interface{}, args ...interface{}) *Client_Find_Call {
	return &Client_Find_Call{Call: _e.mock.On("Find",
		append([]interface{}{ctx, out, query}, args...)...)}
}

func (_c *Client_Find_Call) Run(run func(ctx context.Context, out interface{}, query string, args ...interface{})) *Client_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Find_Call) Return(_a0 error) *Client_Find_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Find_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *Client_Find_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItem provides a mock function with given fields: ctx, key, updates
func (_m *Client) UpdateItem(ctx context.Context, key interface{}, updates map[string]interface{}) error {
	ret := _m.Called(ctx, key, updates)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, map[string]interface{}) error); ok {
		r0 = rf(ctx, key, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_UpdateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItem'
type Client_UpdateItem_Call struct {
	*mock.Call
}

// UpdateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - key interface{}
//   - updates map[string]interface{}
func (_e *Client_Expecter) UpdateItem(ctx interface{}, key interface{}, updates interface{}) *Client_UpdateItem_Call {
	return &Client_UpdateItem_Call{Call: _e.mock.On("UpdateItem", ctx, key, updates)}
}

func (_c *Client_UpdateItem_Call) Run(run func(ctx context.Context, key interface{}, updates map[string]interface{})) *Client_UpdateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *Client_UpdateItem_Call) Return(_a0 error) *Client_UpdateItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_UpdateItem_Call) RunAndReturn(run func(context.Context, interface{}, map[string]interface{}) error) *Client_UpdateItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
