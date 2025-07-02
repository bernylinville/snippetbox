package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home 首页处理器
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 添加响应头
	w.Header().Add("Server", "Go")

	// 模板文件列表，基础模板必须在第一位
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// 解析模板文件
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Println(err.Error())

		// 处理模板解析错误

		app.serverError(w, r, err)
		return
	}

	// 执行模板渲染
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// 处理模板执行错误

		app.serverError(w, r, err)
		return
	}
}

// snippetView 查看代码片段处理器
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// 提取并验证 ID 参数
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// 显示指定 ID 的代码片段

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate 创建代码片段表单处理器
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost 处理创建代码片段请求
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// 设置 201 状态码
	w.WriteHeader(http.StatusCreated) // w.WriteHeader(201)

	// 写入响应内容
	w.Write([]byte("Save a new snippet..."))
}
