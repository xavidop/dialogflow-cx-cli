package profileconversation

type Agent struct {
	Intent   string      `yaml:"intent"`
	Validate []*Validate `yaml:"validate"`
}
