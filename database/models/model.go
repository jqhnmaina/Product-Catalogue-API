package models

import (
	"../database"
	"../migrations"
	"fmt"
)

func CreateCategory(category migrations.Category) (migrations.Category, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&category).Error

	return category, err
}

func GetCategories() ([]migrations.CategoryWithCount, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var categories []migrations.Category

	err := db.Order("name asc").Find(&categories).Error

	if err != nil {
		return nil, err
	}

	var categoriesWithCount []migrations.CategoryWithCount
	for _, category := range categories {
		var categoryWithCount migrations.CategoryWithCount
		var count int64
		db.Model(&migrations.Product{}).Where("category_id = ?", category.ID).Count(&count)

		categoryWithCount.Name = category.Name
		categoryWithCount.ID = category.ID
		categoryWithCount.CreatedAt = category.CreatedAt
		categoryWithCount.DeletedAt = category.DeletedAt
		categoryWithCount.UpdatedAt = category.UpdatedAt
		categoryWithCount.ProductCount = count
		categoriesWithCount = append(categoriesWithCount, categoryWithCount)
	}

	return categoriesWithCount, err
}

func UpdateCategory(category migrations.Category, id int) (migrations.Category, error) {
	db := database.Connect()
	defer db.Close()

	var savedCate migrations.Category

	err := db.First(&savedCate, id).Error

	if err != nil {
		return category, err
	}

	err = db.Model(&savedCate).Update(migrations.Category{Name: category.Name}).Error

	return savedCate, err
}

func DeleteCategory(id int) error {
	db := database.Connect()
	defer db.Close()

	var category migrations.Category
	err := db.First(&category, id).Error

	if err != nil {
		return err
	}

	db.Delete(&category)
	return err
}

func GetCategoryProducts(categoryId int, sort string, sortDir string, limit int, offset int) ([]migrations.Product, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var products []migrations.Product

	err := db.Where("category_id = ?", categoryId).Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error

	return products, err
}

func CreateProduct(product migrations.Product) (migrations.Product, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&product).Error

	return product, err
}

func GetProducts(sort string, sortDir string, limit int, offset int) ([]migrations.Product, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var products []migrations.Product

	err := db.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error

	return products, err
}

func UpdateProduct(product migrations.Product, id int) (migrations.Product, error) {
	db := database.Connect()
	defer db.Close()

	var savedProd migrations.Product

	err := db.First(&savedProd, id).Error

	if err != nil {
		return product, err
	}

	err = db.Model(&savedProd).Update(migrations.Product{Name: product.Name, CategoryID: product.CategoryID}).Error

	return savedProd, err
}

func DeleteProduct(id int) error {
	db := database.Connect()
	defer db.Close()

	var product migrations.Product
	err := db.First(&product, id).Error

	if err != nil {
		return err
	}

	db.Delete(&product)
	return err
}
