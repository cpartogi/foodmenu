package response

type SwaggerMenuType struct {
	Base
	Data []DataMenuType `json:"data"`
}

type DataMenuType struct {
	MenuTypeId   int    `json:"menu_type_id"`
	MenuTypeName string `json:"menu_type_name"`
}
