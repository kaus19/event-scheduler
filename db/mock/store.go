// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaus19/event-scheduler/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/kaus19/event-scheduler/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockStore) CreateEvent(arg0 context.Context, arg1 db.CreateEventParams) (db.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1)
	ret0, _ := ret[0].(db.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockStoreMockRecorder) CreateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockStore)(nil).CreateEvent), arg0, arg1)
}

// CreateTimeSlotEvent mocks base method.
func (m *MockStore) CreateTimeSlotEvent(arg0 context.Context, arg1 db.CreateTimeSlotEventParams) (db.TimeSlotsEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTimeSlotEvent", arg0, arg1)
	ret0, _ := ret[0].(db.TimeSlotsEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTimeSlotEvent indicates an expected call of CreateTimeSlotEvent.
func (mr *MockStoreMockRecorder) CreateTimeSlotEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTimeSlotEvent", reflect.TypeOf((*MockStore)(nil).CreateTimeSlotEvent), arg0, arg1)
}

// CreateTimeSlotUser mocks base method.
func (m *MockStore) CreateTimeSlotUser(arg0 context.Context, arg1 db.CreateTimeSlotUserParams) (db.TimeSlotsUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTimeSlotUser", arg0, arg1)
	ret0, _ := ret[0].(db.TimeSlotsUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTimeSlotUser indicates an expected call of CreateTimeSlotUser.
func (mr *MockStoreMockRecorder) CreateTimeSlotUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTimeSlotUser", reflect.TypeOf((*MockStore)(nil).CreateTimeSlotUser), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteEvent mocks base method.
func (m *MockStore) DeleteEvent(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockStoreMockRecorder) DeleteEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockStore)(nil).DeleteEvent), arg0, arg1)
}

// DeleteTimePreferenceEvent mocks base method.
func (m *MockStore) DeleteTimePreferenceEvent(arg0 context.Context, arg1 db.DeleteTimePreferenceEventParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTimePreferenceEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTimePreferenceEvent indicates an expected call of DeleteTimePreferenceEvent.
func (mr *MockStoreMockRecorder) DeleteTimePreferenceEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTimePreferenceEvent", reflect.TypeOf((*MockStore)(nil).DeleteTimePreferenceEvent), arg0, arg1)
}

// DeleteTimePreferenceUser mocks base method.
func (m *MockStore) DeleteTimePreferenceUser(arg0 context.Context, arg1 db.DeleteTimePreferenceUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTimePreferenceUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTimePreferenceUser indicates an expected call of DeleteTimePreferenceUser.
func (mr *MockStoreMockRecorder) DeleteTimePreferenceUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTimePreferenceUser", reflect.TypeOf((*MockStore)(nil).DeleteTimePreferenceUser), arg0, arg1)
}

// GetEventByID mocks base method.
func (m *MockStore) GetEventByID(arg0 context.Context, arg1 int32) (db.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventByID", arg0, arg1)
	ret0, _ := ret[0].(db.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventByID indicates an expected call of GetEventByID.
func (mr *MockStoreMockRecorder) GetEventByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventByID", reflect.TypeOf((*MockStore)(nil).GetEventByID), arg0, arg1)
}

// GetTimePreferencesByEvent mocks base method.
func (m *MockStore) GetTimePreferencesByEvent(arg0 context.Context, arg1 int32) ([]db.TimeSlotsEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimePreferencesByEvent", arg0, arg1)
	ret0, _ := ret[0].([]db.TimeSlotsEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimePreferencesByEvent indicates an expected call of GetTimePreferencesByEvent.
func (mr *MockStoreMockRecorder) GetTimePreferencesByEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimePreferencesByEvent", reflect.TypeOf((*MockStore)(nil).GetTimePreferencesByEvent), arg0, arg1)
}

// GetTimePreferencesByUser mocks base method.
func (m *MockStore) GetTimePreferencesByUser(arg0 context.Context, arg1 int32) ([]db.TimeSlotsUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimePreferencesByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.TimeSlotsUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimePreferencesByUser indicates an expected call of GetTimePreferencesByUser.
func (mr *MockStoreMockRecorder) GetTimePreferencesByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimePreferencesByUser", reflect.TypeOf((*MockStore)(nil).GetTimePreferencesByUser), arg0, arg1)
}

// GetTimePreferencesForAllUsers mocks base method.
func (m *MockStore) GetTimePreferencesForAllUsers(arg0 context.Context) ([]db.GetTimePreferencesForAllUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimePreferencesForAllUsers", arg0)
	ret0, _ := ret[0].([]db.GetTimePreferencesForAllUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimePreferencesForAllUsers indicates an expected call of GetTimePreferencesForAllUsers.
func (mr *MockStoreMockRecorder) GetTimePreferencesForAllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimePreferencesForAllUsers", reflect.TypeOf((*MockStore)(nil).GetTimePreferencesForAllUsers), arg0)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 context.Context, arg1 int32) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0, arg1)
}

// ListEvents mocks base method.
func (m *MockStore) ListEvents(arg0 context.Context) ([]db.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEvents", arg0)
	ret0, _ := ret[0].([]db.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEvents indicates an expected call of ListEvents.
func (mr *MockStoreMockRecorder) ListEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEvents", reflect.TypeOf((*MockStore)(nil).ListEvents), arg0)
}

// ListEventsByOrganizer mocks base method.
func (m *MockStore) ListEventsByOrganizer(arg0 context.Context, arg1 int32) ([]db.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventsByOrganizer", arg0, arg1)
	ret0, _ := ret[0].([]db.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventsByOrganizer indicates an expected call of ListEventsByOrganizer.
func (mr *MockStoreMockRecorder) ListEventsByOrganizer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventsByOrganizer", reflect.TypeOf((*MockStore)(nil).ListEventsByOrganizer), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0)
}

// UpdateEvent mocks base method.
func (m *MockStore) UpdateEvent(arg0 context.Context, arg1 db.UpdateEventParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockStoreMockRecorder) UpdateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockStore)(nil).UpdateEvent), arg0, arg1)
}

// UpdateTimePreferenceEvent mocks base method.
func (m *MockStore) UpdateTimePreferenceEvent(arg0 context.Context, arg1 db.UpdateTimePreferenceEventParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimePreferenceEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimePreferenceEvent indicates an expected call of UpdateTimePreferenceEvent.
func (mr *MockStoreMockRecorder) UpdateTimePreferenceEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimePreferenceEvent", reflect.TypeOf((*MockStore)(nil).UpdateTimePreferenceEvent), arg0, arg1)
}

// UpdateTimePreferenceUser mocks base method.
func (m *MockStore) UpdateTimePreferenceUser(arg0 context.Context, arg1 db.UpdateTimePreferenceUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimePreferenceUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimePreferenceUser indicates an expected call of UpdateTimePreferenceUser.
func (mr *MockStoreMockRecorder) UpdateTimePreferenceUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimePreferenceUser", reflect.TypeOf((*MockStore)(nil).UpdateTimePreferenceUser), arg0, arg1)
}