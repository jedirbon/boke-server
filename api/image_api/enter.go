package image_api

import (
	"boke-server/models"
)

type ImageApi struct {
}

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
}
