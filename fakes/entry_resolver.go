package fakes

import (
	"sync"

	"github.com/paketo-buildpacks/packit"
)

type EntryResolver struct {
	MergeLayerTypesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Name    string
			Entries []packit.BuildpackPlanEntry
		}
		Returns struct {
			LayerTypeSlice []packit.LayerType
		}
		Stub func(string, []packit.BuildpackPlanEntry) []packit.LayerType
	}
	ResolveEntriesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Name       string
			Entries    []packit.BuildpackPlanEntry
			Priorities map[string]int
		}
		Returns struct {
			BuildpackPlanEntry packit.BuildpackPlanEntry
			Bool               bool
		}
		Stub func(string, []packit.BuildpackPlanEntry, map[string]int) (packit.BuildpackPlanEntry, bool)
	}
}

func (f *EntryResolver) MergeLayerTypes(param1 string, param2 []packit.BuildpackPlanEntry) []packit.LayerType {
	f.MergeLayerTypesCall.Lock()
	defer f.MergeLayerTypesCall.Unlock()
	f.MergeLayerTypesCall.CallCount++
	f.MergeLayerTypesCall.Receives.Name = param1
	f.MergeLayerTypesCall.Receives.Entries = param2
	if f.MergeLayerTypesCall.Stub != nil {
		return f.MergeLayerTypesCall.Stub(param1, param2)
	}
	return f.MergeLayerTypesCall.Returns.LayerTypeSlice
}
func (f *EntryResolver) ResolveEntries(param1 string, param2 []packit.BuildpackPlanEntry, param3 map[string]int) (packit.BuildpackPlanEntry, bool) {
	f.ResolveEntriesCall.Lock()
	defer f.ResolveEntriesCall.Unlock()
	f.ResolveEntriesCall.CallCount++
	f.ResolveEntriesCall.Receives.Name = param1
	f.ResolveEntriesCall.Receives.Entries = param2
	f.ResolveEntriesCall.Receives.Priorities = param3
	if f.ResolveEntriesCall.Stub != nil {
		return f.ResolveEntriesCall.Stub(param1, param2, param3)
	}
	return f.ResolveEntriesCall.Returns.BuildpackPlanEntry, f.ResolveEntriesCall.Returns.Bool
}
