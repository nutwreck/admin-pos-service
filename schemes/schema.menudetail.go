package schemes

type MenuDetail struct {
	ID     string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuID string `json:"menu_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"master produk"`
	Link    string `json:"link"  validate:"required" example:"/master-produk"`
	Image   string `json:"image"`
	Icon    string `json:"icon"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetMenuDetail struct {
	ID       string `json:"id"`
	MenuID   string `json:"menu_id"`
	MenuName string `json:"menu_name"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Image    string `json:"image"`
	Icon     string `json:"icon"`
	Active   *bool  `json:"active"`
}
