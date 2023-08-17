package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"
	"log"
	"time"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := database.StartDB()
	if err != nil {
		log.Fatalln("Error starting DB:", err)
		return
	}

	// Create new product
	// createProduct("Skincare")

	// Get product by ID
	// getProductById(1)

	// Update product by ID
	// updateProduct(1, "Local Skincare")

	// Create variant (triggering hooks)
	// createVariant(1, "Somethinc", -1)

	// Create variant
	// createVariant(1, "Somethinc", 5)
	// createVariant(1, "Harlette", 5)

	// Get product with variant
	// getProductWithVariant()

	// Delete variant by Id
	// deleteVariantById(4)

}

func createProduct(name string) {
	db := database.GetDB()
	product := models.Product{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := db.Create(&product).Error
	if err != nil {
		log.Fatalln("Error creating new product:", err)
		return
	}

	fmt.Println("New product data:", product)
}

func updateProduct(id uint, name string) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Model(&product).Where("id = ?", id).Updates(models.Product{
		Name: name,
	}).Error

	if err != nil {
		log.Fatalln("Error updating product data:", err)
		return
	}

	fmt.Println("Updated product data:", product)
}

func getProductById(id uint) {
	product := models.Product{}

	db := database.GetDB()

	// Retrieve object with primary key
	err := db.First(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			return
		}
		log.Fatalln("Error finding product")
	}

	fmt.Println("Product data:", product)
}

func createVariant(productID uint, name string, quantity int) {
	db := database.GetDB()

	Variants := models.Variants{
		ProductID:   productID,
		VariantName: name,
		Quantity:    quantity,
	}

	err := db.Create(&Variants).Error
	if err != nil {
		log.Fatalln("Error creating variant:", err)
		return
	}

	fmt.Println("New variant data:", Variants)
}

func getProductWithVariant() {
	db := database.GetDB()
	product := models.Product{}
	err := db.Preload("Variants").Find(&product).Error
	if err != nil {
		log.Fatalln("Error getting product data with variants:", err.Error())
		return
	}

	fmt.Println("Products data with variants:")
	fmt.Printf("%+v\n", product)
}

func deleteVariantById(id uint) {
	db := database.GetDB()
	variants := models.Variants{}

	err := db.Where("id = ?", id).Delete(variants).Error
	if err != nil {
		log.Fatalln("Error deleting variant:", err)
		return
	}

	fmt.Println("Deleted variant with id of", id)
}
