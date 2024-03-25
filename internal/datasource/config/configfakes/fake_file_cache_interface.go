// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/nginx/agent/v3/api/grpc/instances"
	"github.com/nginx/agent/v3/internal/datasource/config"
)

type FakeFileCacheInterface struct {
	CacheContentStub        func() map[string]*instances.File
	cacheContentMutex       sync.RWMutex
	cacheContentArgsForCall []struct {
	}
	cacheContentReturns struct {
		result1 map[string]*instances.File
	}
	cacheContentReturnsOnCall map[int]struct {
		result1 map[string]*instances.File
	}
	ReadFileCacheStub        func() (map[string]*instances.File, error)
	readFileCacheMutex       sync.RWMutex
	readFileCacheArgsForCall []struct {
	}
	readFileCacheReturns struct {
		result1 map[string]*instances.File
		result2 error
	}
	readFileCacheReturnsOnCall map[int]struct {
		result1 map[string]*instances.File
		result2 error
	}
	SetCachePathStub        func(string)
	setCachePathMutex       sync.RWMutex
	setCachePathArgsForCall []struct {
		arg1 string
	}
	UpdateFileCacheStub        func(map[string]*instances.File) error
	updateFileCacheMutex       sync.RWMutex
	updateFileCacheArgsForCall []struct {
		arg1 map[string]*instances.File
	}
	updateFileCacheReturns struct {
		result1 error
	}
	updateFileCacheReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileCacheInterface) CacheContent() map[string]*instances.File {
	fake.cacheContentMutex.Lock()
	ret, specificReturn := fake.cacheContentReturnsOnCall[len(fake.cacheContentArgsForCall)]
	fake.cacheContentArgsForCall = append(fake.cacheContentArgsForCall, struct {
	}{})
	stub := fake.CacheContentStub
	fakeReturns := fake.cacheContentReturns
	fake.recordInvocation("CacheContent", []interface{}{})
	fake.cacheContentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileCacheInterface) CacheContentCallCount() int {
	fake.cacheContentMutex.RLock()
	defer fake.cacheContentMutex.RUnlock()
	return len(fake.cacheContentArgsForCall)
}

func (fake *FakeFileCacheInterface) CacheContentCalls(stub func() map[string]*instances.File) {
	fake.cacheContentMutex.Lock()
	defer fake.cacheContentMutex.Unlock()
	fake.CacheContentStub = stub
}

func (fake *FakeFileCacheInterface) CacheContentReturns(result1 map[string]*instances.File) {
	fake.cacheContentMutex.Lock()
	defer fake.cacheContentMutex.Unlock()
	fake.CacheContentStub = nil
	fake.cacheContentReturns = struct {
		result1 map[string]*instances.File
	}{result1}
}

func (fake *FakeFileCacheInterface) CacheContentReturnsOnCall(i int, result1 map[string]*instances.File) {
	fake.cacheContentMutex.Lock()
	defer fake.cacheContentMutex.Unlock()
	fake.CacheContentStub = nil
	if fake.cacheContentReturnsOnCall == nil {
		fake.cacheContentReturnsOnCall = make(map[int]struct {
			result1 map[string]*instances.File
		})
	}
	fake.cacheContentReturnsOnCall[i] = struct {
		result1 map[string]*instances.File
	}{result1}
}

func (fake *FakeFileCacheInterface) ReadFileCache() (map[string]*instances.File, error) {
	fake.readFileCacheMutex.Lock()
	ret, specificReturn := fake.readFileCacheReturnsOnCall[len(fake.readFileCacheArgsForCall)]
	fake.readFileCacheArgsForCall = append(fake.readFileCacheArgsForCall, struct {
	}{})
	stub := fake.ReadFileCacheStub
	fakeReturns := fake.readFileCacheReturns
	fake.recordInvocation("ReadFileCache", []interface{}{})
	fake.readFileCacheMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFileCacheInterface) ReadFileCacheCallCount() int {
	fake.readFileCacheMutex.RLock()
	defer fake.readFileCacheMutex.RUnlock()
	return len(fake.readFileCacheArgsForCall)
}

func (fake *FakeFileCacheInterface) ReadFileCacheCalls(stub func() (map[string]*instances.File, error)) {
	fake.readFileCacheMutex.Lock()
	defer fake.readFileCacheMutex.Unlock()
	fake.ReadFileCacheStub = stub
}

func (fake *FakeFileCacheInterface) ReadFileCacheReturns(result1 map[string]*instances.File, result2 error) {
	fake.readFileCacheMutex.Lock()
	defer fake.readFileCacheMutex.Unlock()
	fake.ReadFileCacheStub = nil
	fake.readFileCacheReturns = struct {
		result1 map[string]*instances.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileCacheInterface) ReadFileCacheReturnsOnCall(i int, result1 map[string]*instances.File, result2 error) {
	fake.readFileCacheMutex.Lock()
	defer fake.readFileCacheMutex.Unlock()
	fake.ReadFileCacheStub = nil
	if fake.readFileCacheReturnsOnCall == nil {
		fake.readFileCacheReturnsOnCall = make(map[int]struct {
			result1 map[string]*instances.File
			result2 error
		})
	}
	fake.readFileCacheReturnsOnCall[i] = struct {
		result1 map[string]*instances.File
		result2 error
	}{result1, result2}
}

func (fake *FakeFileCacheInterface) SetCachePath(arg1 string) {
	fake.setCachePathMutex.Lock()
	fake.setCachePathArgsForCall = append(fake.setCachePathArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetCachePathStub
	fake.recordInvocation("SetCachePath", []interface{}{arg1})
	fake.setCachePathMutex.Unlock()
	if stub != nil {
		fake.SetCachePathStub(arg1)
	}
}

func (fake *FakeFileCacheInterface) SetCachePathCallCount() int {
	fake.setCachePathMutex.RLock()
	defer fake.setCachePathMutex.RUnlock()
	return len(fake.setCachePathArgsForCall)
}

func (fake *FakeFileCacheInterface) SetCachePathCalls(stub func(string)) {
	fake.setCachePathMutex.Lock()
	defer fake.setCachePathMutex.Unlock()
	fake.SetCachePathStub = stub
}

func (fake *FakeFileCacheInterface) SetCachePathArgsForCall(i int) string {
	fake.setCachePathMutex.RLock()
	defer fake.setCachePathMutex.RUnlock()
	argsForCall := fake.setCachePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileCacheInterface) UpdateFileCache(arg1 map[string]*instances.File) error {
	fake.updateFileCacheMutex.Lock()
	ret, specificReturn := fake.updateFileCacheReturnsOnCall[len(fake.updateFileCacheArgsForCall)]
	fake.updateFileCacheArgsForCall = append(fake.updateFileCacheArgsForCall, struct {
		arg1 map[string]*instances.File
	}{arg1})
	stub := fake.UpdateFileCacheStub
	fakeReturns := fake.updateFileCacheReturns
	fake.recordInvocation("UpdateFileCache", []interface{}{arg1})
	fake.updateFileCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFileCacheInterface) UpdateFileCacheCallCount() int {
	fake.updateFileCacheMutex.RLock()
	defer fake.updateFileCacheMutex.RUnlock()
	return len(fake.updateFileCacheArgsForCall)
}

func (fake *FakeFileCacheInterface) UpdateFileCacheCalls(stub func(map[string]*instances.File) error) {
	fake.updateFileCacheMutex.Lock()
	defer fake.updateFileCacheMutex.Unlock()
	fake.UpdateFileCacheStub = stub
}

func (fake *FakeFileCacheInterface) UpdateFileCacheArgsForCall(i int) map[string]*instances.File {
	fake.updateFileCacheMutex.RLock()
	defer fake.updateFileCacheMutex.RUnlock()
	argsForCall := fake.updateFileCacheArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFileCacheInterface) UpdateFileCacheReturns(result1 error) {
	fake.updateFileCacheMutex.Lock()
	defer fake.updateFileCacheMutex.Unlock()
	fake.UpdateFileCacheStub = nil
	fake.updateFileCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileCacheInterface) UpdateFileCacheReturnsOnCall(i int, result1 error) {
	fake.updateFileCacheMutex.Lock()
	defer fake.updateFileCacheMutex.Unlock()
	fake.UpdateFileCacheStub = nil
	if fake.updateFileCacheReturnsOnCall == nil {
		fake.updateFileCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateFileCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileCacheInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cacheContentMutex.RLock()
	defer fake.cacheContentMutex.RUnlock()
	fake.readFileCacheMutex.RLock()
	defer fake.readFileCacheMutex.RUnlock()
	fake.setCachePathMutex.RLock()
	defer fake.setCachePathMutex.RUnlock()
	fake.updateFileCacheMutex.RLock()
	defer fake.updateFileCacheMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFileCacheInterface) recordInvocation(key string, args []interface{}) {
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

var _ config.FileCacheInterface = new(FakeFileCacheInterface)