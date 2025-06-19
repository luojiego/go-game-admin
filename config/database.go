package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "routes.db")
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 创建路由表
	schema := `
	CREATE TABLE IF NOT EXISTS routes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT NOT NULL UNIQUE,
		template TEXT NOT NULL,
		title TEXT NOT NULL,
		enabled BOOLEAN DEFAULT true,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TRIGGER IF NOT EXISTS update_routes_timestamp 
		AFTER UPDATE ON routes
		BEGIN
			UPDATE routes SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
		END;
	`

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatal("Failed to create schema:", err)
	}
}
