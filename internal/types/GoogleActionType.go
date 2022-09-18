package types

type GoogleActionType struct {
	Synonym *Synonym `yaml:"synonym"`
}

type Synonym struct {
	Entities  map[string]*Synonyms `yaml:"entities"`
	MatchType string               `default:"EXACT_MATCH" yaml:"matchType"`
}

type Synonyms struct {
	Synonyms []string `yaml:"synonyms"`
}
