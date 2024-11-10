package cfgo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvVariablesSourceProvider__HappyPath(t *testing.T) {
	os.Setenv("key1", "val1")

	provider := NewEnvVariablesSourceProvider()
	data := provider.GetValues()

	assert.Equal(t, "val1", data["key1"])
}

func TestEnvFileVariableSourceProvider__HappyPath(t *testing.T) {
	provider := NewEnvFileVariableSourceProvider("../fixtures/env/*.env")
	data := provider.GetValues()

	assert.Equal(t, "val1", data["key1"])
}
