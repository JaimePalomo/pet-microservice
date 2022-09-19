package pet

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPet_Validate(t *testing.T) {
	pet := Pet{}

	err := pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "nombre is required", err.Message())

	pet.Nombre = "Tobi"
	err = pet.Validate()
	assert.Equal(t, err.Status(), http.StatusBadRequest)
	assert.Equal(t, "especie is required", err.Message())

	pet.Especie = "perro"
	err = pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "genero is required", err.Message())

	pet.Genero = "fail"
	err = pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "genero must be M or H", err.Message())

	pet.Genero = "H"
	pet.Edad = -3
	err = pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "edad can't be negative", err.Message())

	pet.Edad = 10
	pet.FechaNacimiento = "fail"
	err = pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "fecha_nacimiento invalid date format. Format must be 01/02/2000", err.Message())

	pet.FechaNacimiento = "01/01/2000"
	err = pet.Validate()
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "edad and fecha_nacimiento are not coherent", err.Message())

	pet.FechaNacimiento = "01/01/2012"
	err = pet.Validate()
	assert.Nil(t, err)

}
