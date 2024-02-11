package services

import (
	"github.com/harranali/authority"
	"gorm.io/gorm"
)

func InitAuthority(db *gorm.DB) {
	auth := authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           db,
	})
	Authority = auth
}

var Authority *authority.Authority
