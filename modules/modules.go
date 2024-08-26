package modules

import (
	"app/config"
	"app/modules/product"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB      *bun.DB
	Product *product.ProductModule
}

func Get() *Modules {

	db := config.Database()

	return &Modules{
		DB:      db,
		Product: product.New(db),
	}
}
