package countersign_type

import (
	"eirc.app/internal/v1/service/countersign_type"
	model "eirc.app/internal/v1/structure/countersign_type"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	CountersignTypeService countersign_type.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		CountersignTypeService: countersign_type.New(db),
	}
}
