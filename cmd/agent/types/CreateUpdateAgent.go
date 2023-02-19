package types

type CreateUpdateAgent struct {
	Description              string
	TimeZone                 string
	DefaultLanguageCode      string
	SupportedLanguageCodes   []string
	AvatarURI                string
	EnableStackdriverLogging bool
	EnableInteractionLogging bool
	EnableSpellCorrection    bool
	EnableSpeechAdaptation   bool
}
