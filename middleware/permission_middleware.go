package middleware

import (
	"category-crud/helper"
	auth_model "category-crud/module/auth/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func PermissionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if db == nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "database not initialized in middleware",
				})
			}

			userID, err := helper.GetUserFromToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "unauthorized",
				})
			}

			method := c.Request().Method

			path := c.Path() // Example: /api/v1/category

			fullRoute := fmt.Sprintf("%s:%s", method, path) // Example: GET:/api/v1/category

			var user auth_model.User

			err = db.Preload("Roles.Permissions").First(&user, userID).Error

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "failed to load user with roles",
				})
			}

			for _, role := range user.Roles {
				for _, perm := range role.Permissions {
					if strings.EqualFold(perm.Path, fullRoute) {
						return next(c)
					}
				}
			}

			return c.JSON(http.StatusForbidden, echo.Map{
				"error": "permission denied",
			})
		}
	}
}
