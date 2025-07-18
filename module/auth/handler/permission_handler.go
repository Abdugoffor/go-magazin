package auth_handler

import (
	auth_dto "category-crud/module/auth/dto"
	auth_service "category-crud/module/auth/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type permissionHandler struct {
	db                *gorm.DB
	log               *log.Logger
	permissionService auth_service.PermissionService
}

func NewPermissionHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &permissionHandler{
		db:                db,
		log:               log,
		permissionService: auth_service.NewPermissionService(db),
	}

	routeGroup := group.Group("/permission")
	{
		routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.PUT("/:id", handler.Update)
	}
}

func (h *permissionHandler) All(c echo.Context) error {
	data, err := h.permissionService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *permissionHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.permissionService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *permissionHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req auth_dto.UpdatePermission

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.permissionService.Update(ctx, uint(id), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
