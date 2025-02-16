package model

// User representa la estructura del usuario en el sistema
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}
