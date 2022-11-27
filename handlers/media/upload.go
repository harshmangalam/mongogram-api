package media

import (
	"io"
	"mongogram/database"
	"mongogram/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Upload(c *fiber.Ctx) error {

	userId := c.Locals("userId")

	// parse request multipart form data and extract media
	file, err := c.FormFile("media")
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	bucket := database.Mi.Bucket
	filename := time.Now().Format("2006_01_02_150405") + "_" + file.Filename

	// get file chunks
	fileData, err := file.Open()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	// add meta data in gridfs files collection
	uploadOptions := options.GridFSUpload().SetMetadata(bson.M{"userId": userId})
	// upload data in gridfs  chunks collection
	bucketId, err := bucket.UploadFromStream(filename, io.Reader(fileData), uploadOptions)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return utils.ReturnSuccess(c, fiber.StatusOK, "Uploaded", fiber.Map{"bucketId": bucketId})

}
