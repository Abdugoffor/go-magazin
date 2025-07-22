package main

import (
	"category-crud/config"
	"category-crud/helper"
	"category-crud/middleware"
	auth_cmd "category-crud/module/auth"
	auth_model "category-crud/module/auth/model"
	shop_cmd "category-crud/module/shop"
	category_model "category-crud/module/shop/category/model"
	order_model "category-crud/module/shop/order/model"
	product_model "category-crud/module/shop/product/model"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	env := config.LoadEnv()

	db := config.InitDB(env)

	middleware.SetDB(db)

	if err := migration(db); err != nil {
		log.Fatal("❌ Migration error:", err)
	}

	router := echo.New()

	router.Validator = helper.NewValidator()

	shop_cmd.Cmd(router, db, log.Default())
	auth_cmd.Cmd(router, db, log.Default())

	if err := auth_cmd.SeedPermissions(db, router); err != nil {
		log.Fatal("❌ Permission seed error:", err)
	}

	log.Println("🚀 Server is running on port", env.HTTPPort)

	router.Start(":" + strconv.Itoa(env.HTTPPort))
}

func migration(db *gorm.DB) error {
	// log.Println("⚠️ Dropping all tables...")

	// tables, err := db.Migrator().GetTables()
	// if err != nil {
	// 	return err
	// }

	// for _, table := range tables {
	// 	log.Printf("🗑 Dropping table: %s\n", table)
	// 	if err := db.Migrator().DropTable(table); err != nil {
	// 		return err
	// 	}
	// }

	// log.Println("🔃 Running fresh migrations...")

	return db.AutoMigrate(
		&category_model.Category{},
		&product_model.Product{},
		&auth_model.User{},
		&auth_model.Role{},
		&auth_model.Permission{},
		&auth_model.PermissionGroup{},
		&auth_model.RoleUser{},
		&auth_model.RolePermission{},
		&order_model.Order{},
	)
}
