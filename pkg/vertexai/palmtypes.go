package vertexai

type Content struct {
	Content string `json:"content"`
}

type Parameters struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens"`
	TopP            float64 `json:"topP"`
	TopK            int     `json:"topK"`
}

type Completion struct {
	Text string `json:"text"`
}

type PalmRequest struct {
	Instances  []Content   `json:"instances"`
	Parameters *Parameters `json:"parameters"`
}

type PalmResponse struct {
	Predictions []struct {
		SafetyAttributes struct {
			Categories []string  `json:"categories"`
			Blocked    bool      `json:"blocked"`
			Scores     []float64 `json:"scores"`
		} `json:"safetyAttributes"`
		CitationMetadata struct {
			Citations []any `json:"citations"`
		} `json:"citationMetadata"`
		Content string `json:"content"`
	} `json:"predictions"`
}
