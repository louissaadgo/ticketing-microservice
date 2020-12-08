package user

//Model is the user's authentication model
type Model struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
