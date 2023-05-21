package handlers

import (
	"database/sql"

	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// RegisterHandler handles the registration functionality
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the registration form
		tmpl, err := template.ParseFiles("templates/registration.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		// Parse form values
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		// Validate form inputs
		formErrors := make(map[string]string)
		if username == "" {
			formErrors["username"] = "The username is required and cannot be empty"
		}
		if email == "" {
			formErrors["email"] = "The email is required and cannot be empty"
		}
		if password == "" {
			formErrors["password"] = "The password is required and cannot be empty"
		}
		if confirmPassword == "" {
			formErrors["confirm_password"] = "The confirm password is required and cannot be empty"
		}
		if password != confirmPassword {
			formErrors["confirm_password"] = "The password and its confirm are not the same"
		}

		if len(formErrors) > 0 {
			// Display the registration form with form errors
			tmpl, err := template.ParseFiles("templates/registration.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, formErrors)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		// Establish the database connection
		db, err := sql.Open("postgres", "user=postgres password=postgres dbname=blog sslmode=disable")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insert new user into the database
		_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the home page with a success message
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	// Start the HTTP server
	log.Println("Server started on :8080")
	http.HandleFunc("/registration", RegisterHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
