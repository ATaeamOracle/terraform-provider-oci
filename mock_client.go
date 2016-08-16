package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremetal.IdentityResource), args.Error(1)
}

func (m *MockClient) GetUser(id string) (*baremetal.IdentityResource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateUser(id, description string, opts ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteUser(id string, opts ...baremetal.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreateGroup(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremetal.IdentityResource), args.Error(1)
}

func (m *MockClient) GetGroup(id string) (*baremetal.IdentityResource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateGroup(id, description string, opts ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteGroup(id string, opts ...baremetal.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreatePolicy(name, description string, statements []string, options ...baremetal.Options) (*baremetal.Policy, error) {
	args := m.Called(name, description, statements)
	return args.Get(0).(*baremetal.Policy), args.Error(1)
}

func (m *MockClient) GetPolicy(id string) (*baremetal.Policy, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremetal.Policy)
	return u, args.Error(1)
}

func (m *MockClient) UpdatePolicy(id, description string, statements []string, opts ...baremetal.Options) (*baremetal.Policy, error) {
	args := m.Called(id, description, statements)
	u, _ := args.Get(0).(*baremetal.Policy)
	return u, args.Error(1)
}

func (m *MockClient) DeletePolicy(id string, opts ...baremetal.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreateCompartment(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremetal.IdentityResource), args.Error(1)
}

func (m *MockClient) GetCompartment(id string) (*baremetal.IdentityResource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateCompartment(id, description string, opts ...baremetal.Options) (*baremetal.IdentityResource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremetal.IdentityResource)
	return u, args.Error(1)
}

func (m *MockClient) ListShapes(compartmentID string, opt ...baremetal.CoreOptions) (*baremetal.ShapeList, error) {
	args := m.Called(compartmentID, opt)
	u, _ := args.Get(0).(*baremetal.ShapeList)
	return u, args.Error(1)

}

func (m *MockClient) ListVnicAttachments(compartmentID string, opt ...baremetal.CoreOptions) (*baremetal.VnicAttachmentList, error) {
	args := m.Called(compartmentID, opt)
	u, _ := args.Get(0).(*baremetal.VnicAttachmentList)
	return u, args.Error(1)
}
