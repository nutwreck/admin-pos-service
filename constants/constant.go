package constants

// GENERAL
var (
	TRUE_VALUE  = true
	FALSE_VALUE = false

	EMPTY_VALUE  = ""
	EMPTY_NUMBER = 0
)

// ROLE USER
type RoleUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var RoleUsers = []RoleUser{
	{ID: "superadmin", Name: "SuperAdmin"},
	{ID: "admin", Name: "Admin"},
}

// JENIS KELAMIN
type JenisKelamin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var JenisKelamins = []JenisKelamin{
	{ID: 1, Name: "laki-laki"},
	{ID: 2, Name: "perempuan"},
}

// STATUS PERNIKAHAN
type StatusPernikahan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var StatusPernikahans = []StatusPernikahan{
	{ID: 1, Name: "belum menikah"},
	{ID: 2, Name: "menikah"},
}
