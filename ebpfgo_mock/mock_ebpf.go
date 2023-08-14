package ebpfgo_mock

import (
	"io"

	"github.com/khulnasoft-lab/ebpfgo"
)

// MockSystem is mock implementation of eBPF system
type MockSystem struct {
	Programs map[string]ebpfgo.Program
	Maps     map[string]ebpfgo.Map

	// ErrorLoadElf specifies return value for LoadElf() method.
	ErrorLoadElf error
}

// NewMockSystem creates mocked eBPF system with:
// - empty program map
// - all linked mock maps
func NewMockSystem() *MockSystem {
	return &MockSystem{
		Programs: make(map[string]ebpfgo.Program),
		Maps:     MockMaps,
	}
}

// LoadElf does nothing, just a mock for original LoadElf
func (m *MockSystem) LoadElf(path string) error {
	return m.ErrorLoadElf
}

// Load does nothing, just a mock for original Load
func (m *MockSystem) Load(r io.ReaderAt) error {
	return m.ErrorLoadElf
}

// GetMaps returns all linked eBPF maps
func (m *MockSystem) GetMaps() map[string]ebpfgo.Map {
	return m.Maps
}

// GetPrograms returns map of added eBPF programs
func (m *MockSystem) GetPrograms() map[string]ebpfgo.Program {
	return m.Programs
}

// GetMapByName returns eBPF map by name or nil if not found
func (m *MockSystem) GetMapByName(name string) ebpfgo.Map {
	if result, ok := m.Maps[name]; ok {
		return result
	}
	return nil
}

// GetProgramByName returns eBPF program by name or nil if not found
func (m *MockSystem) GetProgramByName(name string) ebpfgo.Program {
	if result, ok := m.Programs[name]; ok {
		return result
	}
	return nil
}
