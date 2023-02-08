package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

var (
	Trans ut.Translator
	My_DB *gorm.DB
)
