// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddTag provides a mock function with given fields: tag
func (_m *Client) AddTag(tag string) {
	_m.Called(tag)
}

// SendCountMetric provides a mock function with given fields: name, tags
func (_m *Client) SendCountMetric(name string, tags ...string) error {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(name, tags...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDurationMetric provides a mock function with given fields: name, t1, t2, tags
func (_m *Client) SendDurationMetric(name string, t1 time.Time, t2 time.Time, tags ...string) error {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, t1, t2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time, ...string) error); ok {
		r0 = rf(name, t1, t2, tags...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMetric provides a mock function with given fields: name, tags
func (_m *Client) SendMetric(name string, tags ...string) error {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(name, tags...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
