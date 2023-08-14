package ebpfgo_mock

import (
	"testing"

	"github.com/khulnasoft-lab/ebpfgo"
)

// Just to ensure that MockProgram implements ebpfgo.Program interface
func TestMockProgram(t *testing.T) {
	mockBpf := NewMockSystem()
	mockBpf.Programs["test"] = NewMockProgram("test", ebpfgo.ProgramTypeXdp)
}
