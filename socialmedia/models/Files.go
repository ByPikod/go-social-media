package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	FilePath string `json:"file_path"`
}
