package main

import (
	"net/http"
	"html/template"
	"github.com/OlenaMaximiv/blog/handlers"
	_ "github.com/lib/pq"
	
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/comment", handlers.CommentHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/reset-password", handlers.ResetPasswordHandler)

	http.ListenAndServe(":8080", nil)

	err := InitializeDB()
	if err != nil {
		panic(err)
	}
	defer CloseDB()

}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the home page template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Execute the template, passing nil as the data since the home page doesn't require any dynamic data
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // Parse the login page template
        tmpl := template.Must(template.ParseFiles("templates/login.html"))

        // Execute the template
        err := tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else if r.Method == "POST" {
        // Handle the login form submission
        // Retrieve the username and password from the request form values
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Perform the login logic here
        if username == "admin" && password == "password" {
            // Successful login
            // Redirect the user to another page after successful login
            http.Redirect(w, r, "/profile", http.StatusFound)
        } else {
            // Failed login
            // Render the login page with an error message
            tmpl := template.Must(template.ParseFiles("templates/login.html"))

            // Create a data structure to pass error message to the template
            data := struct {
                Username string
                ErrorMessage string
            }{
                Username: username,
                ErrorMessage: "Invalid username or password",
            }

            // Execute the template with the data
            err := tmpl.Execute(w, data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }
    }
}
