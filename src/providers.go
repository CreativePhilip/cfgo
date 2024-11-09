package cfgo

import (
	"os"
	"strings"
)

type ConfigSourceProvider interface {
	GetValues() map[string]string
}

type EnvVariablesSourceProvider struct {
}

func (p *EnvVariablesSourceProvider) GetValues() map[string]string {
	outValues := map[string]string{}

	for _, entry := range os.Environ() {
		envPair := strings.SplitN(entry, "=", 2)

		outValues[envPair[0]] = envPair[1]
	}

	return outValues
}

func NewEnvVariablesSourceProvider() ConfigSourceProvider {
	return &EnvVariablesSourceProvider{}
}

// ----------------------

type MockVariablesSourceProvider struct {
	values map[string]string
}

func (p *MockVariablesSourceProvider) GetValues() map[string]string {
	return p.values
}

func NewMockVariablesSourceProvider(values map[string]string) ConfigSourceProvider {
	return &MockVariablesSourceProvider{
		values: values,
	}
}
