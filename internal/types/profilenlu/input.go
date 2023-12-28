package profilenlu

type Input struct {
	Type   string `yaml:"type"`
	Text   string `yaml:"text,omitempty"`
	Prompt string `yaml:"prompt,omitempty"`
	Audio  string `yaml:"audio,omitempty"`
}
