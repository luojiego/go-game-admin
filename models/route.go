package models

import "time"

// Route 定义路由配置模型
type Route struct {
	ID        int64     `db:"id" json:"id"`                 // 路由ID
	Path      string    `db:"path" json:"path"`             // 路由路径
	Template  string    `db:"template" json:"template"`     // 模板名称
	Title     string    `db:"title" json:"title"`           // 页面标题
	Enabled   bool      `db:"enabled" json:"enabled"`       // 是否启用
	CreatedAt time.Time `db:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"` // 更新时间
}
