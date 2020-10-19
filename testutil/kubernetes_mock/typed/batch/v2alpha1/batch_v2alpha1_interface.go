// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
)

// BatchV2alpha1Interface is an autogenerated mock type for the BatchV2alpha1Interface type
type BatchV2alpha1Interface struct {
	mock.Mock
}

// CronJobs provides a mock function with given fields: namespace
func (_m *BatchV2alpha1Interface) CronJobs(namespace string) v2alpha1.CronJobInterface {
	ret := _m.Called(namespace)

	var r0 v2alpha1.CronJobInterface
	if rf, ok := ret.Get(0).(func(string) v2alpha1.CronJobInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2alpha1.CronJobInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *BatchV2alpha1Interface) RESTClient() rest.Interface {
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
