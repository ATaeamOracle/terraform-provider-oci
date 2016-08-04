package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlsdk.Resource), args.Error(1)
}

func (m *MockClient) GetUser(id string) (*baremtlsdk.Resource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateUser(id, description string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteUser(id string, opts ...baremtlsdk.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreateGroup(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlsdk.Resource), args.Error(1)
}

func (m *MockClient) GetGroup(id string) (*baremtlsdk.Resource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateGroup(id, description string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteGroup(id string, opts ...baremtlsdk.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreateCompartment(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlsdk.Resource), args.Error(1)
}

func (m *MockClient) GetCompartment(id string) (*baremtlsdk.Resource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateCompartment(id, description string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}
