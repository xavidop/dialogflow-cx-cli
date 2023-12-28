package profilenlu

type Check struct {
	ID       string    `yaml:"id"`
	Input    *Input    `yaml:"input"`
	Validate *Validate `yaml:"validate"`
}
