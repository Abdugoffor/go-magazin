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

type roleHandler struct {
	db          *gorm.DB
	log         *log.Logger
	roleService auth_service.RoleService
}

func NewRoleHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &roleHandler{
		db:          db,
		log:         log,
		roleService: auth_service.NewRoleService(db),
	}

	routeGroup := group.Group("/role")
	{
		routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.POST("", handler.Create)
		routeGroup.PUT("/:id", handler.Update)
		routeGroup.DELETE("/:id", handler.Delete)

	}
}

func (h *roleHandler) All(c echo.Context) error {
	data, err := h.roleService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *roleHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.roleService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *roleHandler) Create(c echo.Context) error {

	ctx := c.Request().Context()

	var req auth_dto.CreateRole

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := h.roleService.Create(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *roleHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req auth_dto.UpdateRole

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := h.roleService.Update(ctx, uint(id), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
func (h *roleHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	if err := h.roleService.Delete(ctx, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
