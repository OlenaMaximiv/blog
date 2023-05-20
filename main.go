package main

import (
	"net/http"

	"github.com/OlenaMaximiv/blog/handlers"
	
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/comment", handlers.CommentHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/registration", handlers.RegistrationHandler)
	http.HandleFunc("/reset-password", handlers.ResetPasswordHandler)

	http.ListenAndServe(":8080", nil)
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




