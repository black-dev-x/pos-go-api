package database

import (
	"testing"

	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error("failed to connect to database:", err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("john@doe.com", "12345")
	userDB := NewUserDB(db)

	error := userDB.Create(user)
	assert.Nil(t, error)

	userFound, err := userDB.FindByEmail("john@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
}
