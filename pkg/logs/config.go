package logs

type Config struct {
	Output           string `yaml:"output"`
	ValidateIDRegexp string `yaml:"validate-id-regexp"`
	SaveDebug        bool   `yaml:"save-debug"`
	SaveJSON         bool   `yaml:"save-json"`
}
