package configs

import "github.com/nutwreck/admin-pos-service/constants"

var RoleConfig = map[string]bool{
	"superadmin": constants.TRUE_VALUE,
	"admin":      constants.TRUE_VALUE,
}
