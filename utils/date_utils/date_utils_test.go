package date_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	date, err := Parse("fail")
	assert.Equal(t, "", date)
	assert.Error(t, err)

	date, err = Parse("08/10/2017")
	assert.Equal(t, "08/10/2017", date)
	assert.Nil(t, err)
}

func TestCheckAgeAndBirthCoherence(t *testing.T) {
	result := CheckAgeAndBirthCoherence(1, "01/01/2021")
	assert.True(t, result)

	result = CheckAgeAndBirthCoherence(10, "01/01/2021")
	assert.False(t, result)

	result = CheckAgeAndBirthCoherence(0, "01/01/2021")
	assert.False(t, result)
}
