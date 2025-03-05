package repositories

import (
	"testing"

	"github.com/bybhub/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db             *mongo.Database
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
	userId, err := CreateNewUser(db, "user", testUser)
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

func TestCreateNewUser(t *testing.T) {
	setup()
	defer teardown()

	user := &models.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}

	id, err := CreateNewUser(db, "user", user)
	assert.NoError(t, err)
	assert.NotEqual(t, id, primitive.NilObjectID)
}

func TestFindAllUsers(t *testing.T) {
	setup()
	defer teardown()

	user1 := &models.User{Name: "Alice", Email: "alice@example.com"}
	user2 := &models.User{Name: "Bob", Email: "bob@example.com"}
	_, _ = CreateNewUser(db, "user", user1)
	_, _ = CreateNewUser(db, "user", user2)

	users, err := FindAllUsers(db, "user")
	assert.NoError(t, err)
	assert.Len(t, users, 3)

	assert.Equal(t, "Alice", users[1].Name)
}

func TestFindUserByID(t *testing.T) {
	setup()
	defer teardown()

	userResponse, err := FindUserByID(db, "user", testUser.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, testUser.Name, userResponse.Name)
	assert.Equal(t, testUser.Email, userResponse.Email)

	// Teste: ID inválido
	_, err = FindUserByID(db, "user", "invalid_id")
	assert.Error(t, err)
	assert.Equal(t, "ID inválido", err.Error())
}

func TestUpdateUserByID(t *testing.T) {
	setup()
	defer teardown()

	update := bson.M{"name": "Updated Name"}
	updatedUser, err := UpdateUserByID(db, "user", testUser.ID.Hex(), update)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.Name)

	updateNoChange := bson.M{"name": "Updated Name"} // O nome já está "Updated Name"
	_, err = UpdateUserByID(db, "user", testUser.ID.Hex(), updateNoChange)
	assert.Error(t, err)
	assert.Equal(t, "nenhuma alteração realizada, dados já estavam atualizados", err.Error())

	_, err = UpdateUserByID(db, "user", "invalid_id", update)
	assert.Error(t, err)
	assert.Equal(t, "ID inválido", err.Error())
}

func TestDeleteUserByID(t *testing.T) {
	setup()
	defer teardown()

	err := DeleteUserByID(db, "user", testUser.ID.Hex())
	assert.NoError(t, err)

	err = DeleteUserByID(db, "user", "invalid_id")
	assert.Error(t, err)
	assert.Equal(t, "ID inválido", err.Error())
}
