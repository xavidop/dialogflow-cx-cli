package profileconversation

type Interaction struct {
	ID    string `yaml:"id"`
	User  *User  `yaml:"user"`
	Agent *Agent `yaml:"agent"`
}
