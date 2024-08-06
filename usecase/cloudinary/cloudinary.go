package cloudinary

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/model"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryUseCase interface {
	SendingMail(payload model.PayloadMail)
}

func ImageUploadHelper(file multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.AppConfig.CloudName, config.AppConfig.CloudKey, config.AppConfig.ApiSecret)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{Folder: config.AppConfig.CloudFolder},
	)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
