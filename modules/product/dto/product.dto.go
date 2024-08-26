package productdto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
	IsActived   bool    `json:"isActived"`
}

type UpdateProductRequest struct {
	CreateProductRequest
}

// change " " to " back tick " in ID
type GetProductByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// change " " to " back tick " in ID
type GetProductListRequest struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	SortBy   string `form:"sortBy"`
	OrderBy  string `form:"orderBy"`
	Search   string `form:"search"`
	SearchBy string `form:"searchBy"`
}
