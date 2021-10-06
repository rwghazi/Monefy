// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	investplans "github.com/Rahmanwghazi/Monefy/business/investplans"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: idProduct, domain
func (_m *Usecase) Create(idProduct string, domain *investplans.InvestPlanDomain) (investplans.InvestPlanDomain, error) {
	ret := _m.Called(idProduct, domain)

	var r0 investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(string, *investplans.InvestPlanDomain) investplans.InvestPlanDomain); ok {
		r0 = rf(idProduct, domain)
	} else {
		r0 = ret.Get(0).(investplans.InvestPlanDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *investplans.InvestPlanDomain) error); ok {
		r1 = rf(idProduct, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlan provides a mock function with given fields: domain, id
func (_m *Usecase) DeletePlan(domain *investplans.InvestPlanDomain, id uint) (string, error) {
	ret := _m.Called(domain, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain, uint) string); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditPlan provides a mock function with given fields: domain, id
func (_m *Usecase) EditPlan(domain *investplans.InvestPlanDomain, id uint) (investplans.InvestPlanDomain, error) {
	ret := _m.Called(domain, id)

	var r0 investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain, uint) investplans.InvestPlanDomain); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(investplans.InvestPlanDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlanById provides a mock function with given fields: domain, id
func (_m *Usecase) GetPlanById(domain *investplans.InvestPlanDomain, id uint) (investplans.InvestPlanDomain, error) {
	ret := _m.Called(domain, id)

	var r0 investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain, uint) investplans.InvestPlanDomain); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(investplans.InvestPlanDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlans provides a mock function with given fields: domain
func (_m *Usecase) GetPlans(domain *investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	ret := _m.Called(domain)

	var r0 []investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain) []investplans.InvestPlanDomain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]investplans.InvestPlanDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnfinishedPlans provides a mock function with given fields: domain
func (_m *Usecase) GetUnfinishedPlans(domain *investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	ret := _m.Called(domain)

	var r0 []investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain) []investplans.InvestPlanDomain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]investplans.InvestPlanDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetfinishedPlans provides a mock function with given fields: domain
func (_m *Usecase) GetfinishedPlans(domain *investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	ret := _m.Called(domain)

	var r0 []investplans.InvestPlanDomain
	if rf, ok := ret.Get(0).(func(*investplans.InvestPlanDomain) []investplans.InvestPlanDomain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]investplans.InvestPlanDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*investplans.InvestPlanDomain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
