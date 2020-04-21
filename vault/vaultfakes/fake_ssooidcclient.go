// Code generated by counterfeiter. DO NOT EDIT.
package vaultfakes

import (
	"sync"

	"github.com/99designs/aws-vault/vault"
	"github.com/aws/aws-sdk-go/service/ssooidc"
)

type FakeSSOOIDCClient struct {
	CreateTokenStub        func(*ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error)
	createTokenMutex       sync.RWMutex
	createTokenArgsForCall []struct {
		arg1 *ssooidc.CreateTokenInput
	}
	createTokenReturns struct {
		result1 *ssooidc.CreateTokenOutput
		result2 error
	}
	createTokenReturnsOnCall map[int]struct {
		result1 *ssooidc.CreateTokenOutput
		result2 error
	}
	RegisterClientStub        func(*ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error)
	registerClientMutex       sync.RWMutex
	registerClientArgsForCall []struct {
		arg1 *ssooidc.RegisterClientInput
	}
	registerClientReturns struct {
		result1 *ssooidc.RegisterClientOutput
		result2 error
	}
	registerClientReturnsOnCall map[int]struct {
		result1 *ssooidc.RegisterClientOutput
		result2 error
	}
	StartDeviceAuthorizationStub        func(*ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error)
	startDeviceAuthorizationMutex       sync.RWMutex
	startDeviceAuthorizationArgsForCall []struct {
		arg1 *ssooidc.StartDeviceAuthorizationInput
	}
	startDeviceAuthorizationReturns struct {
		result1 *ssooidc.StartDeviceAuthorizationOutput
		result2 error
	}
	startDeviceAuthorizationReturnsOnCall map[int]struct {
		result1 *ssooidc.StartDeviceAuthorizationOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSOOIDCClient) CreateToken(arg1 *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
	fake.createTokenMutex.Lock()
	ret, specificReturn := fake.createTokenReturnsOnCall[len(fake.createTokenArgsForCall)]
	fake.createTokenArgsForCall = append(fake.createTokenArgsForCall, struct {
		arg1 *ssooidc.CreateTokenInput
	}{arg1})
	fake.recordInvocation("CreateToken", []interface{}{arg1})
	fake.createTokenMutex.Unlock()
	if fake.CreateTokenStub != nil {
		return fake.CreateTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSSOOIDCClient) CreateTokenCallCount() int {
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	return len(fake.createTokenArgsForCall)
}

func (fake *FakeSSOOIDCClient) CreateTokenCalls(stub func(*ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error)) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = stub
}

func (fake *FakeSSOOIDCClient) CreateTokenArgsForCall(i int) *ssooidc.CreateTokenInput {
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	argsForCall := fake.createTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSSOOIDCClient) CreateTokenReturns(result1 *ssooidc.CreateTokenOutput, result2 error) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = nil
	fake.createTokenReturns = struct {
		result1 *ssooidc.CreateTokenOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) CreateTokenReturnsOnCall(i int, result1 *ssooidc.CreateTokenOutput, result2 error) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = nil
	if fake.createTokenReturnsOnCall == nil {
		fake.createTokenReturnsOnCall = make(map[int]struct {
			result1 *ssooidc.CreateTokenOutput
			result2 error
		})
	}
	fake.createTokenReturnsOnCall[i] = struct {
		result1 *ssooidc.CreateTokenOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) RegisterClient(arg1 *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
	fake.registerClientMutex.Lock()
	ret, specificReturn := fake.registerClientReturnsOnCall[len(fake.registerClientArgsForCall)]
	fake.registerClientArgsForCall = append(fake.registerClientArgsForCall, struct {
		arg1 *ssooidc.RegisterClientInput
	}{arg1})
	fake.recordInvocation("RegisterClient", []interface{}{arg1})
	fake.registerClientMutex.Unlock()
	if fake.RegisterClientStub != nil {
		return fake.RegisterClientStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.registerClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSSOOIDCClient) RegisterClientCallCount() int {
	fake.registerClientMutex.RLock()
	defer fake.registerClientMutex.RUnlock()
	return len(fake.registerClientArgsForCall)
}

func (fake *FakeSSOOIDCClient) RegisterClientCalls(stub func(*ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error)) {
	fake.registerClientMutex.Lock()
	defer fake.registerClientMutex.Unlock()
	fake.RegisterClientStub = stub
}

func (fake *FakeSSOOIDCClient) RegisterClientArgsForCall(i int) *ssooidc.RegisterClientInput {
	fake.registerClientMutex.RLock()
	defer fake.registerClientMutex.RUnlock()
	argsForCall := fake.registerClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSSOOIDCClient) RegisterClientReturns(result1 *ssooidc.RegisterClientOutput, result2 error) {
	fake.registerClientMutex.Lock()
	defer fake.registerClientMutex.Unlock()
	fake.RegisterClientStub = nil
	fake.registerClientReturns = struct {
		result1 *ssooidc.RegisterClientOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) RegisterClientReturnsOnCall(i int, result1 *ssooidc.RegisterClientOutput, result2 error) {
	fake.registerClientMutex.Lock()
	defer fake.registerClientMutex.Unlock()
	fake.RegisterClientStub = nil
	if fake.registerClientReturnsOnCall == nil {
		fake.registerClientReturnsOnCall = make(map[int]struct {
			result1 *ssooidc.RegisterClientOutput
			result2 error
		})
	}
	fake.registerClientReturnsOnCall[i] = struct {
		result1 *ssooidc.RegisterClientOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorization(arg1 *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	fake.startDeviceAuthorizationMutex.Lock()
	ret, specificReturn := fake.startDeviceAuthorizationReturnsOnCall[len(fake.startDeviceAuthorizationArgsForCall)]
	fake.startDeviceAuthorizationArgsForCall = append(fake.startDeviceAuthorizationArgsForCall, struct {
		arg1 *ssooidc.StartDeviceAuthorizationInput
	}{arg1})
	fake.recordInvocation("StartDeviceAuthorization", []interface{}{arg1})
	fake.startDeviceAuthorizationMutex.Unlock()
	if fake.StartDeviceAuthorizationStub != nil {
		return fake.StartDeviceAuthorizationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.startDeviceAuthorizationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorizationCallCount() int {
	fake.startDeviceAuthorizationMutex.RLock()
	defer fake.startDeviceAuthorizationMutex.RUnlock()
	return len(fake.startDeviceAuthorizationArgsForCall)
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorizationCalls(stub func(*ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error)) {
	fake.startDeviceAuthorizationMutex.Lock()
	defer fake.startDeviceAuthorizationMutex.Unlock()
	fake.StartDeviceAuthorizationStub = stub
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorizationArgsForCall(i int) *ssooidc.StartDeviceAuthorizationInput {
	fake.startDeviceAuthorizationMutex.RLock()
	defer fake.startDeviceAuthorizationMutex.RUnlock()
	argsForCall := fake.startDeviceAuthorizationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorizationReturns(result1 *ssooidc.StartDeviceAuthorizationOutput, result2 error) {
	fake.startDeviceAuthorizationMutex.Lock()
	defer fake.startDeviceAuthorizationMutex.Unlock()
	fake.StartDeviceAuthorizationStub = nil
	fake.startDeviceAuthorizationReturns = struct {
		result1 *ssooidc.StartDeviceAuthorizationOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) StartDeviceAuthorizationReturnsOnCall(i int, result1 *ssooidc.StartDeviceAuthorizationOutput, result2 error) {
	fake.startDeviceAuthorizationMutex.Lock()
	defer fake.startDeviceAuthorizationMutex.Unlock()
	fake.StartDeviceAuthorizationStub = nil
	if fake.startDeviceAuthorizationReturnsOnCall == nil {
		fake.startDeviceAuthorizationReturnsOnCall = make(map[int]struct {
			result1 *ssooidc.StartDeviceAuthorizationOutput
			result2 error
		})
	}
	fake.startDeviceAuthorizationReturnsOnCall[i] = struct {
		result1 *ssooidc.StartDeviceAuthorizationOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSSOOIDCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	fake.registerClientMutex.RLock()
	defer fake.registerClientMutex.RUnlock()
	fake.startDeviceAuthorizationMutex.RLock()
	defer fake.startDeviceAuthorizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSSOOIDCClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ vault.SSOOIDCClient = new(FakeSSOOIDCClient)
