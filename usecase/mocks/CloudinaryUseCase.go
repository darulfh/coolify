// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "github.com/darulfh/skuy_pay_be/model"

	mock "github.com/stretchr/testify/mock"
)

// CloudinaryUseCase is an autogenerated mock type for the CloudinaryUseCase type
type CloudinaryUseCase struct {
	mock.Mock
}

// SendingMail provides a mock function with given fields: payload
func (_m *CloudinaryUseCase) SendingMail(payload model.PayloadMail) {
	_m.Called(payload)
}

// NewCloudinaryUseCase creates a new instance of CloudinaryUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudinaryUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudinaryUseCase {
	mock := &CloudinaryUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
