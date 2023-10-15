package schemes

type User struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name string `json:"name" validate:"required,lowercase" example:"superadmin"`
	// Email is Unique
	Email      string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password   string `json:"password" validate:"required,gte=8" example:"12345678"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	RoleID     string `json:"role_id" validate:"uuid,required" example:"ca7af74f-2fbf-4dd7-b9bd-eba0903170d0"`
	Active     bool   `json:"active" validate:"required,boolean" example:"true"`
}

type GetUser struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Merchant GetMerchant     `json:"merchant"`
	Outlet   []GetOutletUser `json:"outlet"`
	Role     GetRole         `json:"role"`
	Active   bool            `json:"active"`
}

type UpdateUser struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Email is Unique
	Email string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	// Input with Lowercase
	Name         string `json:"name" validate:"required,lowercase" example:"superadmin"`
	OldPassword  string `json:"old_password" validate:"omitempty,gte=8" example:"12345678"`
	NewPassword  string `json:"new_password" validate:"omitempty,gte=8" example:"12345679"`
	DataPassword string `json:"data_password" validate:"gte=8" example:"12345678"`
	MerchantID   string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	RoleID       string `json:"role_id" validate:"uuid,required" example:"01f33858-0cf9-45eb-9e1f-c6a26ca759c4"`
	Active       *bool  `json:"active" validate:"required,boolean" example:"true"`
}

type UpdateUserExample struct {
	// Input with Lowercase
	Name        string `json:"name" validate:"required,lowercase" example:"superadmin"`
	OldPassword string `json:"old_password" validate:"omitempty,gte=8" example:"12345678"`
	NewPassword string `json:"new_password" validate:"omitempty,gte=8" example:"12345679"`
	MerchantID  string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	RoleID      string `json:"role_id" validate:"uuid,required" example:"01f33858-0cf9-45eb-9e1f-c6a26ca759c4"`
	Active      *bool  `json:"active" validate:"required,boolean" example:"true"`
}

type UserRequest struct {
	// Input with Lowercase
	Name string `json:"name" validate:"required,lowercase" example:"superadmin"`
	// Email is Unique
	Email      string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password   string `json:"password" validate:"required,gte=8" example:"12345678"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	RoleID     string `json:"role_id" validate:"uuid,required" example:"01f33858-0cf9-45eb-9e1f-c6a26ca759c4"`
	Active     bool   `json:"active" validate:"required,boolean" example:"true"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email" example:"pos.superadmin@digy.com" format:"email"`
	Password string `json:"password" validate:"required,gte=8" example:"12345678"`
}

type JWTConvert struct {
	ID       string
	Email    string
	Role     string
	Merchant string
}
