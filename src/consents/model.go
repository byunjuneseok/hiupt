package consents

type Consent struct {
	PetitionId string `json:"petition_id"`
	UserId     string `json:"user_id"`
	Comment    string `json:"comment"`
}
