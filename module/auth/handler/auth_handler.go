package auth_handler

import (
	auth_dto "category-crud/module/auth/dto"
	auth_service "category-crud/module/auth/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authHandler struct {
	db          *gorm.DB
	log         *log.Logger
	authService auth_service.UserService
}

func NewAuthHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &authHandler{
		db:          db,
		log:         log,
		authService: auth_service.NewUserService(db),
	}

	routeGroup := group.Group("/auth")
	{
		routeGroup.POST("/login", handler.Login)
		routeGroup.POST("/register", handler.Register)
	}
}

func (handler *authHandler) Login(c echo.Context) error {
	return nil
}

func (handler *authHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var req auth_dto.RegisterUser

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := handler.authService.Register(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)
}
