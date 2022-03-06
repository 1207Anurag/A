// Code generated by MockGen. DO NOT EDIT.
// Source: Interface.go

// Package store is a generated GoMock package.
package store

import (
	model "Hospital/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPatient_er is a mock of Patient_er interface.
type MockPatient_er struct {
	ctrl     *gomock.Controller
	recorder *MockPatient_erMockRecorder
}

// MockPatient_erMockRecorder is the mock recorder for MockPatient_er.
type MockPatient_erMockRecorder struct {
	mock *MockPatient_er
}

// NewMockPatient_er creates a new mock instance.
func NewMockPatient_er(ctrl *gomock.Controller) *MockPatient_er {
	mock := &MockPatient_er{ctrl: ctrl}
	mock.recorder = &MockPatient_erMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPatient_er) EXPECT() *MockPatient_erMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPatient_er) Get(id int) (*model.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPatient_erMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPatient_er)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockPatient_er) GetAll() (model.Patient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(model.Patient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPatient_erMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPatient_er)(nil).GetAll))
}

// PostPatient mocks base method.
func (m *MockPatient_er) PostPatient(Id int, Name, Phone string, Discharge bool, BloodGroup, Description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostPatient", Id, Name, Phone, Discharge, BloodGroup, Description)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostPatient indicates an expected call of PostPatient.
func (mr *MockPatient_erMockRecorder) PostPatient(Id, Name, Phone, Discharge, BloodGroup, Description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostPatient", reflect.TypeOf((*MockPatient_er)(nil).PostPatient), Id, Name, Phone, Discharge, BloodGroup, Description)
}

// Update mocks base method.
func (m *MockPatient_er) Update(id int, Name, desc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", id, Name, desc)
}

// Update indicates an expected call of Update.
func (mr *MockPatient_erMockRecorder) Update(id, Name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPatient_er)(nil).Update), id, Name, desc)
}