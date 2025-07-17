package auth

import (
	auth_model "category-crud/module/auth/model"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB, e *echo.Echo) error {
	routes := e.Routes()
	// 1. Avval hamma tegishli jadvalni tozalaymiz

	// db.Exec("TRUNCATE TABLE role_permissions RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE permissions RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE permission_groups RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE roles RESTART IDENTITY CASCADE")

	// 1. PermissionGroup larni aniqlash va yaratish
	groupMap := make(map[string]auth_model.PermissionGroup)
	for _, route := range routes {
		segments := strings.Split(route.Path, "/")
		if len(segments) > 3 {
			groupName := strings.ToLower(segments[3]) // masalan: /api/v1/product → product
			if _, ok := groupMap[groupName]; !ok {
				group := auth_model.PermissionGroup{
					Name:     groupName,
					IsActive: true,
				}
				if err := db.Where("name = ?", groupName).FirstOrCreate(&group).Error; err != nil {
					return err
				}
				groupMap[groupName] = group
			}
		}
	}

	// 2. Har bir route uchun permission yaratish
	for _, route := range routes {
		segments := strings.Split(route.Path, "/")
		if len(segments) <= 3 {
			continue // /api yoki /api/v1 kabi route bo‘lsa o‘tkazamiz
		}

		groupName := strings.ToLower(segments[3])
		group, ok := groupMap[groupName]
		if !ok {
			continue
		}

		perm := auth_model.Permission{
			Name:     generatePermissionName(route.Method, route.Path),
			Path:     fmt.Sprintf("%s:%s", route.Method, route.Path),
			GroupID:  &group.ID,
			IsActive: true,
		}
		if err := db.Where("path = ?", perm.Path).FirstOrCreate(&perm).Error; err != nil {
			return err
		}
	}

	// 3. Role lar
	roles := []string{"admin", "moderator", "user"}
	roleMap := make(map[string]auth_model.Role)
	for _, roleName := range roles {
		role := auth_model.Role{
			Name:     roleName,
			IsActive: true,
		}
		if err := db.Where("name = ?", roleName).FirstOrCreate(&role).Error; err != nil {
			return err
		}
		roleMap[roleName] = role
	}

	// 4. RolePermission biriktirish
	var allPerms []auth_model.Permission
	if err := db.Find(&allPerms).Error; err != nil {
		return err
	}

	for _, perm := range allPerms {
		path := perm.Path

		// admin → barcha permission
		db.FirstOrCreate(&auth_model.RolePermission{}, auth_model.RolePermission{
			RoleID:       roleMap["admin"].ID,
			PermissionID: perm.ID,
		})

		// moderator → category va product uchun CRUD
		if strings.Contains(path, "/category") || strings.Contains(path, "/product") {
			db.FirstOrCreate(&auth_model.RolePermission{}, auth_model.RolePermission{
				RoleID:       roleMap["moderator"].ID,
				PermissionID: perm.ID,
			})
		}

		// user → category faqat GET va POST
		if strings.Contains(path, "/category") &&
			(strings.HasPrefix(path, "GET") || strings.HasPrefix(path, "POST")) {
			db.FirstOrCreate(&auth_model.RolePermission{}, auth_model.RolePermission{
				RoleID:       roleMap["user"].ID,
				PermissionID: perm.ID,
			})
		}
	}

	return nil
}

func generatePermissionName(method, path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		entity := strings.Title(parts[3])
		if len(parts) > 4 && !strings.Contains(parts[4], ":") {
			action := strings.Title(parts[4])
			return fmt.Sprintf("%s %s %s", entity, action, method)
		}

		switch method {
		case "GET":
			if strings.Contains(path, ":id") {
				return entity + " View"
			}
			return entity + " List"
		case "POST":
			return entity + " Create"
		case "PUT", "PATCH":
			return entity + " Update"
		case "DELETE":
			return entity + " Delete"
		}
	}

	return method + " " + path
}
