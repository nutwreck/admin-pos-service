package schemes

type MenuDetailFunction struct {
	ID           string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuID       string `json:"menu_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID string `json:"menu_detail_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"add produk"`
	Link    string `json:"link"  validate:"required" example:"/add"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetMenuDetailFunction struct {
	ID             string `json:"id"`
	MenuID         string `json:"menu_id"`
	MenuName       string `json:"menu_name"`
	MenuDetailID   string `json:"menu_detail_id"`
	MenuDetailName string `json:"menu_detail_name"`
	Name           string `json:"name"`
	Link           string `json:"link"`
	Active         *bool  `json:"active"`
}

type MenuDetailFunctionRequest struct {
	MenuID       string `json:"menu_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID string `json:"menu_detail_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=200" example:"add produk"`
	Link   string `json:"link" validate:"required" example:"/add"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
