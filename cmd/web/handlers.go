package main

import (
	"errors"
	"fmt"
	"net/http"
	"snippetbox.xmxxmx.us/internal/models"
	"strconv"
)

// home 首页处理器
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 添加响应头
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// 模板文件列表，基础模板必须在第一位
	//files := []string{
	//	"./ui/html/base.tmpl",
	//	"./ui/html/partials/nav.tmpl",
	//	"./ui/html/pages/home.tmpl",
	//}

	// 解析模板文件
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	// log.Println(err.Error())
	//
	//	// 处理模板解析错误
	//
	//	app.serverError(w, r, err)
	//	return
	//}

	// 执行模板渲染
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	// 处理模板执行错误
	//
	//	app.serverError(w, r, err)
	//	return
	//}
}

// snippetView 查看代码片段处理器
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// 提取并验证 ID 参数
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the SnippetModel's Get() method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	snnipet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%+v", snnipet)
}

// snippetCreate 创建代码片段表单处理器
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost 处理创建代码片段请求
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
	// during development.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

	// 设置 201 状态码
	//w.WriteHeader(http.StatusCreated) // w.WriteHeader(201)

	// 写入响应内容
	//w.Write([]byte("Save a new snippet..."))
}
