package main

import (
	"github.com/MRizki28/go-rest/controllers/MahasiswaController"
	"github.com/MRizki28/go-rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r :=	gin.Default();
	models.ConnectionDatabase()
	models.Migrate()

	r.GET("/api/mahasiswa" , mahasiswacontroller.GetAllData)
	r.GET("/api/mahasiswa/:id" , mahasiswacontroller.GetDataById)

	r.Run()
}