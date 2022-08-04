// Code generated by MockGen. DO NOT EDIT.
// Source: /home/basili4/GolandProjects/service/pkg/keyvalue/memcached.go

// Package mock_keyvalue is a generated GoMock package.
package mock_cache

import (
	reflect "reflect"

	memcache "github.com/bradfitz/gomemcache/memcache"
	gomock "github.com/golang/mock/gomock"
)

// MockMemcachedClient is a mock of MemcachedClient interface.
type MockMemcachedClient struct {
	ctrl     *gomock.Controller
	recorder *MockMemcachedClientMockRecorder
}

// MockMemcachedClientMockRecorder is the mock recorder for MockMemcachedClient.
type MockMemcachedClientMockRecorder struct {
	mock *MockMemcachedClient
}

// NewMockMemcachedClient creates a new mock instance.
func NewMockMemcachedClient(ctrl *gomock.Controller) *MockMemcachedClient {
	mock := &MockMemcachedClient{ctrl: ctrl}
	mock.recorder = &MockMemcachedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemcachedClient) EXPECT() *MockMemcachedClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMemcachedClient) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMemcachedClientMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemcachedClient)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockMemcachedClient) Get(key string) (*memcache.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*memcache.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMemcachedClientMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMemcachedClient)(nil).Get), key)
}

// Set mocks base method.
func (m *MockMemcachedClient) Set(item *memcache.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockMemcachedClientMockRecorder) Set(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMemcachedClient)(nil).Set), item)
}
