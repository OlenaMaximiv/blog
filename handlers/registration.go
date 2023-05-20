package handlers

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the registration form
		http.ServeFile(w, r, "templates/registration.html")
	} else if r.Method == http.MethodPost {
		// Parse form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Establish the database connection
		db, err := sql.Open("postgres", "user=postgres password=postgres dbname=blog sslmode=disable")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insert new user into the database
		stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to a success page or another endpoint
		http.Redirect(w, r, "/success", http.StatusSeeOther)
	}
}
