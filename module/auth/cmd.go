package auth

import (
	auth_handler "category-crud/module/auth/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {

	routerGroup := route.Group("/api/v1")
	{
		auth_handler.NewRoleHandler(routerGroup, db, log)
		auth_handler.NewPermissionGroupHandler(routerGroup, db, log)
		auth_handler.NewPermissionHandler(routerGroup, db, log)
		auth_handler.NewUserHandler(routerGroup, db, log)
		auth_handler.NewAuthHandler(routerGroup, db, log)
	}
}
