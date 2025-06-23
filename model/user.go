package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"bio"`
}

type UserPostBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"bio"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  User   `json:"data,omitempty"`
}
