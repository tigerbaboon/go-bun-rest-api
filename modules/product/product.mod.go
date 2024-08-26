package product

import "github.com/uptrace/bun"

type ProductModule struct {
	Ctl *ProductController
	Svc *ProductService
}

func New(db *bun.DB) *ProductModule {
	svc := newService(db)
	return &ProductModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
