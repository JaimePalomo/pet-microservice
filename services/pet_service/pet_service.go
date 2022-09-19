package pet_service

import (
	"pet-microservice/domain/pet"
	"pet-microservice/repository/db"
	"pet-microservice/utils/rest_errors"
)

type petService struct {
	petDb db.PetDb
}

type PetService interface {
	InsertPet(pet pet.Pet) (pet.Pet, rest_errors.RestError)
	GetAllPets() ([]pet.Pet, rest_errors.RestError)
	GetKpiOfEspecie(string) (pet.Kpi, rest_errors.RestError)
	GetKpiMostPopularEspecie() (pet.Kpi, rest_errors.RestError)
}

//New inicia el controlador del servicio
func New(petDb db.PetDb) PetService {
	return &petService{petDb: petDb}
}

//InsertPet inserta la mascota en la base de datos usando el controlado de la bbdd
func (p *petService) InsertPet(givenPet pet.Pet) (pet.Pet, rest_errors.RestError) {
	err := givenPet.Validate()
	if err != nil {
		return pet.Pet{}, err
	}
	err = p.petDb.InsertPet(givenPet)
	if err != nil {
		return pet.Pet{}, err
	}
	return givenPet, nil
}

//GetAllPets obtiene todos las mascotas
func (p *petService) GetAllPets() ([]pet.Pet, rest_errors.RestError) {
	return p.petDb.GetAllPets()
}

//GetKpiOfEspecie obtiene todos las mascotas de una especie de la bbdd y calcula su edad media y deviación típica
func (p *petService) GetKpiOfEspecie(especie string) (pet.Kpi, rest_errors.RestError) {
	petsOfEspecie, err := p.petDb.GetPetsByEspecie(especie)
	if err != nil {
		return pet.Kpi{}, err
	}
	kpi, err := petsOfEspecie.GetKpi()
	if err != nil {
		return pet.Kpi{}, err
	}
	kpi.Especie = especie
	return kpi, nil
}

//GetKpiMostPopularEspecie obtiene edad media y deviación típica de la especie más repetida
func (p *petService) GetKpiMostPopularEspecie() (pet.Kpi, rest_errors.RestError) {
	especie, err := p.petDb.GetMostPopularEspecie()
	if err != nil {
		return pet.Kpi{}, err
	}
	return p.GetKpiOfEspecie(especie)
}
