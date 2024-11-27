package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Main Function
func main() {
	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

	//Create router
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	// Wrap router with CORS and JSON content type middleware

	enchancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	log.Fatal(http.ListenAndServe(":8080", enchancedRouter))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Get all users
func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{} // Initialize empty array of users
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users) // Send the list of users as response
	}
}


// Get a single user
func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var u User
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		json.NewEncoder(w).Encode("User retrieved successfully")
	}
}

// Create a user
func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		err := db.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id", u.Username, u.Email).Scan(&u.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode("User created successfully")
	}
}

// Update a user
func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]
		// Update the user
		_, err := db.Exec("UPDATE users SET username = $1, email = $2 WHERE id = $3", u.Username, u.Email, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Retrieve the updated user from the database
		err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode("User updated successfully")
	}
}

// Delete a user
func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var u User
		err := db.QueryRow("DELETE FROM users WHERE id = $1", id).Scan(&u.ID)
		if err != nil {
			// Todo: Handle error
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("User deleted successfully")
	}
}