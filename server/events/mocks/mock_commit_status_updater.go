// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommitStatusUpdater)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockCommitStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommitStatusUpdater(options ...pegomock.Option) *MockCommitStatusUpdater {
	mock := &MockCommitStatusUpdater{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommitStatusUpdater) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommitStatusUpdater) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommitStatusUpdater) UpdateCombined(repo models.Repo, pull models.PullRequest, status models.CommitStatus, cmdName command.Name) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{repo, pull, status, cmdName}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateCombined", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdateCombinedCount(repo models.Repo, pull models.PullRequest, status models.CommitStatus, cmdName command.Name, numSuccess int, numTotal int) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{repo, pull, status, cmdName, numSuccess, numTotal}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateCombinedCount", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdatePostWorkflowHook(pull models.PullRequest, status models.CommitStatus, hookDescription string, runtimeDescription string, url string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{pull, status, hookDescription, runtimeDescription, url}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdatePostWorkflowHook", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdatePreWorkflowHook(pull models.PullRequest, status models.CommitStatus, hookDescription string, runtimeDescription string, url string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{pull, status, hookDescription, runtimeDescription, url}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdatePreWorkflowHook", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledOnce() *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommitStatusUpdater struct {
	mock                   *MockCommitStatusUpdater
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommitStatusUpdater) UpdateCombined(repo models.Repo, pull models.PullRequest, status models.CommitStatus, cmdName command.Name) *MockCommitStatusUpdater_UpdateCombined_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, cmdName}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateCombined", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdateCombined_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdateCombined_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdateCombined_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, command.Name) {
	repo, pull, status, cmdName := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], cmdName[len(cmdName)-1]
}

func (c *MockCommitStatusUpdater_UpdateCombined_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []command.Name) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(command.Name)
		}
	}
	return
}

func (verifier *VerifierMockCommitStatusUpdater) UpdateCombinedCount(repo models.Repo, pull models.PullRequest, status models.CommitStatus, cmdName command.Name, numSuccess int, numTotal int) *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, cmdName, numSuccess, numTotal}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateCombinedCount", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, command.Name, int, int) {
	repo, pull, status, cmdName, numSuccess, numTotal := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], cmdName[len(cmdName)-1], numSuccess[len(numSuccess)-1], numTotal[len(numTotal)-1]
}

func (c *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []command.Name, _param4 []int, _param5 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(command.Name)
		}
		_param4 = make([]int, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(int)
		}
		_param5 = make([]int, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockCommitStatusUpdater) UpdatePostWorkflowHook(pull models.PullRequest, status models.CommitStatus, hookDescription string, runtimeDescription string, url string) *MockCommitStatusUpdater_UpdatePostWorkflowHook_OngoingVerification {
	params := []pegomock.Param{pull, status, hookDescription, runtimeDescription, url}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdatePostWorkflowHook", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdatePostWorkflowHook_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdatePostWorkflowHook_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdatePostWorkflowHook_OngoingVerification) GetCapturedArguments() (models.PullRequest, models.CommitStatus, string, string, string) {
	pull, status, hookDescription, runtimeDescription, url := c.GetAllCapturedArguments()
	return pull[len(pull)-1], status[len(status)-1], hookDescription[len(hookDescription)-1], runtimeDescription[len(runtimeDescription)-1], url[len(url)-1]
}

func (c *MockCommitStatusUpdater_UpdatePostWorkflowHook_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []models.CommitStatus, _param2 []string, _param3 []string, _param4 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.CommitStatus)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCommitStatusUpdater) UpdatePreWorkflowHook(pull models.PullRequest, status models.CommitStatus, hookDescription string, runtimeDescription string, url string) *MockCommitStatusUpdater_UpdatePreWorkflowHook_OngoingVerification {
	params := []pegomock.Param{pull, status, hookDescription, runtimeDescription, url}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdatePreWorkflowHook", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdatePreWorkflowHook_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdatePreWorkflowHook_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdatePreWorkflowHook_OngoingVerification) GetCapturedArguments() (models.PullRequest, models.CommitStatus, string, string, string) {
	pull, status, hookDescription, runtimeDescription, url := c.GetAllCapturedArguments()
	return pull[len(pull)-1], status[len(status)-1], hookDescription[len(hookDescription)-1], runtimeDescription[len(runtimeDescription)-1], url[len(url)-1]
}

func (c *MockCommitStatusUpdater_UpdatePreWorkflowHook_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []models.CommitStatus, _param2 []string, _param3 []string, _param4 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.CommitStatus)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
	}
	return
}
