package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bybhub/backend/internal/models"
	"github.com/bybhub/backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx            = context.TODO()
	userCollection *mongo.Collection
	testUser       *models.User
)

func setup() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@mongo:27017/"))
	if err != nil {
		panic(err)
	}
	db = client.Database("testdb")

	userCollection = db.Collection("user")

	testUser = &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	userId, err := repositories.CreateNewUser(db, "user", testUser)
	if err != nil {
		panic(err)
	}
	testUser.ID = userId

}

func teardown() {
	_, err := db.Collection("user").DeleteMany(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
}

func TestCreateUserHandler(t *testing.T) {
	setup()
	defer teardown()

	user := &models.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/user", CreateUserHandler)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	fmt.Println(response)

	assert.NotNil(t, response["data"].(map[string]any)["id"])
}

func TestGetUserHandler(t *testing.T) {
	setup()
	defer teardown()

	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	id, err := repositories.CreateNewUser(db, "user", user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	req, _ := http.NewRequest("GET", "/user/"+id.Hex(), nil)

	rr := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/user/:id", GetUserHandler)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, user.Name, response["data"].(map[string]any)["name"])
}

func TestGetAllUsersHandler(t *testing.T) {
	setup()
	defer teardown()

	_, _ = repositories.CreateNewUser(db, "user", &models.User{Name: "Alice", Email: "alice@example.com"})
	_, _ = repositories.CreateNewUser(db, "user", &models.User{Name: "Bob", Email: "bob@example.com"})

	req, _ := http.NewRequest("GET", "/users", nil)

	rr := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/users", GetAllUsers)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	users := response["data"].([]any)
	assert.Len(t, users, 3) // Incluindo o usuário de setup
}

func TestUpdateUserHandler(t *testing.T) {
	setup()
	defer teardown()

	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	id, err := repositories.CreateNewUser(db, "user", user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	updateData := bson.M{"name": "Updated Name"}
	jsonData, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("Failed to marshal update data: %v", err)
	}

	req, _ := http.NewRequest("PUT", "/user/"+id.Hex(), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r := gin.Default()
	r.PUT("/user/:id", UpdateUserHandler)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, "Updated Name", response["data"].(map[string]any)["name"])
}

func TestDeleteUserHandler(t *testing.T) {
	setup()
	defer teardown()

	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	id, err := repositories.CreateNewUser(db, "user", user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	req, _ := http.NewRequest("DELETE", "/user/"+id.Hex(), nil)

	rr := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/user/:id", DeleteUserHandler)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
