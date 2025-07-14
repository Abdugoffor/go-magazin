package main

import (
	"category-crud/config"
	category_cmd "category-crud/module/category"
	category_model "category-crud/module/category/model"
	product_cmd "category-crud/module/product"
	product_model "category-crud/module/product/model"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	env := config.LoadEnv()

	db := config.InitDB(env)

	if err := migration(db); err != nil {
		log.Fatal("❌ Migration error:", err)
	}

	router := echo.New()

	category_cmd.Cmd(router, db, log.Default())
	product_cmd.Cmd(router, db, log.Default())

	log.Println("🚀 Server is running on port", env.HTTPPort)

	router.Start(":" + strconv.Itoa(env.HTTPPort))
}

func migration(db *gorm.DB) error {
	// log.Println("⚠️ Dropping all tables...")

	// // Barcha jadvallar nomini olish
	// tables, err := db.Migrator().GetTables()
	// if err != nil {
	// 	return err
	// }

	// // Har bir jadvalni o‘chirish
	// for _, table := range tables {
	// 	log.Printf("🗑 Dropping table: %s\n", table)
	// 	if err := db.Migrator().DropTable(table); err != nil {
	// 		return err
	// 	}
	// }

	// log.Println("🔃 Running fresh migrations...")

	// Endi kerakli modellaringizni qayta yaratish
	return db.AutoMigrate(
		&category_model.Category{},
		&product_model.Product{},
	)
}
