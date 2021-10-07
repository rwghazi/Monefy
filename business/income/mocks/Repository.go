// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	income "github.com/Rahmanwghazi/Monefy/business/income"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateIncome provides a mock function with given fields: domain
func (_m *Repository) CreateIncome(domain *income.IncomeDomain) (income.IncomeDomain, error) {
	ret := _m.Called(domain)

	var r0 income.IncomeDomain
	if rf, ok := ret.Get(0).(func(*income.IncomeDomain) income.IncomeDomain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(income.IncomeDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*income.IncomeDomain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIncome provides a mock function with given fields: domain, id
func (_m *Repository) DeleteIncome(domain *income.IncomeDomain, id uint) (string, error) {
	ret := _m.Called(domain, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(*income.IncomeDomain, uint) string); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*income.IncomeDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditIncome provides a mock function with given fields: domain, id
func (_m *Repository) EditIncome(domain *income.IncomeDomain, id uint) (income.IncomeDomain, error) {
	ret := _m.Called(domain, id)

	var r0 income.IncomeDomain
	if rf, ok := ret.Get(0).(func(*income.IncomeDomain, uint) income.IncomeDomain); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(income.IncomeDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*income.IncomeDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncome provides a mock function with given fields: domain
func (_m *Repository) GetIncome(domain *income.IncomeDomain) ([]income.IncomeDomain, error) {
	ret := _m.Called(domain)

	var r0 []income.IncomeDomain
	if rf, ok := ret.Get(0).(func(*income.IncomeDomain) []income.IncomeDomain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]income.IncomeDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*income.IncomeDomain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncomeById provides a mock function with given fields: domain, id
func (_m *Repository) GetIncomeById(domain *income.IncomeDomain, id uint) (income.IncomeDomain, error) {
	ret := _m.Called(domain, id)

	var r0 income.IncomeDomain
	if rf, ok := ret.Get(0).(func(*income.IncomeDomain, uint) income.IncomeDomain); ok {
		r0 = rf(domain, id)
	} else {
		r0 = ret.Get(0).(income.IncomeDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*income.IncomeDomain, uint) error); ok {
		r1 = rf(domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}