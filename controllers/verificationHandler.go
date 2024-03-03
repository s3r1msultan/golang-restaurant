package controllers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"final/db"
	"final/initializers"
	"final/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendVerificationEmail(firstName, lastName, to, token string, r *http.Request) error {
	//baseURL := fmt.Sprintf("%s://%s", r.URL.Scheme, r.Header.Get("Host"))
	//verificationURL := fmt.Sprintf("%s/verify?token=%s", baseURL, token)
	url := os.Getenv("VERIFICATION_URL")
	verificationURL := url + "/verify?token=" + token
	subject := "Verify Your Email Address"
	fmt.Println(verificationURL)
	emailTemplate := `
    <h2>Dear %s %s,</h2>
    <p>Thank you for registering with us. Please click on the link below to verify your email address and activate your account:</p>
    <p><a href="%s" target="_blank">Verify Email</a></p>
    <p>If you did not request this, please ignore this email.</p>
`
	message := fmt.Sprintf(emailTemplate, firstName, lastName, verificationURL)
	return SendMessage(to, subject, message)
}

func VerificationHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		initializers.LogError("token checking", nil, nil)
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}
	var result models.User
	err := db.GetUsersCollection().
		FindOneAndUpdate(
			context.TODO(),
			bson.M{"verification_token": token},
			bson.M{"$set": bson.M{"email_verified": true, "verification_token": ""}},
		).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Invalid token", http.StatusNotFound)
		} else {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, "/auth", http.StatusSeeOther)
}
