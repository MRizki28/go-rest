package mahasiswacontroller

import (
	"net/http"

	"github.com/MRizki28/go-rest/models"
	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	models.DB.Find(&mahasiswa)

	if mahasiswa == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Success get All Data",
			"data":    mahasiswa,
		})
	}

}

func GetDataById(c *gin.Context) {

}
