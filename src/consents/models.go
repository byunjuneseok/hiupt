package consents

type Consent struct {
	PetitionHashKey string `json:"petition_id"`
	UserHashKey     string `json:"user_id"`
	Comment         string `json:"comment"`
}
