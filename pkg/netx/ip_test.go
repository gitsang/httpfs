package netx

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetIPv4s(t *testing.T) {
	ips, err := GetIPv4s()
	assert.NilError(t, err)
	_ = ips
	// t.Log(ips)
}
