package mocks

import "github.com/stretchr/testify/mock"

type UUIDGeneratorMock struct {
	mock.Mock
}

func (u *UUIDGeneratorMock) New() string {
	args := u.Called()
	return args.String(0)
}
