package http_pets

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-microservice/domain/pet"
	"pet-microservice/services/pet_service"
	"pet-microservice/utils/rest_errors"
)

type petsHandler struct {
	petService pet_service.PetService
}

type PetsHandler interface {
	CreatePet(*gin.Context)
	GetAllPets(*gin.Context)
	GetKpi(*gin.Context)
}

//New inicia el controlador de los handlers de los endpoints
func New(petService pet_service.PetService) PetsHandler {
	return &petsHandler{petService: petService}
}

//CreatePet handler del endpoint /creamascota
func (p *petsHandler) CreatePet(c *gin.Context) {
	var pet pet.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		restErr := rest_errors.New(http.StatusBadRequest, "invalid json body")
		c.JSON(restErr.Status(), restErr.Message())
		return
	}

	result, err := p.petService.InsertPet(pet)
	if err != nil {
		c.JSON(err.Status(), err.Message())
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetAllPets handler del endpoint /lismascotas
func (p *petsHandler) GetAllPets(c *gin.Context) {
	pets, err := p.petService.GetAllPets()
	if err != nil {
		c.JSON(err.Status(), err.Message())
		return
	}
	c.JSON(http.StatusOK, pets)
}

//GetKpi handler del endpoint /kpidemascotas
func (p *petsHandler) GetKpi(c *gin.Context) {
	var err rest_errors.RestError
	var kpi pet.Kpi

	especie := getEspecieFromQuery(c)
	if especie == "" {
		kpi, err = p.petService.GetKpiMostPopularEspecie()
	} else {
		kpi, err = p.petService.GetKpiOfEspecie(especie)
	}
	if err != nil {
		c.JSON(err.Status(), err.Message())
		return
	}
	c.JSON(http.StatusOK, kpi)

}

//getEspecieFromQuery obtiene la especie obtenida por query
func getEspecieFromQuery(c *gin.Context) string {
	value, exist := c.GetQuery("especie")
	if !exist {
		return ""
	}
	return value
}
