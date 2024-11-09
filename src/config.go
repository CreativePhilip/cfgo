package cfgo

type EnvConfiguration struct {
	BoolValidTrueValues []string

	Providers []ConfigSourceProvider
}

var (
	DefaultBoolValidTrueValues = []string{"t", "true", "y", "yes"}
)

func (c *EnvConfiguration) getAllValues() map[string]string {
	outMap := map[string]string{}

	for _, provider := range c.Providers {
		outMap = joinMaps(outMap, provider.GetValues())
	}

	return outMap
}

func NewEnvConfiguration(cfg EnvConfiguration) *EnvConfiguration {
	return &EnvConfiguration{
		BoolValidTrueValues: sliceOrDefault(cfg.BoolValidTrueValues, DefaultBoolValidTrueValues),
		Providers:           cfg.Providers,
	}
}
