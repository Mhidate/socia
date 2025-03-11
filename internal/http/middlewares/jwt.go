package middlewares

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := strings.TrimSpace(c.Get("Authorization"))
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	// Pastikan header menggunakan format "Bearer <token>"
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Println("⚠ WARNING: JWT_SECRET tidak ditemukan di environment!")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server misconfiguration"})
	}

	// Verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan token menggunakan metode HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("⚠ Invalid signing method:", token.Header["alg"])
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("⚠ Invalid token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Ambil user_id dari claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		log.Println("⚠ Invalid claims:", claims)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid claims"})
	}

	userID, ok := claims["user_id"]
	if !ok {
		log.Println("⚠ user_id tidak ditemukan di claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid claims"})
	}

	var userIDInt int
	switch v := userID.(type) {
	case float64:
		userIDInt = int(v)
	case int:
		userIDInt = v
	default:
		log.Println("⚠ Invalid user_id type:", v)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id type"})
	}

	c.Locals("user_id", userIDInt)

	return c.Next()
}
