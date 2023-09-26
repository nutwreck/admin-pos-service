package schemes

type SchemeRoleMenu struct {
	ID                   string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	RoleID               string `json:"role_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuID               string `json:"menu_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID         string `json:"menu_detail_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailFunctionID string `json:"menu_detail_function_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"Add Produk"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type SchemeGetRoleMenu struct {
	ID                     string `json:"id"`
	RoleID                 string `json:"role_id"`
	RoleName               string `json:"role_name"`
	MenuID                 string `json:"menu_id"`
	MenuName               string `json:"menu_name"`
	MenuDetailID           string `json:"menu_detail_id"`
	MenuDetailName         string `json:"menu_detail_name"`
	MenuDetailFunctionID   string `json:"menu_detail_function_id"`
	MenuDetailFunctionName string `json:"menu_detail_function_name"`
	Name                   string `json:"name"`
	Active                 *bool  `json:"active"`
}

type SchemeRoleMenuRequest struct {
	RoleID               string `json:"role_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuID               string `json:"menu_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID         string `json:"menu_detail_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailFunctionID string `json:"menu_detail_function_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=200" example:"Add Produk"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
