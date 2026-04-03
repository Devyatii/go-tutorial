package verify

type VerifyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
