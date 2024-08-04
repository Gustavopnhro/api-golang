package database

import (
	"testing"

	"github.com/Gustavopnhro/api-golang/configs"
	"github.com/Gustavopnhro/api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	cfg, _ := configs.LoadConfig("../../..")
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + "127.0.0.1:" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.CreateUser("Jane Doe", "Janedoe@jane.com", "123456")
	userDB := NewUserInstance(db)

	err = userDB.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}
