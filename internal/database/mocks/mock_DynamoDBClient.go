// Code generated by mockery v2.50.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	mock "github.com/stretchr/testify/mock"
)

// DynamoDBClient is an autogenerated mock type for the DynamoDBClient type
type DynamoDBClient struct {
	mock.Mock
}

type DynamoDBClient_Expecter struct {
	mock *mock.Mock
}

func (_m *DynamoDBClient) EXPECT() *DynamoDBClient_Expecter {
	return &DynamoDBClient_Expecter{mock: &_m.Mock}
}

// DeleteItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDBClient) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItem")
	}

	var r0 *dynamodb.DeleteItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) *dynamodb.DeleteItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DeleteItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDBClient_DeleteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItem'
type DynamoDBClient_DeleteItem_Call struct {
	*mock.Call
}

// DeleteItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.DeleteItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDBClient_Expecter) DeleteItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDBClient_DeleteItem_Call {
	return &DynamoDBClient_DeleteItem_Call{Call: _e.mock.On("DeleteItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDBClient_DeleteItem_Call) Run(run func(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options))) *DynamoDBClient_DeleteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.DeleteItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDBClient_DeleteItem_Call) Return(_a0 *dynamodb.DeleteItemOutput, _a1 error) *DynamoDBClient_DeleteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDBClient_DeleteItem_Call) RunAndReturn(run func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)) *DynamoDBClient_DeleteItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDBClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *dynamodb.GetItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) *dynamodb.GetItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDBClient_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type DynamoDBClient_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.GetItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDBClient_Expecter) GetItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDBClient_GetItem_Call {
	return &DynamoDBClient_GetItem_Call{Call: _e.mock.On("GetItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDBClient_GetItem_Call) Run(run func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options))) *DynamoDBClient_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.GetItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDBClient_GetItem_Call) Return(_a0 *dynamodb.GetItemOutput, _a1 error) *DynamoDBClient_GetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDBClient_GetItem_Call) RunAndReturn(run func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)) *DynamoDBClient_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// PutItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDBClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutItem")
	}

	var r0 *dynamodb.PutItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) *dynamodb.PutItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDBClient_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type DynamoDBClient_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.PutItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDBClient_Expecter) PutItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDBClient_PutItem_Call {
	return &DynamoDBClient_PutItem_Call{Call: _e.mock.On("PutItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDBClient_PutItem_Call) Run(run func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options))) *DynamoDBClient_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.PutItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDBClient_PutItem_Call) Return(_a0 *dynamodb.PutItemOutput, _a1 error) *DynamoDBClient_PutItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDBClient_PutItem_Call) RunAndReturn(run func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)) *DynamoDBClient_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, input, optFns
func (_m *DynamoDBClient) Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *dynamodb.QueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)); ok {
		return rf(ctx, input, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) *dynamodb.QueryOutput); ok {
		r0 = rf(ctx, input, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.QueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, input, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDBClient_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type DynamoDBClient_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - input *dynamodb.QueryInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDBClient_Expecter) Query(ctx interface{}, input interface{}, optFns ...interface{}) *DynamoDBClient_Query_Call {
	return &DynamoDBClient_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, input}, optFns...)...)}
}

func (_c *DynamoDBClient_Query_Call) Run(run func(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options))) *DynamoDBClient_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.QueryInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDBClient_Query_Call) Return(_a0 *dynamodb.QueryOutput, _a1 error) *DynamoDBClient_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDBClient_Query_Call) RunAndReturn(run func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)) *DynamoDBClient_Query_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDBClient) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItem")
	}

	var r0 *dynamodb.UpdateItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) *dynamodb.UpdateItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.UpdateItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDBClient_UpdateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItem'
type DynamoDBClient_UpdateItem_Call struct {
	*mock.Call
}

// UpdateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.UpdateItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDBClient_Expecter) UpdateItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDBClient_UpdateItem_Call {
	return &DynamoDBClient_UpdateItem_Call{Call: _e.mock.On("UpdateItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDBClient_UpdateItem_Call) Run(run func(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options))) *DynamoDBClient_UpdateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.UpdateItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDBClient_UpdateItem_Call) Return(_a0 *dynamodb.UpdateItemOutput, _a1 error) *DynamoDBClient_UpdateItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDBClient_UpdateItem_Call) RunAndReturn(run func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)) *DynamoDBClient_UpdateItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewDynamoDBClient creates a new instance of DynamoDBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDynamoDBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DynamoDBClient {
	mock := &DynamoDBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
