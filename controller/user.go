package controller

import (
	"api-go/db"
	"api-go/model"
	"api-go/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func HandleDeleteUser(db db.AppDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle DELETE request
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.SendJSON(w, utils.Response{Error: "user ID is required"}, http.StatusBadRequest)
			return
		}
		if _, exists := db.Users[id]; !exists {
			utils.SendJSON(w, utils.Response{Error: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return
		}

		err := deleteUser(db, id)
		if err != nil {
			utils.SendJSON(w, utils.Response{Error: "The user could not be removed"}, http.StatusInternalServerError)
			return
		}
		utils.SendJSON(w, utils.Response{Data: "user deleted successfully"}, http.StatusNoContent)

	}
}

func deleteUser(db db.AppDB, id string) error {
	delete(db.Users, id)
	return nil
}

func HandleUpdateUser(db db.AppDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle PUT request
		var body model.User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJSON(w, utils.Response{Error: "invalid request body"}, http.StatusUnprocessableEntity)
			return
		}
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.SendJSON(w, utils.Response{Error: "user ID is required"}, http.StatusBadRequest)
			return
		}
		if _, exists := db.Users[id]; !exists {
			utils.SendJSON(w, utils.Response{Error: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return
		}
		if body.ID != id {
			utils.SendJSON(w, utils.Response{Error: "user ID in body does not match URL"}, http.StatusBadRequest)
			return
		}

		// Validate that FirstName, LastName, and Biography are not empty
		if body.FirstName == "" || body.LastName == "" || body.Biography == "" {
			utils.SendJSON(w, utils.Response{Error: "Please provide name and bio for the user"}, http.StatusBadRequest)
			return
		}

		// Validate the length of FirstName, LastName, and Biography
		if len(body.FirstName) < 2 || len(body.FirstName) > 20 ||
			len(body.LastName) < 2 || len(body.LastName) > 20 ||
			len(body.Biography) < 20 || len(body.Biography) > 450 {
			utils.SendJSON(w, utils.Response{Error: "FirstName and LastName should be between 2 and 20 characters, Biography should be between 20 and 450 characters"}, http.StatusBadRequest)
			return
		}

		// Update user
		err := updateUser(db, id, body)
		if err != nil {
			utils.SendJSON(w, utils.Response{Error: "The user information could not be modified"}, http.StatusInternalServerError)
			return
		}
		utils.SendJSON(w, utils.Response{Data: body}, http.StatusOK)

	}
}

func updateUser(db db.AppDB, id string, body model.User) error {
	db.Users[id] = body
	return nil
}

func HandleCreateUser(db db.AppDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle user creation
		var body model.UserPostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJSON(w, utils.Response{Error: "invalid request body"}, http.StatusUnprocessableEntity)
			return
		}

		// Validate that FirstName, LastName, and Biography are not empty
		if body.FirstName == "" || body.LastName == "" || body.Biography == "" {
			utils.SendJSON(w, utils.Response{Error: "Please provide FirstName LastName and bio for the user"}, http.StatusBadRequest)
			return
		}

		// Validate the length of FirstName, LastName, and Biography
		if len(body.FirstName) < 2 || len(body.FirstName) > 20 ||
			len(body.LastName) < 2 || len(body.LastName) > 20 ||
			len(body.Biography) < 20 || len(body.Biography) > 450 {
			utils.SendJSON(w, utils.Response{Error: "FirstName and LastName should be between 2 and 20 characters, Biography should be between 20 and 450 characters"}, http.StatusBadRequest)
			return
		}

		// Create a new user
		user, err := insertUser(body, db)
		if err != nil {
			utils.SendJSON(w, utils.Response{Error: "There was an error while saving the user to the database"}, http.StatusInternalServerError)
			return
		}

		utils.SendJSON(w, utils.Response{Data: user}, http.StatusCreated)
	}
}

func insertUser(body model.UserPostBody, db db.AppDB) (*model.User, error) {
	user := &model.User{
		ID:        uuid.NewString(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Biography: body.Biography,
	}

	db.Users[user.ID] = *user
	return user, nil
}

func HandleGetAllUsers(db db.AppDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle GET request, return an array of users
		users, err := getUsersFromDB(db)
		if err != nil {
			utils.SendJSON(w, utils.Response{Error: "The users information could not be retrieved"}, http.StatusInternalServerError)
			return
		}
		if len(users) == 0 {
			// return an empty array if no users found
			utils.SendJSON(w, utils.Response{Data: []model.User{}}, http.StatusOK)
			return
		}
		// Convert a map to slice
		var userSlice []model.User
		for _, user := range users {
			userSlice = append(userSlice, user)

		}
		utils.SendJSON(w, utils.Response{Data: userSlice}, http.StatusOK)
	}
}

func getUsersFromDB(db db.AppDB) (map[string]model.User, error) {
	users := db.Users
	return users, nil
}

func HandleGetUserByID(db db.AppDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle GET request by ID, id is on path
		id := chi.URLParam(r, "id")
		user, exists := db.Users[id]
		if !exists {
			utils.SendJSON(w, utils.Response{Error: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return
		}

		// check if user is empty
		if user.ID == "" || user.FirstName == "" || user.LastName == "" || user.Biography == "" {
			utils.SendJSON(w, utils.Response{Error: "The user information could not be retrieved"}, http.StatusInternalServerError)
			return
		}

		utils.SendJSON(w, utils.Response{Data: user}, http.StatusOK)
	}
}
