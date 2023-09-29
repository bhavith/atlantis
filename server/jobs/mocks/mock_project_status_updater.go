// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/jobs (interfaces: ProjectStatusUpdater)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockProjectStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectStatusUpdater(options ...pegomock.Option) *MockProjectStatusUpdater {
	mock := &MockProjectStatusUpdater{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectStatusUpdater) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectStatusUpdater) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectStatusUpdater) UpdateProject(ctx command.ProjectContext, cmdName command.Name, status models.CommitStatus, url string, res *command.ProjectResult) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectStatusUpdater().")
	}
	params := []pegomock.Param{ctx, cmdName, status, url, res}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProject", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledOnce() *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectStatusUpdater struct {
	mock                   *MockProjectStatusUpdater
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectStatusUpdater) UpdateProject(ctx command.ProjectContext, cmdName command.Name, status models.CommitStatus, url string, res *command.ProjectResult) *MockProjectStatusUpdater_UpdateProject_OngoingVerification {
	params := []pegomock.Param{ctx, cmdName, status, url, res}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProject", params, verifier.timeout)
	return &MockProjectStatusUpdater_UpdateProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectStatusUpdater_UpdateProject_OngoingVerification struct {
	mock              *MockProjectStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectStatusUpdater_UpdateProject_OngoingVerification) GetCapturedArguments() (command.ProjectContext, command.Name, models.CommitStatus, string, *command.ProjectResult) {
	ctx, cmdName, status, url, res := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], cmdName[len(cmdName)-1], status[len(status)-1], url[len(url)-1], res[len(res)-1]
}

func (c *MockProjectStatusUpdater_UpdateProject_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []command.Name, _param2 []models.CommitStatus, _param3 []string, _param4 []*command.ProjectResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.ProjectContext)
		}
		_param1 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(command.Name)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]*command.ProjectResult, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(*command.ProjectResult)
		}
	}
	return
}
