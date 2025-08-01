// Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wmcram/dcgm-exporter/internal/pkg/dcgmprovider (interfaces: DCGM)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/pkg/dcgmprovider/mock_client.go -package=dcgmprovider -copyright_file=../../../hack/header.txt . DCGM
//

// Package dcgmprovider is a generated GoMock package.
package dcgmprovider

import (
	reflect "reflect"
	time "time"

	dcgm "github.com/NVIDIA/go-dcgm/pkg/dcgm"
	gomock "go.uber.org/mock/gomock"
)

// MockDCGM is a mock of DCGM interface.
type MockDCGM struct {
	ctrl     *gomock.Controller
	recorder *MockDCGMMockRecorder
	isgomock struct{}
}

// MockDCGMMockRecorder is the mock recorder for MockDCGM.
type MockDCGMMockRecorder struct {
	mock *MockDCGM
}

// NewMockDCGM creates a new mock instance.
func NewMockDCGM(ctrl *gomock.Controller) *MockDCGM {
	mock := &MockDCGM{ctrl: ctrl}
	mock.recorder = &MockDCGMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDCGM) EXPECT() *MockDCGMMockRecorder {
	return m.recorder
}

// AddEntityToGroup mocks base method.
func (m *MockDCGM) AddEntityToGroup(arg0 dcgm.GroupHandle, arg1 dcgm.Field_Entity_Group, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntityToGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntityToGroup indicates an expected call of AddEntityToGroup.
func (mr *MockDCGMMockRecorder) AddEntityToGroup(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntityToGroup", reflect.TypeOf((*MockDCGM)(nil).AddEntityToGroup), arg0, arg1, arg2)
}

// AddLinkEntityToGroup mocks base method.
func (m *MockDCGM) AddLinkEntityToGroup(arg0 dcgm.GroupHandle, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLinkEntityToGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLinkEntityToGroup indicates an expected call of AddLinkEntityToGroup.
func (mr *MockDCGMMockRecorder) AddLinkEntityToGroup(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLinkEntityToGroup", reflect.TypeOf((*MockDCGM)(nil).AddLinkEntityToGroup), arg0, arg1, arg2)
}

// Cleanup mocks base method.
func (m *MockDCGM) Cleanup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cleanup")
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockDCGMMockRecorder) Cleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockDCGM)(nil).Cleanup))
}

// CreateFakeEntities mocks base method.
func (m *MockDCGM) CreateFakeEntities(entities []dcgm.MigHierarchyInfo) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFakeEntities", entities)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFakeEntities indicates an expected call of CreateFakeEntities.
func (mr *MockDCGMMockRecorder) CreateFakeEntities(entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFakeEntities", reflect.TypeOf((*MockDCGM)(nil).CreateFakeEntities), entities)
}

// CreateGroup mocks base method.
func (m *MockDCGM) CreateGroup(arg0 string) (dcgm.GroupHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0)
	ret0, _ := ret[0].(dcgm.GroupHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockDCGMMockRecorder) CreateGroup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockDCGM)(nil).CreateGroup), arg0)
}

// DestroyGroup mocks base method.
func (m *MockDCGM) DestroyGroup(groupID dcgm.GroupHandle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyGroup", groupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyGroup indicates an expected call of DestroyGroup.
func (mr *MockDCGMMockRecorder) DestroyGroup(groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyGroup", reflect.TypeOf((*MockDCGM)(nil).DestroyGroup), groupID)
}

// EntitiesGetLatestValues mocks base method.
func (m *MockDCGM) EntitiesGetLatestValues(arg0 []dcgm.GroupEntityPair, arg1 []dcgm.Short, arg2 uint) ([]dcgm.FieldValue_v2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntitiesGetLatestValues", arg0, arg1, arg2)
	ret0, _ := ret[0].([]dcgm.FieldValue_v2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EntitiesGetLatestValues indicates an expected call of EntitiesGetLatestValues.
func (mr *MockDCGMMockRecorder) EntitiesGetLatestValues(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntitiesGetLatestValues", reflect.TypeOf((*MockDCGM)(nil).EntitiesGetLatestValues), arg0, arg1, arg2)
}

// EntityGetLatestValues mocks base method.
func (m *MockDCGM) EntityGetLatestValues(arg0 dcgm.Field_Entity_Group, arg1 uint, arg2 []dcgm.Short) ([]dcgm.FieldValue_v1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityGetLatestValues", arg0, arg1, arg2)
	ret0, _ := ret[0].([]dcgm.FieldValue_v1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EntityGetLatestValues indicates an expected call of EntityGetLatestValues.
func (mr *MockDCGMMockRecorder) EntityGetLatestValues(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityGetLatestValues", reflect.TypeOf((*MockDCGM)(nil).EntityGetLatestValues), arg0, arg1, arg2)
}

// FieldGetByID mocks base method.
func (m *MockDCGM) FieldGetByID(arg0 dcgm.Short) dcgm.FieldMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldGetByID", arg0)
	ret0, _ := ret[0].(dcgm.FieldMeta)
	return ret0
}

// FieldGetByID indicates an expected call of FieldGetByID.
func (mr *MockDCGMMockRecorder) FieldGetByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldGetByID", reflect.TypeOf((*MockDCGM)(nil).FieldGetByID), arg0)
}

// FieldGroupCreate mocks base method.
func (m *MockDCGM) FieldGroupCreate(arg0 string, arg1 []dcgm.Short) (dcgm.FieldHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldGroupCreate", arg0, arg1)
	ret0, _ := ret[0].(dcgm.FieldHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FieldGroupCreate indicates an expected call of FieldGroupCreate.
func (mr *MockDCGMMockRecorder) FieldGroupCreate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldGroupCreate", reflect.TypeOf((*MockDCGM)(nil).FieldGroupCreate), arg0, arg1)
}

// FieldGroupDestroy mocks base method.
func (m *MockDCGM) FieldGroupDestroy(arg0 dcgm.FieldHandle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldGroupDestroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FieldGroupDestroy indicates an expected call of FieldGroupDestroy.
func (mr *MockDCGMMockRecorder) FieldGroupDestroy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldGroupDestroy", reflect.TypeOf((*MockDCGM)(nil).FieldGroupDestroy), arg0)
}

// Fv2_String mocks base method.
func (m *MockDCGM) Fv2_String(fv dcgm.FieldValue_v2) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fv2_String", fv)
	ret0, _ := ret[0].(string)
	return ret0
}

// Fv2_String indicates an expected call of Fv2_String.
func (mr *MockDCGMMockRecorder) Fv2_String(fv any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fv2_String", reflect.TypeOf((*MockDCGM)(nil).Fv2_String), fv)
}

// GetAllDeviceCount mocks base method.
func (m *MockDCGM) GetAllDeviceCount() (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDeviceCount")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDeviceCount indicates an expected call of GetAllDeviceCount.
func (mr *MockDCGMMockRecorder) GetAllDeviceCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDeviceCount", reflect.TypeOf((*MockDCGM)(nil).GetAllDeviceCount))
}

// GetCPUHierarchy mocks base method.
func (m *MockDCGM) GetCPUHierarchy() (dcgm.CPUHierarchy_v1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCPUHierarchy")
	ret0, _ := ret[0].(dcgm.CPUHierarchy_v1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCPUHierarchy indicates an expected call of GetCPUHierarchy.
func (mr *MockDCGMMockRecorder) GetCPUHierarchy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCPUHierarchy", reflect.TypeOf((*MockDCGM)(nil).GetCPUHierarchy))
}

// GetDeviceInfo mocks base method.
func (m *MockDCGM) GetDeviceInfo(arg0 uint) (dcgm.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceInfo", arg0)
	ret0, _ := ret[0].(dcgm.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceInfo indicates an expected call of GetDeviceInfo.
func (mr *MockDCGMMockRecorder) GetDeviceInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceInfo", reflect.TypeOf((*MockDCGM)(nil).GetDeviceInfo), arg0)
}

// GetEntityGroupEntities mocks base method.
func (m *MockDCGM) GetEntityGroupEntities(entityGroup dcgm.Field_Entity_Group) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntityGroupEntities", entityGroup)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntityGroupEntities indicates an expected call of GetEntityGroupEntities.
func (mr *MockDCGMMockRecorder) GetEntityGroupEntities(entityGroup any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntityGroupEntities", reflect.TypeOf((*MockDCGM)(nil).GetEntityGroupEntities), entityGroup)
}

// GetGPUInstanceHierarchy mocks base method.
func (m *MockDCGM) GetGPUInstanceHierarchy() (dcgm.MigHierarchy_v2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGPUInstanceHierarchy")
	ret0, _ := ret[0].(dcgm.MigHierarchy_v2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGPUInstanceHierarchy indicates an expected call of GetGPUInstanceHierarchy.
func (mr *MockDCGMMockRecorder) GetGPUInstanceHierarchy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGPUInstanceHierarchy", reflect.TypeOf((*MockDCGM)(nil).GetGPUInstanceHierarchy))
}

// GetGroupInfo mocks base method.
func (m *MockDCGM) GetGroupInfo(groupID dcgm.GroupHandle) (*dcgm.GroupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupInfo", groupID)
	ret0, _ := ret[0].(*dcgm.GroupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupInfo indicates an expected call of GetGroupInfo.
func (mr *MockDCGMMockRecorder) GetGroupInfo(groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupInfo", reflect.TypeOf((*MockDCGM)(nil).GetGroupInfo), groupID)
}

// GetNvLinkLinkStatus mocks base method.
func (m *MockDCGM) GetNvLinkLinkStatus() ([]dcgm.NvLinkStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNvLinkLinkStatus")
	ret0, _ := ret[0].([]dcgm.NvLinkStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNvLinkLinkStatus indicates an expected call of GetNvLinkLinkStatus.
func (mr *MockDCGMMockRecorder) GetNvLinkLinkStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNvLinkLinkStatus", reflect.TypeOf((*MockDCGM)(nil).GetNvLinkLinkStatus))
}

// GetSupportedDevices mocks base method.
func (m *MockDCGM) GetSupportedDevices() ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedDevices")
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedDevices indicates an expected call of GetSupportedDevices.
func (mr *MockDCGMMockRecorder) GetSupportedDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedDevices", reflect.TypeOf((*MockDCGM)(nil).GetSupportedDevices))
}

// GetSupportedMetricGroups mocks base method.
func (m *MockDCGM) GetSupportedMetricGroups(arg0 uint) ([]dcgm.MetricGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedMetricGroups", arg0)
	ret0, _ := ret[0].([]dcgm.MetricGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedMetricGroups indicates an expected call of GetSupportedMetricGroups.
func (mr *MockDCGMMockRecorder) GetSupportedMetricGroups(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedMetricGroups", reflect.TypeOf((*MockDCGM)(nil).GetSupportedMetricGroups), arg0)
}

// GetValuesSince mocks base method.
func (m *MockDCGM) GetValuesSince(arg0 dcgm.GroupHandle, arg1 dcgm.FieldHandle, arg2 time.Time) ([]dcgm.FieldValue_v2, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValuesSince", arg0, arg1, arg2)
	ret0, _ := ret[0].([]dcgm.FieldValue_v2)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetValuesSince indicates an expected call of GetValuesSince.
func (mr *MockDCGMMockRecorder) GetValuesSince(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValuesSince", reflect.TypeOf((*MockDCGM)(nil).GetValuesSince), arg0, arg1, arg2)
}

// GroupAllGPUs mocks base method.
func (m *MockDCGM) GroupAllGPUs() dcgm.GroupHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupAllGPUs")
	ret0, _ := ret[0].(dcgm.GroupHandle)
	return ret0
}

// GroupAllGPUs indicates an expected call of GroupAllGPUs.
func (mr *MockDCGMMockRecorder) GroupAllGPUs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupAllGPUs", reflect.TypeOf((*MockDCGM)(nil).GroupAllGPUs))
}

// HealthCheck mocks base method.
func (m *MockDCGM) HealthCheck(groupID dcgm.GroupHandle) (dcgm.HealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", groupID)
	ret0, _ := ret[0].(dcgm.HealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockDCGMMockRecorder) HealthCheck(groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDCGM)(nil).HealthCheck), groupID)
}

// HealthGet mocks base method.
func (m *MockDCGM) HealthGet(groupID dcgm.GroupHandle) (dcgm.HealthSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthGet", groupID)
	ret0, _ := ret[0].(dcgm.HealthSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthGet indicates an expected call of HealthGet.
func (mr *MockDCGMMockRecorder) HealthGet(groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthGet", reflect.TypeOf((*MockDCGM)(nil).HealthGet), groupID)
}

// HealthSet mocks base method.
func (m *MockDCGM) HealthSet(groupID dcgm.GroupHandle, systems dcgm.HealthSystem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthSet", groupID, systems)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthSet indicates an expected call of HealthSet.
func (mr *MockDCGMMockRecorder) HealthSet(groupID, systems any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthSet", reflect.TypeOf((*MockDCGM)(nil).HealthSet), groupID, systems)
}

// InjectFieldValue mocks base method.
func (m *MockDCGM) InjectFieldValue(gpu uint, fieldID dcgm.Short, fieldType uint, status int, ts int64, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectFieldValue", gpu, fieldID, fieldType, status, ts, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectFieldValue indicates an expected call of InjectFieldValue.
func (mr *MockDCGMMockRecorder) InjectFieldValue(gpu, fieldID, fieldType, status, ts, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectFieldValue", reflect.TypeOf((*MockDCGM)(nil).InjectFieldValue), gpu, fieldID, fieldType, status, ts, value)
}

// LinkGetLatestValues mocks base method.
func (m *MockDCGM) LinkGetLatestValues(arg0, arg1 uint, arg2 []dcgm.Short) ([]dcgm.FieldValue_v1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkGetLatestValues", arg0, arg1, arg2)
	ret0, _ := ret[0].([]dcgm.FieldValue_v1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkGetLatestValues indicates an expected call of LinkGetLatestValues.
func (mr *MockDCGMMockRecorder) LinkGetLatestValues(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkGetLatestValues", reflect.TypeOf((*MockDCGM)(nil).LinkGetLatestValues), arg0, arg1, arg2)
}

// NewDefaultGroup mocks base method.
func (m *MockDCGM) NewDefaultGroup(arg0 string) (dcgm.GroupHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDefaultGroup", arg0)
	ret0, _ := ret[0].(dcgm.GroupHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDefaultGroup indicates an expected call of NewDefaultGroup.
func (mr *MockDCGMMockRecorder) NewDefaultGroup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDefaultGroup", reflect.TypeOf((*MockDCGM)(nil).NewDefaultGroup), arg0)
}

// UpdateAllFields mocks base method.
func (m *MockDCGM) UpdateAllFields() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllFields")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAllFields indicates an expected call of UpdateAllFields.
func (mr *MockDCGMMockRecorder) UpdateAllFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllFields", reflect.TypeOf((*MockDCGM)(nil).UpdateAllFields))
}

// WatchFieldsWithGroupEx mocks base method.
func (m *MockDCGM) WatchFieldsWithGroupEx(arg0 dcgm.FieldHandle, arg1 dcgm.GroupHandle, arg2 int64, arg3 float64, arg4 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchFieldsWithGroupEx", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchFieldsWithGroupEx indicates an expected call of WatchFieldsWithGroupEx.
func (mr *MockDCGMMockRecorder) WatchFieldsWithGroupEx(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchFieldsWithGroupEx", reflect.TypeOf((*MockDCGM)(nil).WatchFieldsWithGroupEx), arg0, arg1, arg2, arg3, arg4)
}
