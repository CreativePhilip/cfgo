package cfgo

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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

// ----------------------

type EnvFileVariableSourceProvider struct {
	GlobPattern string
}

func (p *EnvFileVariableSourceProvider) GetValues() map[string]string {
	files, err := filepath.Glob(p.GlobPattern)
	outMap := map[string]string{}

	if err != nil {
		panic("cfgo: invalid pattern")
	}

	for _, file := range files {
		fileContent := must(os.ReadFile(file))
		lines := strings.Split(string(fileContent), "\n")

		for _, line := range lines {
			parsedLine := strings.SplitN(line, "=", 2)
			outMap[parsedLine[0]] = parsedLine[1]
		}
	}

	return outMap
}

func NewEnvFileVariableSourceProvider(globPattern string) ConfigSourceProvider {
	return &EnvFileVariableSourceProvider{GlobPattern: globPattern}
}

// ----------------------

type JsonFileVariableSourceProvider struct {
	GlobPattern string
}

func (p *JsonFileVariableSourceProvider) GetValues() map[string]string {
	files, err := filepath.Glob(p.GlobPattern)
	outMap := map[string]string{}

	if err != nil {
		panic("cfgo: invalid pattern")
	}

	for _, file := range files {
		fileContent := must(os.ReadFile(file))

		var data map[string]string
		if err := json.Unmarshal(fileContent, &data); err != nil {
			panic(err)
		}

		outMap = joinMaps(outMap, data)
	}

	return outMap
}

func NewJsonFileVariableSourceProvider(globPattern string) ConfigSourceProvider {
	return &JsonFileVariableSourceProvider{GlobPattern: globPattern}
}

// ----------------------

type YamlFileVariableSourceProvider struct {
	GlobPattern string
}

func (p *YamlFileVariableSourceProvider) GetValues() map[string]string {
	files, err := filepath.Glob(p.GlobPattern)
	outMap := map[string]string{}

	if err != nil {
		panic("cfgo: invalid pattern")
	}

	for _, file := range files {
		fileContent := must(os.ReadFile(file))
		var data map[string]string

		if err := yaml.Unmarshal(fileContent, &data); err != nil {
			panic(err)
		}

		outMap = joinMaps(outMap, data)
	}

	return outMap
}

func NewYamlFileVariableSourceProvider(globPattern string) ConfigSourceProvider {
	return &YamlFileVariableSourceProvider{GlobPattern: globPattern}
}
