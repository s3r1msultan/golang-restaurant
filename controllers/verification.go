package controllers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"final/db"
	"final/initializers"
	"final/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/smtp"
	"os"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendVerificationEmail(to, token string, r *http.Request) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var scheme string
	if r.TLS != nil {
		scheme = "https"
	} else {
		scheme = "http"
	}
	verificationURL := scheme + "://" + r.Host + "/verify?token=" + token
	message := []byte("To: " + to + "\r\n" +
		"Subject: Verify your email address\r\n\r\n" +
		"Click the link to verify your email address: " + verificationURL)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
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
