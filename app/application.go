package app

import (
	"github.com/gin-gonic/gin"
	"pet-microservice/http/http_pets"
	"pet-microservice/http/http_ping"
	"pet-microservice/repository/db"
	"pet-microservice/services/pet_service"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbPets := db.New()
	petsService := pet_service.New(dbPets)
	petsHandler := http_pets.New(petsService)

	router.POST("/creamascota", petsHandler.CreatePet)
	router.GET("/kpidemascotas", petsHandler.GetKpi)
	router.GET("/lismascotas", petsHandler.GetAllPets)

	router.GET("/ping", http_ping.Pong)

	router.Run(":8080")

}
