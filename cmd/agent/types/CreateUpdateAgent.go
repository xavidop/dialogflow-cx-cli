package types

type CreateUpdateAgent struct {
	Description              string
	TimeZone                 string
	DefaultLanguageCode      string
	SupportedLanguageCodes   []string
	AvatarURI                string
	EnableStackdriverLogging string
	EnableInteractionLogging string
	EnableSpellCorrection    string
	EnableSpeechAdaptation   string
}
