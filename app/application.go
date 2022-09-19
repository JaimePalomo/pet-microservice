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

//StartApplication inicia los controladores necesarios y el servidor
func StartApplication() {
	//Inicio del controlador de la base de datos
	dbPets := db.New()
	//Inicio del controlador del servicio
	petsService := pet_service.New(dbPets)
	//Inicio del controlador de los handlers para los endpoints
	petsHandler := http_pets.New(petsService)

	//Endpoints requeridos
	router.POST("/creamascota", petsHandler.CreatePet)
	router.GET("/kpidemascotas", petsHandler.GetKpi)
	router.GET("/lismascotas", petsHandler.GetAllPets)

	//Endpoint para comprobar el buen funcionamiento del microservicio
	router.GET("/ping", http_ping.Pong)

	router.Run(":8080")

}
