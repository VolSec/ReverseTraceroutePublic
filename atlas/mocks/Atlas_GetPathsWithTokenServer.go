package mocks

import "github.com/VolSec/ReverseTraceroutePublic/atlas/pb"
import "github.com/stretchr/testify/mock"

type Atlas_GetPathsWithTokenServer struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *Atlas_GetPathsWithTokenServer) Send(_a0 *pb.TokenResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pb.TokenResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Recv provides a mock function with given fields:
func (_m *Atlas_GetPathsWithTokenServer) Recv() (*pb.TokenRequest, error) {
	ret := _m.Called()

	var r0 *pb.TokenRequest
	if rf, ok := ret.Get(0).(func() *pb.TokenRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.TokenRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
