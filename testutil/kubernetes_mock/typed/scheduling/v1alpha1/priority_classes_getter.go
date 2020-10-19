// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
)

// PriorityClassesGetter is an autogenerated mock type for the PriorityClassesGetter type
type PriorityClassesGetter struct {
	mock.Mock
}

// PriorityClasses provides a mock function with given fields:
func (_m *PriorityClassesGetter) PriorityClasses() v1alpha1.PriorityClassInterface {
	ret := _m.Called()

	var r0 v1alpha1.PriorityClassInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.PriorityClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.PriorityClassInterface)
		}
	}

	return r0
}
