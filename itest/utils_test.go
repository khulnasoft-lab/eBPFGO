package itest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft-labs/ebpfgo"
)

func TestGetNumOfPossibleCpus(t *testing.T) {
	cpus, err := ebpfgo.GetNumOfPossibleCpus()
	assert.NoError(t, err)
	assert.True(t, cpus > 0)
}
