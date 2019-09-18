// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/model"
	mock "github.com/stretchr/testify/mock"
)

// CommandWebhookStore is an autogenerated mock type for the CommandWebhookStore type
type CommandWebhookStore struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *CommandWebhookStore) Cleanup() {
	_m.Called()
}

// Get provides a mock function with given fields: id
func (_m *CommandWebhookStore) Get(id string) (*model.CommandWebhook, *model.AppError) {
	ret := _m.Called(id)

	var r0 *model.CommandWebhook
	if rf, ok := ret.Get(0).(func(string) *model.CommandWebhook); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Save provides a mock function with given fields: webhook
func (_m *CommandWebhookStore) Save(webhook *model.CommandWebhook) (*model.CommandWebhook, *model.AppError) {
	ret := _m.Called(webhook)

	var r0 *model.CommandWebhook
	if rf, ok := ret.Get(0).(func(*model.CommandWebhook) *model.CommandWebhook); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.CommandWebhook) *model.AppError); ok {
		r1 = rf(webhook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// TryUse provides a mock function with given fields: id, limit
func (_m *CommandWebhookStore) TryUse(id string, limit int) *model.AppError {
	ret := _m.Called(id, limit)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int) *model.AppError); ok {
		r0 = rf(id, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}
