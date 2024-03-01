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
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendVerificationEmail(to string, token string) error {
	verificationURL := "https://golang-restaurant.onrender.com" + "/verify?token=" + token
	subject := "Verify your email address"
	message := "To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		"Click the link to verify your email address: " + verificationURL
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
