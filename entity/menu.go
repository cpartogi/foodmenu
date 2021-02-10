package entity

type MenuType struct {
	MenuTypeId   int    `json:"menu_type_id"`
	MenuTypeName string `validate:"required" json:"menu_type_name"`
}

type Menu struct {
	MenuId      string `validate:"required" json:"menu_id"`
	MenuTypeId  int    `validate:"required" json:"menu_type_id"`
	MenuName    string `validate:"required" json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `validate:"required,number" json:"menu_price"`
}

type MenuRestaurant struct {
	MenuRestId string `validate:"required" json:"menu_rest_id"`
	MenuId     string `validate:"required" json:"menu_id"`
	WartegId   string `validate:"required" json:"warteg_id"`
}
