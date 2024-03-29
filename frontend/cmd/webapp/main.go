package main

import (
	"mime"
	"net/http"

	"github.com/charukak/todo-app-htmx/frontend/internal/handlers"
	"github.com/charukak/todo-app-htmx/frontend/pkg/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	loadMimeTypes()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("web/static/out"))

	r.Handle("/*", fs)

	h := handlers.NewHandler()

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", h.TodoList)
		r.Post("/", h.CreateTodo)
		r.Put("/{id}", h.UpdateTodoStatus)
		r.Delete("/{id}", h.DeleteTodoById)
	})

	s := server.NewServer()
	s.Start(r)

}

func loadMimeTypes() {
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")
}
