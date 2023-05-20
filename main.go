package main

import (
	"net/http"

	"github.com/your-username/mywebapp/handlers"
)

func main() {
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
