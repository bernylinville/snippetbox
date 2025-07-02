package models

import (
	"database/sql"
	"time"
)

// Snippet 定义代码片段结构体，用于存储单个代码片段的数据
// 结构体字段与 MySQL 数据库中 snippets 表的字段一一对应
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel 定义代码片段模型结构体，封装数据库连接池
type SnippetModel struct {
	DB *sql.DB
}

// Insert 向数据库插入新的代码片段
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get 根据 ID 获取指定的代码片段
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Latest 获取最新创建的 10 个代码片段
func (m *SnippetModel) Latest(id int) ([]Snippet, error) {
	return nil, nil
}
