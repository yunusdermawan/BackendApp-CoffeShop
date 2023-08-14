package pkg

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Cloudinary(file interface{}) (string, error) {
	name := "dsfrrcvrs"
	key := "554933634147368"
	secret := "dhVHQ_1QnUJ66EZvfEIHjZUI0jo"

	cld, _ := cloudinary.NewFromParams(name, key, secret)

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
