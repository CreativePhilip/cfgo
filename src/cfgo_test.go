package cfgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadType__PanicOnMissingValue(t *testing.T) {
	defer func() { recover() }()

	type TestType struct {
		Key1 string `env:"key_1"`
		Key2 string `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{}),
	}})

	LoadType(v, cfg)

	t.Errorf("did not panic on missing value2")
}

func TestLoadType__HandlesString(t *testing.T) {
	type TestType struct {
		Key1 string `env:"key_1"`
		Key2 string `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{"key_1": "value_1", "key_2": "value_2"}),
	}})

	LoadType(v, cfg)

	assert.Equal(t, "value_1", v.Key1)
	assert.Equal(t, "value_2", v.Key2)
}

func TestLoadType__HandlesBool(t *testing.T) {
	type TestType struct {
		Key1 bool `env:"key_1"`
		Key2 bool `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{"key_1": "yes", "key_2": "value_2"}),
	}})

	LoadType(v, cfg)

	assert.True(t, v.Key1)
	assert.False(t, v.Key2)
}

func TestLoadType__HandlesInt(t *testing.T) {
	type TestType struct {
		Key1 int `env:"key_1"`
		Key2 int `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{"key_1": "1", "key_2": "-2"}),
	}})

	LoadType(v, cfg)

	assert.Equal(t, 1, v.Key1)
	assert.Equal(t, -2, v.Key2)
}

func TestLoadType__HandlesUint(t *testing.T) {
	type TestType struct {
		Key1 uint `env:"key_1"`
		Key2 uint `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{"key_1": "1", "key_2": "2"}),
	}})

	LoadType(v, cfg)

	assert.Equal(t, uint(1), v.Key1)
	assert.Equal(t, uint(2), v.Key2)
}

func TestLoadType__HandlesFloat(t *testing.T) {
	type TestType struct {
		Key1 float64 `env:"key_1"`
		Key2 float64 `env:"key_2"`
	}

	v := &TestType{}
	cfg := NewEnvConfiguration(EnvConfiguration{Providers: []ConfigSourceProvider{
		NewMockVariablesSourceProvider(map[string]string{"key_1": "1.23", "key_2": "2.34"}),
	}})

	LoadType(v, cfg)

	assert.InDelta(t, 1.23, v.Key1, 0.001)
	assert.InDelta(t, 2.34, v.Key2, 0.001)
}
