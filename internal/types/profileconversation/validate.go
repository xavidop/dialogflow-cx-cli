package profileconversation

import (
	"github.com/adrg/strutil/metrics"
	"github.com/xavidop/dialogflow-cx-cli/internal/types/profileconversation/configurations"
)

// type Validate struct {
// 	ValidateConfigRegexp
// 	ValidateConfigContains
// 	ValidateConfigEquals
// 	ValidateConfigHamming
// 	ValidateConfigJaccard
// 	ValidateConfigJaro
// 	ValidateConfigJaroWinkler
// 	ValidateConfigLevenshtein
// 	ValidateConfigOverlapCoefficient
// 	ValidateConfigSmithWatermanGotoh
// 	ValidateConfigSorensenDice
// }

// type Validate struct {
// 	ValidateConfigBase
// 	ValidateConfigSimilarityBase
// 	ConfigurationContains           *configurations.Contains    `yaml:"configuration-contains,omitempty"`
// 	ConfigurationRegexp             *configurations.Regexp      `yaml:"configuration-regexp,omitempty"`
// 	ConfigurationEquals             *configurations.Equals      `yaml:"configuration-equals,omitempty"`
// 	ConfigurationHamming            *metrics.Hamming            `yaml:"configuration-hamming,omitempty"`
// 	ConfigurationLevenshtein        *metrics.Levenshtein        `yaml:"configuration-levenshtein,omitempty"`
// 	ConfigurationJaro               *metrics.Jaro               `yaml:"configuration-jaro,omitempty"`
// 	ConfigurationJaroWinkler        *metrics.JaroWinkler        `yaml:"configuration-jaro-winkler,omitempty"`
// 	ConfigurationSmithWatermanGotoh *metrics.SmithWatermanGotoh `yaml:"configuration-smith-waterman-gotoh,omitempty"`
// 	ConfigurationSorensenDice       *metrics.SorensenDice       `yaml:"configuration-sorensen-dice,omitempty"`
// 	ConfigurationJaccard            *metrics.Jaccard            `yaml:"configuration-jaccard,omitempty"`
// 	ConfigurationOverlapCoefficient *metrics.OverlapCoefficient `yaml:"configuration-overlap-coefficient,omitempty"`
// }

type ValidateConfigBase struct {
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type ValidateConfigSimilarityBase struct {
	Algorithm string  `yaml:"algorithm,omitempty"`
	Threshold float64 `yaml:"threshold,omitempty"`
}

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

type ValidateConfigContains struct {
	ValidateConfigBase
	Configuration *configurations.Contains `yaml:"configuration"`
}

type ValidateConfigRegexp struct {
	ValidateConfigBase
	Configuration *configurations.Regexp `yaml:"configuration"`
}

type ValidateConfigEquals struct {
	ValidateConfigBase
	Configuration *configurations.Equals `yaml:"configuration"`
}

type ValidateConfigHamming struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.Hamming `yaml:"configuration"`
}

type ValidateConfigLevenshtein struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.Levenshtein `yaml:"configuration"`
}

type ValidateConfigJaro struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.Jaro `yaml:"configuration"`
}

type ValidateConfigJaroWinkler struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.JaroWinkler `yaml:"configuration"`
}

type ValidateConfigSmithWatermanGotoh struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.SmithWatermanGotoh `yaml:"configuration"`
}

type ValidateConfigSorensenDice struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.SorensenDice `yaml:"configuration"`
}

type ValidateConfigJaccard struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.Jaccard `yaml:"configuration"`
}

type ValidateConfigOverlapCoefficient struct {
	ValidateConfigBase
	ValidateConfigSimilarityBase
	Configuration *metrics.OverlapCoefficient `yaml:"configuration"`
}
