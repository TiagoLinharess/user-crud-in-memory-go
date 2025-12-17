package user

type User struct {
	FirstName string
	LastName  string
	Biography string
}

type UserBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string
}

type UserResponse struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string
}
