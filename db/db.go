package db

import (
	"api-go/model"
)

type AppDB struct {
	Users map[string]model.User
}

func NewAppDB() (*AppDB, error) {
	users := make(map[string]model.User)
	appDB := &AppDB{
		Users: users,
	}

	// Initialize with some dummy data
	appDB.Users["1"] = model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Biography: "A sample user",
	}

	return appDB, nil
}
