// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/common/log.go

// Package common is a generated GoMock package.
package common

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), i...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Error mocks base method.
func (m *MockLogger) Error(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), i...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), i...)
}

// Fatalf mocks base method.
func (m *MockLogger) Fatalf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockLoggerMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// Info mocks base method.
func (m *MockLogger) Info(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), i...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Output mocks base method.
func (m *MockLogger) Output() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Output indicates an expected call of Output.
func (mr *MockLoggerMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockLogger)(nil).Output))
}

// Panic mocks base method.
func (m *MockLogger) Panic(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic.
func (mr *MockLoggerMockRecorder) Panic(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockLogger)(nil).Panic), i...)
}

// Panicf mocks base method.
func (m *MockLogger) Panicf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panicf", varargs...)
}

// Panicf indicates an expected call of Panicf.
func (mr *MockLoggerMockRecorder) Panicf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panicf", reflect.TypeOf((*MockLogger)(nil).Panicf), varargs...)
}

// Prefix mocks base method.
func (m *MockLogger) Prefix() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prefix")
	ret0, _ := ret[0].(string)
	return ret0
}

// Prefix indicates an expected call of Prefix.
func (mr *MockLoggerMockRecorder) Prefix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefix", reflect.TypeOf((*MockLogger)(nil).Prefix))
}

// Print mocks base method.
func (m *MockLogger) Print(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Print", varargs...)
}

// Print indicates an expected call of Print.
func (mr *MockLoggerMockRecorder) Print(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockLogger)(nil).Print), i...)
}

// Printf mocks base method.
func (m *MockLogger) Printf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockLoggerMockRecorder) Printf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockLogger)(nil).Printf), varargs...)
}

// SetHeader mocks base method.
func (m *MockLogger) SetHeader(h string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeader", h)
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockLoggerMockRecorder) SetHeader(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockLogger)(nil).SetHeader), h)
}

// SetOutput mocks base method.
func (m *MockLogger) SetOutput(w io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOutput", w)
}

// SetOutput indicates an expected call of SetOutput.
func (mr *MockLoggerMockRecorder) SetOutput(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOutput", reflect.TypeOf((*MockLogger)(nil).SetOutput), w)
}

// SetPrefix mocks base method.
func (m *MockLogger) SetPrefix(p string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPrefix", p)
}

// SetPrefix indicates an expected call of SetPrefix.
func (mr *MockLoggerMockRecorder) SetPrefix(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrefix", reflect.TypeOf((*MockLogger)(nil).SetPrefix), p)
}

// Warn mocks base method.
func (m *MockLogger) Warn(i ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), i...)
}

// Warnf mocks base method.
func (m *MockLogger) Warnf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockLoggerMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}
