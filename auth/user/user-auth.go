package user

//Auth is the user's authentication model
type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
