package models

// Error representa una estructura de error genérica para la API.
type Error struct {
	Code    int    `json:"code"`    // Código del error (por ejemplo, 400, 404)
	Message string `json:"message"` // Mensaje descriptivo del error
}
