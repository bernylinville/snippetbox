package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	// 导入我们创建的 models 包
	"snippetbox.xmxxmx.us/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// application 应用程序结构体，用于保存全局依赖项
// 包含结构化日志器和代码片段模型
// Add a templateCache field to the application struct.
type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// 定义命令行参数 addr，默认值为 ":4000"
	addr := flag.String("addr", ":4000", "HTTP network address")
	// 定义 MySQL 数据源名称参数
	dsn := flag.String("dsn", "snippetbox:Mz8nQ3vR7sT2uW5yE9aF4bG6cH1jL0kP@/snippetbox?parseTime=true", "MySQL data source name")
	// 解析命令行参数，必须在使用参数前调用
	flag.Parse()

	// 初始化结构化日志器
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// 创建数据库连接池
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// 延迟关闭数据库连接池
	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// 初始化应用程序实例，包含依赖项
	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// 路由配置已移至 routes() 方法

	// 记录服务器启动信息
	logger.Info("Starting server", "addr", *addr)

	// 启动 HTTP 服务器
	err = http.ListenAndServe(*addr, app.routes())
	// 记录错误并退出
	logger.Error(err.Error())
	os.Exit(1)
}

// openDB 创建并返回数据库连接池
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
