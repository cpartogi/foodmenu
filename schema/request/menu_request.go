package request

type Menu struct {
	MenuTypeId  int    `validate:"required,number" json:"menu_type_id"`
	WartegId    string `validate:"required" json:"warteg_id"`
	MenuName    string `validate:"required" json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `validate:"required,number" json:"menu_price"`
}

type MenuUpdate struct {
	MenuTypeId  int    `validate:"required,number" json:"menu_type_id"`
	WartegId    string `validate:"required" json:"warteg_id"`
	MenuName    string `validate:"required" json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `validate:"required,number" json:"menu_price"`
}
