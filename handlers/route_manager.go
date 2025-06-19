package handlers

import (
	"net/http"
	"sync"
	"time"

	"management/config"
	"management/models"

	"github.com/gin-gonic/gin"
)

type RouteManager struct {
	router *gin.Engine
	mu     sync.RWMutex
}

var Manager *RouteManager

func NewRouteManager(router *gin.Engine) *RouteManager {
	Manager = &RouteManager{
		router: router,
	}
	return Manager
}

// LoadRoutes 从数据库加载所有路由
func (rm *RouteManager) LoadRoutes() error {
	var routes []models.Route
	err := config.DB.Select(&routes, "SELECT * FROM routes WHERE enabled = ?", true)
	if err != nil {
		return err
	}

	rm.mu.Lock()
	defer rm.mu.Unlock()

	for _, route := range routes {
		rm.addRoute(route)
	}
	return nil
}

// addRoute 添加单个路由
func (rm *RouteManager) addRoute(route models.Route) {
	rm.router.GET(route.Path, func(c *gin.Context) {
		c.HTML(http.StatusOK, route.Template, gin.H{
			"Title": route.Title,
		})
	})
}

// UpdateRoute 更新路由配置
func (rm *RouteManager) UpdateRoute(route models.Route) error {
	now := time.Now()
	if route.ID == 0 {
		// 新建路由
		route.CreatedAt = now
		route.UpdatedAt = now
		query := `
			INSERT INTO routes (path, template, title, enabled, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`
		result, err := config.DB.Exec(query,
			route.Path, route.Template, route.Title, route.Enabled,
			route.CreatedAt, route.UpdatedAt)
		if err != nil {
			return err
		}
		route.ID, _ = result.LastInsertId()
	} else {
		// 更新路由
		route.UpdatedAt = now
		query := `
			UPDATE routes 
			SET path = ?, template = ?, title = ?, enabled = ?, updated_at = ?
			WHERE id = ?
		`
		_, err := config.DB.Exec(query,
			route.Path, route.Template, route.Title, route.Enabled,
			route.UpdatedAt, route.ID)
		if err != nil {
			return err
		}
	}

	rm.mu.Lock()
	defer rm.mu.Unlock()

	if route.Enabled {
		rm.addRoute(route)
	}
	return nil
}

// GetRoutes 获取所有路由配置
func (rm *RouteManager) GetRoutes() ([]models.Route, error) {
	var routes []models.Route
	err := config.DB.Select(&routes, "SELECT * FROM routes")
	return routes, err
}

// GetRoute 获取单个路由配置
func (rm *RouteManager) GetRoute(id int64) (models.Route, error) {
	var route models.Route
	err := config.DB.Get(&route, "SELECT * FROM routes WHERE id = ?", id)
	return route, err
}
