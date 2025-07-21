package category_cmd

import (
	category_handler "category-crud/module/shop/category/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/api/v1")
	{
		category_handler.NewCategoryHandler(routerGroup, db, log)
	}
}
