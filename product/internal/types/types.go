// Code generated by goctl. DO NOT EDIT.
package types

type CategroyProductListRequest struct {
	Name string `json:"name, optional"`
}

type CategroyProductListResponse struct {
	BaseResponse
	Data []*ProductBasic `json:"data"`
}

type GetProductRequest struct {
	Id uint `json:"id, optional"`
}

type GetProductResponse struct {
	BaseResponse
	Data Product `json:"data"`
}

type GetOnePerCategoryRequest struct {
}

type GetOnePerCategoryResponse struct {
	CategroyProductListResponse
}

type GetProductListRequest struct {
}

type GetProductListResponse struct {
	BaseResponse
	Data []*Product `json:"data"`
}

type Product struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	Img           string  `json:"image"`
	Price         float64 `json:"price"`
	Origin        string  `json:"origin"`
	Brand         string  `json:"brand"`
	Specification string  `json:"specification"`
	ShelfLife     string  `json:"shelflife"`
	Description   string  `json:"description"`
	Count         int     `json:"count"`
	Type          string  `json:"type"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Location      string  `json:"location"`
}

type ProductBasic struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image" gorm:"column:img"`
}

type BaseResponse struct {
	Status  string `json:"status"`
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
