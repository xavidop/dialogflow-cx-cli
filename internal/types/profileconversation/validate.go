package profileconversation

import (
	"github.com/adrg/strutil/metrics"
	"github.com/xavidop/dialogflow-cx-cli/internal/types/profileconversation/configurations"
)

type Validate struct {
	Type                            string                      `yaml:"type"`
	Value                           string                      `yaml:"value"`
	Algorithm                       string                      `yaml:"algorithm,omitempty"`
	Threshold                       float64                     `yaml:"threshold,omitempty"`
	ConfigurationContains           *configurations.Contains    `yaml:"configuration-contains,omitempty"`
	ConfigurationRegexp             *configurations.Regexp      `yaml:"configuration-regexp,omitempty"`
	ConfigurationEquals             *configurations.Equals      `yaml:"configuration-equals,omitempty"`
	ConfigurationHamming            *metrics.Hamming            `yaml:"configuration-hamming,omitempty"`
	ConfigurationLevenshtein        *metrics.Levenshtein        `yaml:"configuration-levenshtein,omitempty"`
	ConfigurationJaro               *metrics.Jaro               `yaml:"configuration-jaro,omitempty"`
	ConfigurationJaroWinkler        *metrics.JaroWinkler        `yaml:"configuration-jaro-winkler,omitempty"`
	ConfigurationSmithWatermanGotoh *metrics.SmithWatermanGotoh `yaml:"configuration-smith-waterman-gotoh,omitempty"`
	ConfigurationSorensenDice       *metrics.SorensenDice       `yaml:"configuration-sorensen-dice,omitempty"`
	ConfigurationJaccard            *metrics.Jaccard            `yaml:"configuration-jaccard,omitempty"`
	ConfigurationOverlapCoefficient *metrics.OverlapCoefficient `yaml:"configuration-overlap-coefficient,omitempty"`
}
