package models

import "fmt"

type ImageModel struct {
	FileName string `gorm:"size:32"json:"fileName"'`
	Path     string `gorm:"size:256"json:"path"`
	Size     int64  `gorm:"size"json:"size"`
	Hash     string `gorm:"size:32"json:"hash"`
	Model
}

func (i ImageModel) WebPath() string {
	return fmt.Sprintf("/")
}
