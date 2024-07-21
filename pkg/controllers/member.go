package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"

	"Goland-Jam/pkg/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MemberController handles member CRUD operations
type MemberController struct {
	collection *mongo.Collection
}

// NewMemberController creates a new MemberController
func NewMemberController(client *mongo.Client) *MemberController {
	collection := client.Database("golandjam").Collection("members")
	return &MemberController{collection}
}

// CreateMember godoc
// @Summary Create a new member
// @Description Create a new member
// @Tags members
// @Accept  json
// @Produce  json
// @Param member body models.Member true "Member"
// @Success 200 {object} models.Member
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /members [post]
func (mc *MemberController) CreateMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := mc.collection.InsertOne(context.TODO(), member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// GetMember godoc
// @Summary Get a member by ID
// @Description Get a member by ID
// @Tags members
// @Produce  json
// @Param id query string true "Member ID"
// @Success 200 {object} models.Member
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Member Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /member [get]
func (mc *MemberController) GetMember(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	var member models.Member
	if err := mc.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&member); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Member not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(member)
}

// UpdateMember godoc
// @Summary Update a member by ID
// @Description Update a member by ID
// @Tags members
// @Accept  json
// @Produce  json
// @Param id query string true "Member ID"
// @Param member body models.Member true "Member"
// @Success 200 {object} models.Member
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /member [put]
func (mc *MemberController) UpdateMember(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := mc.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": member})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// DeleteMember godoc
// @Summary Delete a member by ID
// @Description Delete a member by ID
// @Tags members
// @Produce  json
// @Param id query string true "Member ID"
// @Success 200 {string} string "Member deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /member [delete]
func (mc *MemberController) DeleteMember(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	result, err := mc.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// ListMembers godoc
// @Summary List members with pagination
// @Description List members with pagination
// @Tags members
// @Produce  json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} models.Member
// @Failure 500 {string} string "Internal Server Error"
// @Router /members [get]
func (mc *MemberController) ListMembers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	options := options.Find()
	options.SetSkip(int64((page - 1) * size))
	options.SetLimit(int64(size))

	cursor, err := mc.collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var members []models.Member
	for cursor.Next(context.TODO()) {
		var member models.Member
		if err := cursor.Decode(&member); err != nil {
			log.Println("Error decoding member:", err)
			continue
		}
		members = append(members, member)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(members)
}
