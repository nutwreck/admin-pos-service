package constants

// GENERAL
var (
	TRUE_VALUE  = true
	FALSE_VALUE = false

	EMPTY_VALUE  = ""
	EMPTY_NUMBER = 0

	ROLE_SYS  = "sys"
	ROLE_USER = "user"
)

// ROLE TYPE
/* Untuk membedakan role untuk internal (sys) dan external (user) */
type RoleType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var RoleTypes = []RoleType{
	{ID: "sys", Name: "sys", Description: "Digunakan untuk mengelompokkan role yang digunakan internal sistem"},
	{ID: "user", Name: "user", Description: "Digunakan untuk mengelompokkan role yang digunakan user"},
}
