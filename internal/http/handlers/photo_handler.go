package handlers

import (
	"socia/config"
	"socia/models"

	"github.com/gofiber/fiber/v2"
)

// Get all photos
func GetPhotos(c *fiber.Ctx) error {
	rows, err := config.DB.Query("SELECT id, title, caption, photo_url, user_id FROM photos")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	var photos []models.Photo
	for rows.Next() {
		var photo models.Photo
		if err := rows.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan row"})
		}
		photos = append(photos, photo)
	}

	return c.JSON(photos)
}

// Upload photo
func UploadPhoto(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64) // Ambil user_id dari JWT

	photo := new(models.Photo)
	if err := c.BodyParser(photo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec("INSERT INTO photos (title, caption, photo_url, user_id) VALUES ($1, $2, $3, $4)",
		photo.Title, photo.Caption, photo.PhotoURL, int(userID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload photo"})
	}

	return c.JSON(fiber.Map{"message": "Photo uploaded successfully"})
}

// Edit photo (hanya oleh pemilik)
func EditPhoto(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)
	photoID := c.Params("id")

	var photo models.Photo
	if err := c.BodyParser(&photo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Pastikan foto hanya bisa diedit oleh pemiliknya
	res, err := config.DB.Exec("UPDATE photos SET title=$1, caption=$2, photo_url=$3 WHERE id=$4 AND user_id=$5",
		photo.Title, photo.Caption, photo.PhotoURL, photoID, int(userID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update photo"})
	}

	// Pastikan ada row yang terupdate
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized or photo not found"})
	}

	return c.JSON(fiber.Map{"message": "Photo updated successfully"})
}

// Delete photo (hanya oleh pemilik)
func DeletePhoto(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)
	photoID := c.Params("id")

	// Hapus hanya jika user adalah pemiliknya
	res, err := config.DB.Exec("DELETE FROM photos WHERE id=$1 AND user_id=$2", photoID, int(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete photo"})
	}

	// Pastikan ada row yang terhapus
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized or photo not found"})
	}

	return c.JSON(fiber.Map{"message": "Photo deleted successfully"})
}
