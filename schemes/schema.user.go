package schemes

type SchemeUser struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	FirstName string `json:"first_name" validate:"required,lowercase" example:"super"`
	// Input with Lowercase
	LastName string `json:"last_name" validate:"required,lowercase" example:"admin"`
	// Email is Unique
	Email    string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password string `json:"password" validate:"required,gte=8" example:"12345678"`
	// Input superadmin or admin with lowercase
	Role   string `json:"role" validate:"required,lowercase" example:"superadmin"`
	Active bool   `json:"active" validate:"required,boolean" example:"true"`
}

type SchemeUpdateUser struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Email is Unique
	Email string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	// Input with Lowercase
	FirstName string `json:"first_name" validate:"required,lowercase" example:"super"`
	// Input with Lowercase
	LastName     string `json:"last_name" validate:"required,lowercase" example:"admin"`
	OldPassword  string `json:"old_password" validate:"omitempty,gte=8" example:"12345678"`
	NewPassword  string `json:"new_password" validate:"omitempty,gte=8" example:"12345679"`
	DataPassword string `json:"data_password" validate:"gte=8" example:"12345678"`
	// Input superadmin or admin with lowercase
	Role   string `json:"role" validate:"required,lowercase" example:"superadmin"`
	Active *bool  `json:"active" validate:"required,boolean" example:"true"`
}

type SchemeUpdateUserExample struct {
	// Input with Lowercase
	FirstName string `json:"first_name" validate:"required,lowercase" example:"super"`
	// Input with Lowercase
	LastName    string `json:"last_name" validate:"required,lowercase" example:"admin"`
	OldPassword string `json:"old_password" validate:"omitempty,gte=8" example:"12345678"`
	NewPassword string `json:"new_password" validate:"omitempty,gte=8" example:"12345679"`
	// Input superadmin or admin with lowercase
	Role   string `json:"role" validate:"required,lowercase" example:"superadmin"`
	Active *bool  `json:"active" validate:"required,boolean" example:"true"`
}

type SchemeUserRequest struct {
	// Input with Lowercase
	FirstName string `json:"first_name" validate:"required,lowercase" example:"super"`
	// Input with Lowercase
	LastName string `json:"last_name" validate:"required,lowercase" example:"admin"`
	// Email is Unique
	Email    string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password string `json:"password" validate:"required,gte=8" example:"12345678"`
	// Input superadmin or admin with lowercase
	Role   string `json:"role" validate:"required,lowercase" example:"superadmin"`
	Active bool   `json:"active" validate:"required,boolean" example:"true"`
}

type SchemeLoginUser struct {
	Email    string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password string `json:"password" validate:"required,gte=8" example:"12345678"`
}

type SchemeJWTConvert struct {
	ID    string
	Email string
	Role  string
}
