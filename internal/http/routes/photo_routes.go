package routes

import (
	"socia/internal/http/handlers"
	"socia/internal/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PhotoRoutes(app *fiber.App) {
	app.Get("/photos", handlers.GetPhotos) // Public: Semua bisa lihat foto

	photo := app.Group("/photos", middlewares.JWTMiddleware)
	photo.Post("/", handlers.UploadPhoto)      // Private: Hanya user yang bisa upload
	photo.Put("/:id", handlers.EditPhoto)      // Private: Hanya pemilik bisa edit
	photo.Delete("/:id", handlers.DeletePhoto) // Private: Hanya pemilik bisa hapus
}
