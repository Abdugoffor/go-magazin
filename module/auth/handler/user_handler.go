package auth_handler

import (
	auth_dto "category-crud/module/auth/dto"
	auth_service "category-crud/module/auth/service"
	"category-crud/module/middleware"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	db          *gorm.DB
	log         *log.Logger
	userService auth_service.UserService
}

func NewUserHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &userHandler{
		db:          db,
		log:         log,
		userService: auth_service.NewUserService(db),
	}

	routeGroup := group.Group("/user", middleware.PermissionMiddleware())
	{
		routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.POST("", handler.Create)
		routeGroup.PUT("/:id", handler.Update)
		routeGroup.DELETE("/:id", handler.Delete)
	}
}

func (u *userHandler) All(c echo.Context) error {
	data, err := u.userService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (u *userHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := u.userService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (u *userHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req auth_dto.CreateUser

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := u.userService.Create(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)

}

func (u *userHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req auth_dto.UpdateUser

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := u.userService.Update(ctx, uint(id), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (u *userHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	if err := u.userService.Delete(ctx, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
