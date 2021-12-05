// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/jiachen/Git/mgcache/storage/redis_storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisClientInterface is a mock of RedisClientInterface interface.
type MockRedisClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientInterfaceMockRecorder
}

// MockRedisClientInterfaceMockRecorder is the mock recorder for MockRedisClientInterface.
type MockRedisClientInterfaceMockRecorder struct {
	mock *MockRedisClientInterface
}

// NewMockRedisClientInterface creates a new mock instance.
func NewMockRedisClientInterface(ctrl *gomock.Controller) *MockRedisClientInterface {
	mock := &MockRedisClientInterface{ctrl: ctrl}
	mock.recorder = &MockRedisClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClientInterface) EXPECT() *MockRedisClientInterfaceMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockRedisClientInterface) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisClientInterfaceMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClientInterface)(nil).Del), varargs...)
}

// Get mocks base method.
func (m *MockRedisClientInterface) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisClientInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClientInterface)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockRedisClientInterface) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisClientInterfaceMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClientInterface)(nil).Set), ctx, key, value, expiration)
}
