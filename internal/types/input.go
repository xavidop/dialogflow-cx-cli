package types

type Input struct {
	Type  string `yaml:"type"`
	Text  string `yaml:"text,omitempty"`
	Audio string `yaml:"audio,omitempty"`
}
