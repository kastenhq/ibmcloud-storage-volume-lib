// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.ibm.com/alchemy-containers/armada-cluster/iaas/softlayer/backend"
)

type VirtualGuestService struct {
	FilterStub        func(string) backend.ServerService
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 string
	}
	filterReturns struct {
		result1 backend.ServerService
	}
	filterReturnsOnCall map[int]struct {
		result1 backend.ServerService
	}
	MaskStub        func(string) backend.ServerService
	maskMutex       sync.RWMutex
	maskArgsForCall []struct {
		arg1 string
	}
	maskReturns struct {
		result1 backend.ServerService
	}
	maskReturnsOnCall map[int]struct {
		result1 backend.ServerService
	}
	IDStub        func(int) backend.ServerService
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
		arg1 int
	}
	iDReturns struct {
		result1 backend.ServerService
	}
	iDReturnsOnCall map[int]struct {
		result1 backend.ServerService
	}
	CreateObjectStub        func(templateObject backend.Server) (resp backend.Server, err error)
	createObjectMutex       sync.RWMutex
	createObjectArgsForCall []struct {
		templateObject backend.Server
	}
	createObjectReturns struct {
		result1 backend.Server
		result2 error
	}
	createObjectReturnsOnCall map[int]struct {
		result1 backend.Server
		result2 error
	}
	GetObjectStub        func() (backend.Server, error)
	getObjectMutex       sync.RWMutex
	getObjectArgsForCall []struct{}
	getObjectReturns     struct {
		result1 backend.Server
		result2 error
	}
	getObjectReturnsOnCall map[int]struct {
		result1 backend.Server
		result2 error
	}
	DeleteObjectStub        func() (bool, error)
	deleteObjectMutex       sync.RWMutex
	deleteObjectArgsForCall []struct{}
	deleteObjectReturns     struct {
		result1 bool
		result2 error
	}
	deleteObjectReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetActiveTransactionsStub        func() ([]datatypes.Provisioning_Version1_Transaction, error)
	getActiveTransactionsMutex       sync.RWMutex
	getActiveTransactionsArgsForCall []struct{}
	getActiveTransactionsReturns     struct {
		result1 []datatypes.Provisioning_Version1_Transaction
		result2 error
	}
	getActiveTransactionsReturnsOnCall map[int]struct {
		result1 []datatypes.Provisioning_Version1_Transaction
		result2 error
	}
	RebootHardStub        func() (bool, error)
	rebootHardMutex       sync.RWMutex
	rebootHardArgsForCall []struct{}
	rebootHardReturns     struct {
		result1 bool
		result2 error
	}
	rebootHardReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RebootSoftStub        func() (bool, error)
	rebootSoftMutex       sync.RWMutex
	rebootSoftArgsForCall []struct{}
	rebootSoftReturns     struct {
		result1 bool
		result2 error
	}
	rebootSoftReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	PowerOnStub        func() (bool, error)
	powerOnMutex       sync.RWMutex
	powerOnArgsForCall []struct{}
	powerOnReturns     struct {
		result1 bool
		result2 error
	}
	powerOnReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ReloadOperatingSystemStub        func(token *string, config *datatypes.Container_Hardware_Server_Configuration) (string, error)
	reloadOperatingSystemMutex       sync.RWMutex
	reloadOperatingSystemArgsForCall []struct {
		token  *string
		config *datatypes.Container_Hardware_Server_Configuration
	}
	reloadOperatingSystemReturns struct {
		result1 string
		result2 error
	}
	reloadOperatingSystemReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SetTagsStub        func(tags *string) (bool, error)
	setTagsMutex       sync.RWMutex
	setTagsArgsForCall []struct {
		tags *string
	}
	setTagsReturns struct {
		result1 bool
		result2 error
	}
	setTagsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetUserMetadataStub        func(metadata string) (bool, error)
	setUserMetadataMutex       sync.RWMutex
	setUserMetadataArgsForCall []struct {
		metadata string
	}
	setUserMetadataReturns struct {
		result1 bool
		result2 error
	}
	setUserMetadataReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VirtualGuestService) Filter(arg1 string) backend.ServerService {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Filter", []interface{}{arg1})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.filterReturns.result1
}

func (fake *VirtualGuestService) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *VirtualGuestService) FilterArgsForCall(i int) string {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return fake.filterArgsForCall[i].arg1
}

func (fake *VirtualGuestService) FilterReturns(result1 backend.ServerService) {
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) FilterReturnsOnCall(i int, result1 backend.ServerService) {
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 backend.ServerService
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) Mask(arg1 string) backend.ServerService {
	fake.maskMutex.Lock()
	ret, specificReturn := fake.maskReturnsOnCall[len(fake.maskArgsForCall)]
	fake.maskArgsForCall = append(fake.maskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Mask", []interface{}{arg1})
	fake.maskMutex.Unlock()
	if fake.MaskStub != nil {
		return fake.MaskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.maskReturns.result1
}

func (fake *VirtualGuestService) MaskCallCount() int {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return len(fake.maskArgsForCall)
}

func (fake *VirtualGuestService) MaskArgsForCall(i int) string {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return fake.maskArgsForCall[i].arg1
}

func (fake *VirtualGuestService) MaskReturns(result1 backend.ServerService) {
	fake.MaskStub = nil
	fake.maskReturns = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) MaskReturnsOnCall(i int, result1 backend.ServerService) {
	fake.MaskStub = nil
	if fake.maskReturnsOnCall == nil {
		fake.maskReturnsOnCall = make(map[int]struct {
			result1 backend.ServerService
		})
	}
	fake.maskReturnsOnCall[i] = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) ID(arg1 int) backend.ServerService {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ID", []interface{}{arg1})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *VirtualGuestService) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *VirtualGuestService) IDArgsForCall(i int) int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return fake.iDArgsForCall[i].arg1
}

func (fake *VirtualGuestService) IDReturns(result1 backend.ServerService) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) IDReturnsOnCall(i int, result1 backend.ServerService) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 backend.ServerService
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 backend.ServerService
	}{result1}
}

func (fake *VirtualGuestService) CreateObject(templateObject backend.Server) (resp backend.Server, err error) {
	fake.createObjectMutex.Lock()
	ret, specificReturn := fake.createObjectReturnsOnCall[len(fake.createObjectArgsForCall)]
	fake.createObjectArgsForCall = append(fake.createObjectArgsForCall, struct {
		templateObject backend.Server
	}{templateObject})
	fake.recordInvocation("CreateObject", []interface{}{templateObject})
	fake.createObjectMutex.Unlock()
	if fake.CreateObjectStub != nil {
		return fake.CreateObjectStub(templateObject)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createObjectReturns.result1, fake.createObjectReturns.result2
}

func (fake *VirtualGuestService) CreateObjectCallCount() int {
	fake.createObjectMutex.RLock()
	defer fake.createObjectMutex.RUnlock()
	return len(fake.createObjectArgsForCall)
}

func (fake *VirtualGuestService) CreateObjectArgsForCall(i int) backend.Server {
	fake.createObjectMutex.RLock()
	defer fake.createObjectMutex.RUnlock()
	return fake.createObjectArgsForCall[i].templateObject
}

func (fake *VirtualGuestService) CreateObjectReturns(result1 backend.Server, result2 error) {
	fake.CreateObjectStub = nil
	fake.createObjectReturns = struct {
		result1 backend.Server
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) CreateObjectReturnsOnCall(i int, result1 backend.Server, result2 error) {
	fake.CreateObjectStub = nil
	if fake.createObjectReturnsOnCall == nil {
		fake.createObjectReturnsOnCall = make(map[int]struct {
			result1 backend.Server
			result2 error
		})
	}
	fake.createObjectReturnsOnCall[i] = struct {
		result1 backend.Server
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) GetObject() (backend.Server, error) {
	fake.getObjectMutex.Lock()
	ret, specificReturn := fake.getObjectReturnsOnCall[len(fake.getObjectArgsForCall)]
	fake.getObjectArgsForCall = append(fake.getObjectArgsForCall, struct{}{})
	fake.recordInvocation("GetObject", []interface{}{})
	fake.getObjectMutex.Unlock()
	if fake.GetObjectStub != nil {
		return fake.GetObjectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getObjectReturns.result1, fake.getObjectReturns.result2
}

func (fake *VirtualGuestService) GetObjectCallCount() int {
	fake.getObjectMutex.RLock()
	defer fake.getObjectMutex.RUnlock()
	return len(fake.getObjectArgsForCall)
}

func (fake *VirtualGuestService) GetObjectReturns(result1 backend.Server, result2 error) {
	fake.GetObjectStub = nil
	fake.getObjectReturns = struct {
		result1 backend.Server
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) GetObjectReturnsOnCall(i int, result1 backend.Server, result2 error) {
	fake.GetObjectStub = nil
	if fake.getObjectReturnsOnCall == nil {
		fake.getObjectReturnsOnCall = make(map[int]struct {
			result1 backend.Server
			result2 error
		})
	}
	fake.getObjectReturnsOnCall[i] = struct {
		result1 backend.Server
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) DeleteObject() (bool, error) {
	fake.deleteObjectMutex.Lock()
	ret, specificReturn := fake.deleteObjectReturnsOnCall[len(fake.deleteObjectArgsForCall)]
	fake.deleteObjectArgsForCall = append(fake.deleteObjectArgsForCall, struct{}{})
	fake.recordInvocation("DeleteObject", []interface{}{})
	fake.deleteObjectMutex.Unlock()
	if fake.DeleteObjectStub != nil {
		return fake.DeleteObjectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteObjectReturns.result1, fake.deleteObjectReturns.result2
}

func (fake *VirtualGuestService) DeleteObjectCallCount() int {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return len(fake.deleteObjectArgsForCall)
}

func (fake *VirtualGuestService) DeleteObjectReturns(result1 bool, result2 error) {
	fake.DeleteObjectStub = nil
	fake.deleteObjectReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) DeleteObjectReturnsOnCall(i int, result1 bool, result2 error) {
	fake.DeleteObjectStub = nil
	if fake.deleteObjectReturnsOnCall == nil {
		fake.deleteObjectReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.deleteObjectReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) GetActiveTransactions() ([]datatypes.Provisioning_Version1_Transaction, error) {
	fake.getActiveTransactionsMutex.Lock()
	ret, specificReturn := fake.getActiveTransactionsReturnsOnCall[len(fake.getActiveTransactionsArgsForCall)]
	fake.getActiveTransactionsArgsForCall = append(fake.getActiveTransactionsArgsForCall, struct{}{})
	fake.recordInvocation("GetActiveTransactions", []interface{}{})
	fake.getActiveTransactionsMutex.Unlock()
	if fake.GetActiveTransactionsStub != nil {
		return fake.GetActiveTransactionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getActiveTransactionsReturns.result1, fake.getActiveTransactionsReturns.result2
}

func (fake *VirtualGuestService) GetActiveTransactionsCallCount() int {
	fake.getActiveTransactionsMutex.RLock()
	defer fake.getActiveTransactionsMutex.RUnlock()
	return len(fake.getActiveTransactionsArgsForCall)
}

func (fake *VirtualGuestService) GetActiveTransactionsReturns(result1 []datatypes.Provisioning_Version1_Transaction, result2 error) {
	fake.GetActiveTransactionsStub = nil
	fake.getActiveTransactionsReturns = struct {
		result1 []datatypes.Provisioning_Version1_Transaction
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) GetActiveTransactionsReturnsOnCall(i int, result1 []datatypes.Provisioning_Version1_Transaction, result2 error) {
	fake.GetActiveTransactionsStub = nil
	if fake.getActiveTransactionsReturnsOnCall == nil {
		fake.getActiveTransactionsReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Provisioning_Version1_Transaction
			result2 error
		})
	}
	fake.getActiveTransactionsReturnsOnCall[i] = struct {
		result1 []datatypes.Provisioning_Version1_Transaction
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) RebootHard() (bool, error) {
	fake.rebootHardMutex.Lock()
	ret, specificReturn := fake.rebootHardReturnsOnCall[len(fake.rebootHardArgsForCall)]
	fake.rebootHardArgsForCall = append(fake.rebootHardArgsForCall, struct{}{})
	fake.recordInvocation("RebootHard", []interface{}{})
	fake.rebootHardMutex.Unlock()
	if fake.RebootHardStub != nil {
		return fake.RebootHardStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.rebootHardReturns.result1, fake.rebootHardReturns.result2
}

func (fake *VirtualGuestService) RebootHardCallCount() int {
	fake.rebootHardMutex.RLock()
	defer fake.rebootHardMutex.RUnlock()
	return len(fake.rebootHardArgsForCall)
}

func (fake *VirtualGuestService) RebootHardReturns(result1 bool, result2 error) {
	fake.RebootHardStub = nil
	fake.rebootHardReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) RebootHardReturnsOnCall(i int, result1 bool, result2 error) {
	fake.RebootHardStub = nil
	if fake.rebootHardReturnsOnCall == nil {
		fake.rebootHardReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.rebootHardReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) RebootSoft() (bool, error) {
	fake.rebootSoftMutex.Lock()
	ret, specificReturn := fake.rebootSoftReturnsOnCall[len(fake.rebootSoftArgsForCall)]
	fake.rebootSoftArgsForCall = append(fake.rebootSoftArgsForCall, struct{}{})
	fake.recordInvocation("RebootSoft", []interface{}{})
	fake.rebootSoftMutex.Unlock()
	if fake.RebootSoftStub != nil {
		return fake.RebootSoftStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.rebootSoftReturns.result1, fake.rebootSoftReturns.result2
}

func (fake *VirtualGuestService) RebootSoftCallCount() int {
	fake.rebootSoftMutex.RLock()
	defer fake.rebootSoftMutex.RUnlock()
	return len(fake.rebootSoftArgsForCall)
}

func (fake *VirtualGuestService) RebootSoftReturns(result1 bool, result2 error) {
	fake.RebootSoftStub = nil
	fake.rebootSoftReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) RebootSoftReturnsOnCall(i int, result1 bool, result2 error) {
	fake.RebootSoftStub = nil
	if fake.rebootSoftReturnsOnCall == nil {
		fake.rebootSoftReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.rebootSoftReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) PowerOn() (bool, error) {
	fake.powerOnMutex.Lock()
	ret, specificReturn := fake.powerOnReturnsOnCall[len(fake.powerOnArgsForCall)]
	fake.powerOnArgsForCall = append(fake.powerOnArgsForCall, struct{}{})
	fake.recordInvocation("PowerOn", []interface{}{})
	fake.powerOnMutex.Unlock()
	if fake.PowerOnStub != nil {
		return fake.PowerOnStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.powerOnReturns.result1, fake.powerOnReturns.result2
}

func (fake *VirtualGuestService) PowerOnCallCount() int {
	fake.powerOnMutex.RLock()
	defer fake.powerOnMutex.RUnlock()
	return len(fake.powerOnArgsForCall)
}

func (fake *VirtualGuestService) PowerOnReturns(result1 bool, result2 error) {
	fake.PowerOnStub = nil
	fake.powerOnReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) PowerOnReturnsOnCall(i int, result1 bool, result2 error) {
	fake.PowerOnStub = nil
	if fake.powerOnReturnsOnCall == nil {
		fake.powerOnReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.powerOnReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (string, error) {
	fake.reloadOperatingSystemMutex.Lock()
	ret, specificReturn := fake.reloadOperatingSystemReturnsOnCall[len(fake.reloadOperatingSystemArgsForCall)]
	fake.reloadOperatingSystemArgsForCall = append(fake.reloadOperatingSystemArgsForCall, struct {
		token  *string
		config *datatypes.Container_Hardware_Server_Configuration
	}{token, config})
	fake.recordInvocation("ReloadOperatingSystem", []interface{}{token, config})
	fake.reloadOperatingSystemMutex.Unlock()
	if fake.ReloadOperatingSystemStub != nil {
		return fake.ReloadOperatingSystemStub(token, config)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.reloadOperatingSystemReturns.result1, fake.reloadOperatingSystemReturns.result2
}

func (fake *VirtualGuestService) ReloadOperatingSystemCallCount() int {
	fake.reloadOperatingSystemMutex.RLock()
	defer fake.reloadOperatingSystemMutex.RUnlock()
	return len(fake.reloadOperatingSystemArgsForCall)
}

func (fake *VirtualGuestService) ReloadOperatingSystemArgsForCall(i int) (*string, *datatypes.Container_Hardware_Server_Configuration) {
	fake.reloadOperatingSystemMutex.RLock()
	defer fake.reloadOperatingSystemMutex.RUnlock()
	return fake.reloadOperatingSystemArgsForCall[i].token, fake.reloadOperatingSystemArgsForCall[i].config
}

func (fake *VirtualGuestService) ReloadOperatingSystemReturns(result1 string, result2 error) {
	fake.ReloadOperatingSystemStub = nil
	fake.reloadOperatingSystemReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) ReloadOperatingSystemReturnsOnCall(i int, result1 string, result2 error) {
	fake.ReloadOperatingSystemStub = nil
	if fake.reloadOperatingSystemReturnsOnCall == nil {
		fake.reloadOperatingSystemReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.reloadOperatingSystemReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) SetTags(tags *string) (bool, error) {
	fake.setTagsMutex.Lock()
	ret, specificReturn := fake.setTagsReturnsOnCall[len(fake.setTagsArgsForCall)]
	fake.setTagsArgsForCall = append(fake.setTagsArgsForCall, struct {
		tags *string
	}{tags})
	fake.recordInvocation("SetTags", []interface{}{tags})
	fake.setTagsMutex.Unlock()
	if fake.SetTagsStub != nil {
		return fake.SetTagsStub(tags)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setTagsReturns.result1, fake.setTagsReturns.result2
}

func (fake *VirtualGuestService) SetTagsCallCount() int {
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	return len(fake.setTagsArgsForCall)
}

func (fake *VirtualGuestService) SetTagsArgsForCall(i int) *string {
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	return fake.setTagsArgsForCall[i].tags
}

func (fake *VirtualGuestService) SetTagsReturns(result1 bool, result2 error) {
	fake.SetTagsStub = nil
	fake.setTagsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) SetTagsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.SetTagsStub = nil
	if fake.setTagsReturnsOnCall == nil {
		fake.setTagsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.setTagsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) SetUserMetadata(metadata string) (bool, error) {
	fake.setUserMetadataMutex.Lock()
	ret, specificReturn := fake.setUserMetadataReturnsOnCall[len(fake.setUserMetadataArgsForCall)]
	fake.setUserMetadataArgsForCall = append(fake.setUserMetadataArgsForCall, struct {
		metadata string
	}{metadata})
	fake.recordInvocation("SetUserMetadata", []interface{}{metadata})
	fake.setUserMetadataMutex.Unlock()
	if fake.SetUserMetadataStub != nil {
		return fake.SetUserMetadataStub(metadata)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setUserMetadataReturns.result1, fake.setUserMetadataReturns.result2
}

func (fake *VirtualGuestService) SetUserMetadataCallCount() int {
	fake.setUserMetadataMutex.RLock()
	defer fake.setUserMetadataMutex.RUnlock()
	return len(fake.setUserMetadataArgsForCall)
}

func (fake *VirtualGuestService) SetUserMetadataArgsForCall(i int) string {
	fake.setUserMetadataMutex.RLock()
	defer fake.setUserMetadataMutex.RUnlock()
	return fake.setUserMetadataArgsForCall[i].metadata
}

func (fake *VirtualGuestService) SetUserMetadataReturns(result1 bool, result2 error) {
	fake.SetUserMetadataStub = nil
	fake.setUserMetadataReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) SetUserMetadataReturnsOnCall(i int, result1 bool, result2 error) {
	fake.SetUserMetadataStub = nil
	if fake.setUserMetadataReturnsOnCall == nil {
		fake.setUserMetadataReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.setUserMetadataReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *VirtualGuestService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.createObjectMutex.RLock()
	defer fake.createObjectMutex.RUnlock()
	fake.getObjectMutex.RLock()
	defer fake.getObjectMutex.RUnlock()
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	fake.getActiveTransactionsMutex.RLock()
	defer fake.getActiveTransactionsMutex.RUnlock()
	fake.rebootHardMutex.RLock()
	defer fake.rebootHardMutex.RUnlock()
	fake.rebootSoftMutex.RLock()
	defer fake.rebootSoftMutex.RUnlock()
	fake.powerOnMutex.RLock()
	defer fake.powerOnMutex.RUnlock()
	fake.reloadOperatingSystemMutex.RLock()
	defer fake.reloadOperatingSystemMutex.RUnlock()
	fake.setTagsMutex.RLock()
	defer fake.setTagsMutex.RUnlock()
	fake.setUserMetadataMutex.RLock()
	defer fake.setUserMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VirtualGuestService) recordInvocation(key string, args []interface{}) {
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

var _ backend.VirtualGuestService = new(VirtualGuestService)
