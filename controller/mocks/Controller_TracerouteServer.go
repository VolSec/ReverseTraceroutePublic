package mocks

import (
	"github.com/VolSec/ReverseTraceroutePublic/datamodel"
	"github.com/stretchr/testify/mock"
)

type Controller_TracerouteServer struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *Controller_TracerouteServer) Send(_a0 *datamodel.Traceroute) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datamodel.Traceroute) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
