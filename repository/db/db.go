package db

import (
	"net/http"
	"pet-microservice/domain/pet"
	"pet-microservice/utils/rest_errors"
)

//La base de datos será una base de datos en memoria
type petDb struct {
	petMap        map[string]pet.Pet
	populationMap map[string]int
}

type PetDb interface {
	InsertPet(pet.Pet) rest_errors.RestError
	GetPetsByEspecie(string) (pet.GroupOfPets, rest_errors.RestError)
	GetAllPets() ([]pet.Pet, rest_errors.RestError)
	GetMostPopularEspecie() (string, rest_errors.RestError)
}

//New construye el controlador de la base de datos
func New() PetDb {
	petMap := make(map[string]pet.Pet)
	populationMap := make(map[string]int)
	return &petDb{petMap: petMap, populationMap: populationMap}
}

//InsertPet inserta una nueva mascota en la base de datos
func (p *petDb) InsertPet(pet pet.Pet) rest_errors.RestError {
	p.petMap[pet.Nombre] = pet
	value, exist := p.populationMap[pet.Especie]
	if !exist {
		p.populationMap[pet.Especie] = 1
	} else {
		p.populationMap[pet.Especie] = value + 1
	}
	return nil
}

//GetAllPets obtiene todos las mascotas presentes en la base de datos
func (p *petDb) GetAllPets() ([]pet.Pet, rest_errors.RestError) {
	var resultSlice []pet.Pet
	for _, pet := range p.petMap {
		resultSlice = append(resultSlice, pet)
	}
	if len(resultSlice) == 0 {
		return nil, rest_errors.New(http.StatusNotFound, "no pets found")
	}
	return resultSlice, nil
}

//GetPetsByEspecie obtiene todas las mascotas de la especie dada
func (p *petDb) GetPetsByEspecie(especie string) (pet.GroupOfPets, rest_errors.RestError) {
	var resultSlice []pet.Pet
	for _, pet := range p.petMap {
		if pet.Especie == especie {
			resultSlice = append(resultSlice, pet)
		}
	}
	if len(resultSlice) == 0 {
		return nil, rest_errors.New(http.StatusNotFound, "no pets found for the given especie")
	}
	return resultSlice, nil
}

//GetMostPopularEspecie obtiene la especie más repetida en la base de datos
func (p *petDb) GetMostPopularEspecie() (string, rest_errors.RestError) {
	var mostPopularEspecie string
	var biggestPopulation int
	for especie, population := range p.populationMap {
		if population > biggestPopulation {
			mostPopularEspecie = especie
			biggestPopulation = population
		}
	}
	return mostPopularEspecie, nil
}
