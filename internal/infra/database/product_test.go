package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/Gustavopnhro/api-golang/configs"
	"github.com/Gustavopnhro/api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	cfg, _ := configs.LoadConfig("../../..")
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + "127.0.0.1:" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.CreateProduct("Keyboard Red Dragon", 10.22, 10, "Eletronics")
	productDB := NewProductInstance(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

}

func TestFindAll(t *testing.T) {
	cfg, _ := configs.LoadConfig("../../..")
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + "127.0.0.1:" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.CreateProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100, 10, "Eletronics")
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProductInstance(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
}
