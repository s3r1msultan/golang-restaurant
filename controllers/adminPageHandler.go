package controllers

import (
	"context"
	"encoding/json"
	"final/db"
	"final/initializers"
	"final/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type notifyInfo struct {
	Email string `bson:"email" json:"email"`
}

type notifyMessage struct {
	Subject string `bson:"subject" json:"subject"`
	Message string `bson:"message" json:"message"`
}

func NotifyVerifiedUsers(w http.ResponseWriter, r *http.Request) {
	var message notifyMessage
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		initializers.LogError("parsing data from json request", err, nil)
		json.NewEncoder(w).Encode(bson.M{"is_sent": false})
		w.WriteHeader(http.StatusBadRequest)
	}

	filter := bson.M{
		"email_verified": true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cursor, err := db.GetUsersCollection().Find(ctx, filter)
	if err != nil {
		initializers.LogError("getting all users", err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(bson.M{"is_sent": false})
	}
	defer cursor.Close(ctx)
	var failedEmails []string
	var info notifyInfo
	for cursor.Next(ctx) {
		err := cursor.Decode(&info)
		if err != nil {
			initializers.LogError("trying to decode user from cursor", err, nil)
			continue
		}
		err = SendMessage(info.Email, message.Subject, message.Message)
		if err != nil {
			initializers.LogError("trying to send message to"+info.Email, err, nil)
			failedEmails = append(failedEmails, info.Email)
		}

	}
	response := bson.M{"is_sent": true}
	if len(failedEmails) > 0 {
		response["is_sent"] = false
		response["failed_emails"] = failedEmails
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = db.GetUsersCollection().FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.GetUsersCollection().InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	cursor, err := db.GetUsersCollection().Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var updateUser models.User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Failed to parse json data", http.StatusBadRequest)
		return
	}
	updateUser.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": userId}
	update := bson.M{"$set": updateUser}

	_, err = db.GetUsersCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bson.M{"message": "User updated successfully"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetUsersCollection().DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bson.M{"message": "User deleted successfully"})
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := initializers.InitTemplates()
	err := tmpl.ExecuteTemplate(w, "Admin.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
