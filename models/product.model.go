package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID          int64   `bun:",type:serial,autoincrement,pk"`
	Name        string  `bun:"name"`
	Price       float64 `bun:"price"`
	Description string  `bun:"description"`
	IsActived   bool    `bun:"is_actived,default:false"`
}
