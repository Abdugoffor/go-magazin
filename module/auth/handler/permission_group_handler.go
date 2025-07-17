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

type permissionGroupHandler struct {
	db                     *gorm.DB
	log                    *log.Logger
	permissionGroupService auth_service.PermissionGroupService
}

func NewPermissionGroupHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &permissionGroupHandler{
		db:                     db,
		log:                    log,
		permissionGroupService: auth_service.NewPermissionGroupService(db),
	}

	routeGroup := group.Group("/permission-group")
	{
		routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.PUT("/:id", handler.Update)
	}
}

func (h *permissionGroupHandler) All(c echo.Context) error {

	data, err := h.permissionGroupService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *permissionGroupHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.permissionGroupService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *permissionGroupHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req auth_dto.UpdatePermissionGroup
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.permissionGroupService.Update(ctx, uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
