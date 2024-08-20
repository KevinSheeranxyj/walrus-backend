package walrus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadBlob(address string, epochs string, blob []byte) (string, error) {

	r := gin.Default()

	storeEndpoint := fmt.Sprintf("http://%s/store", address)

	r.PUT(storeEndpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run(address)
	return "", nil
}
