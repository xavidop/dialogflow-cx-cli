package types

type Check struct {
	ID       string   `yaml:"id"`
	Input    string   `yaml:"input"`
	Validate Validate `yaml:"validate"`
}
