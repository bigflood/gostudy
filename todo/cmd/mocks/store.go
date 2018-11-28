// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	sync "sync"

	cmd "github.com/bigflood/gostudy/todo/cmd"
	store "github.com/bigflood/gostudy/todo/store"
)

type Store struct {
	AddStub        func(string) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	DoneStub        func(int) error
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
		arg1 int
	}
	doneReturns struct {
		result1 error
	}
	doneReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(store.Filter) ([]store.Task, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 store.Filter
	}
	listReturns struct {
		result1 []store.Task
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []store.Task
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) Add(arg1 string) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addReturns
	return fakeReturns.result1
}

func (fake *Store) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Store) AddCalls(stub func(string) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *Store) AddArgsForCall(i int) string {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Store) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) Done(arg1 int) error {
	fake.doneMutex.Lock()
	ret, specificReturn := fake.doneReturnsOnCall[len(fake.doneArgsForCall)]
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Done", []interface{}{arg1})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		return fake.DoneStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doneReturns
	return fakeReturns.result1
}

func (fake *Store) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *Store) DoneCalls(stub func(int) error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *Store) DoneArgsForCall(i int) int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	argsForCall := fake.doneArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Store) DoneReturns(result1 error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	fake.doneReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) DoneReturnsOnCall(i int, result1 error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	if fake.doneReturnsOnCall == nil {
		fake.doneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) List(arg1 store.Filter) ([]store.Task, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 store.Filter
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Store) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *Store) ListCalls(stub func(store.Filter) ([]store.Task, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *Store) ListArgsForCall(i int) store.Filter {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Store) ListReturns(result1 []store.Task, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []store.Task
		result2 error
	}{result1, result2}
}

func (fake *Store) ListReturnsOnCall(i int, result1 []store.Task, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []store.Task
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []store.Task
		result2 error
	}{result1, result2}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Store) recordInvocation(key string, args []interface{}) {
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

var _ cmd.Store = new(Store)
