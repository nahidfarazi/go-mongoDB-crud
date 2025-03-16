package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nahidfarazi/go-mongo2/database"
	"github.com/nahidfarazi/go-mongo2/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var collection = database.GetCollection()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	defer cancel()
	if err != nil {
		http.Error(w, "error finding user", http.StatusInternalServerError)
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, "error decoding users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	defer cancel()
	if err != nil {
		http.Error(w, "error finding user", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if user.ID == primitive.NilObjectID {
		user.ID = primitive.NewObjectID()
	}
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user %v:", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, "error decoding users", http.StatusInternalServerError)
		return
	}
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	result, err := collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"_name": user.Name, "_email": user.Email}})
	defer cancel()
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
