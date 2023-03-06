package storage

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/xuri/excelize/v2"
	"time"

	//"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
)

func InitStorage() *cloudinary.Cloudinary {
	// Add your Cloudinary credentials.
	cld, _ := cloudinary.NewFromParams("logicielapplab", "667244167665823", "zA5hDDm8aykVffsUJvwmEZlEleE")
	return cld
}
func UploadAsset(cld *cloudinary.Cloudinary, file *excelize.File, date time.Time) string {
	// Upload the my_picture.jpg image and set the PublicID to "my_image".
	buff, _ := file.WriteToBuffer()

	resp, err := cld.Upload.Upload(context.TODO(), buff, uploader.UploadParams{PublicID: fmt.Sprintf("Reporte Diario %s.xlsx", date.String()[:10]), ResourceType: "raw", Folder: "GaviotaFerry/Reports"})
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return "No se pudo generar reporte"
	} else {
		return resp.SecureURL
	}
}
