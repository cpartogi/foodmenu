package response

type MenuType struct {
	MenuTypeId   int    `json:"menu_type_id"`
	MenuTypeName string `json:"menu_type_name"`
}

type MenuAdd struct {
	MenuTypeId  int    `json:"menu_type_id"`
	WartegId    string `json:"warteg_id"`
	MenuName    string `json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `json:"menu_price"`
}

type MenuDelete struct {
	MenuId string `json:"menu_id"`
}

type MenuUpdate struct {
	MenuId      string `json:"menu_id"`
	MenuTypeId  int    `json:"menu_type_id"`
	WartegId    string `json:"warteg_id"`
	MenuName    string `json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `json:"menu_price"`
}

type MenuList struct {
	MenuId       string `json:"menu_id"`
	MenuTypeName string `json:"menu_type_name"`
	WartegId     string `json:"warteg_id"`
	MenuName     string `json:"menu_name"`
	MenuPrice    int    `json:"menu_price"`
}

type MenuDetail struct {
	MenuId       string `json:"menu_id"`
	MenuTypeName string `json:"menu_type_name"`
	WartegId     string `json:"warteg_id"`
	MenuName     string `json:"menu_name"`
	MenuDetail   string `json:"menu_detail"`
	MenuPicture  string `json:"menu_picture"`
	MenuPrice    int    `json:"menu_price"`
}
