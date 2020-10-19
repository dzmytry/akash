// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// NetworkingV1beta1Interface is an autogenerated mock type for the NetworkingV1beta1Interface type
type NetworkingV1beta1Interface struct {
	mock.Mock
}

// IngressClasses provides a mock function with given fields:
func (_m *NetworkingV1beta1Interface) IngressClasses() v1beta1.IngressClassInterface {
	ret := _m.Called()

	var r0 v1beta1.IngressClassInterface
	if rf, ok := ret.Get(0).(func() v1beta1.IngressClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.IngressClassInterface)
		}
	}

	return r0
}

// Ingresses provides a mock function with given fields: namespace
func (_m *NetworkingV1beta1Interface) Ingresses(namespace string) v1beta1.IngressInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.IngressInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.IngressInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.IngressInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *NetworkingV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}
