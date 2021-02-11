package response

type SwaggerMenuType struct {
	Base
	Data []DataMenuType `json:"data"`
}

type DataMenuType struct {
	MenuTypeId   int    `json:"menu_type_id"`
	MenuTypeName string `json:"menu_type_name"`
}

type SwaggerMenuAdd struct {
	Base
	Data DataMenu `json:"data"`
}

type DataMenu struct {
	MenuTypeId  int    `json:"menu_type_id"`
	WartegId    string `json:"warteg_id"`
	MenuName    string `json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `json:"menu_price"`
}
