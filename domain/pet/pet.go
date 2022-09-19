package pet

import (
	"github.com/montanaflynn/stats"
	"net/http"
	"pet-microservice/utils/date_utils"
	"pet-microservice/utils/rest_errors"
	"strings"
)

type Pet struct {
	Nombre          string `json:"nombre"`
	Especie         string `json:"especie"`
	Genero          string `json:"genero"`
	Edad            int    `json:"edad"`
	FechaNacimiento string `json:"fecha_nacimiento"`
}

type Kpi struct {
	Especie           string  `json:"especie"`
	AverageAge        float64 `json:"edad_media"`
	StandardDeviation float64 `json:"desviacion_estandar"`
}

type GroupOfPets []Pet

func (p *Pet) Validate() rest_errors.RestError {
	p.Nombre = strings.TrimSpace(p.Nombre)
	if p.Nombre == "" {
		return rest_errors.New(http.StatusBadRequest, "nombre is required")
	}
	p.Especie = strings.TrimSpace(p.Especie)
	if p.Especie == "" {
		return rest_errors.New(http.StatusBadRequest, "especie is required")
	}
	p.Genero = strings.TrimSpace(p.Genero)
	if p.Genero == "" {
		return rest_errors.New(http.StatusBadRequest, "genero is required")
	}
	if p.Genero != "M" && p.Genero != "H" {
		return rest_errors.New(http.StatusBadRequest, "genero must be M or H")
	}
	if p.Edad < 0 {
		return rest_errors.New(http.StatusBadRequest, "edad can't be negative")
	}
	var err error
	p.FechaNacimiento, err = date_utils.Parse(p.FechaNacimiento)
	if err != nil {
		return rest_errors.New(http.StatusBadRequest, "fecha_nacimiento invalid date format. Format must be 01/02/2000")
	}
	if !date_utils.CheckAgeAndBirthCoherence(p.Edad, p.FechaNacimiento) {
		return rest_errors.New(http.StatusBadRequest, "edad and fecha_nacimiento are not coherent")
	}
	return nil
}

func (g GroupOfPets) GetKpi() (Kpi, rest_errors.RestError) {
	var kpi Kpi
	var agesSlice stats.Float64Data
	for _, pet := range g {
		agesSlice = append(agesSlice, float64(pet.Edad))
	}
	mean, err := agesSlice.Mean()
	if err != nil {
		return kpi, rest_errors.New(http.StatusInternalServerError, "error getting average age")
	}
	kpi.AverageAge = mean

	sd, err := agesSlice.StandardDeviation()
	if err != nil {
		return kpi, rest_errors.New(http.StatusInternalServerError, "error getting standard deviation")
	}
	kpi.StandardDeviation = sd
	return kpi, nil
}
