package date_utils

import "time"

const (
	DateLayout = "02/01/2006"
)

//Parse comprueba que la fecha pasada es del formato 01/02/2006
func Parse(date string) (string, error) {
	dateParsed, err := time.Parse(DateLayout, date)
	if err != nil {
		return "", err
	}
	return dateParsed.Format(DateLayout), nil
}

//CheckAgeAndBirthCoherence comprueba si la edad y fecha de nacimiento pasada son coherentes
func CheckAgeAndBirthCoherence(age int, birthDate string) bool {
	date, _ := time.Parse(DateLayout, birthDate)
	birthDatePlusAge := date.AddDate(age, 0, 0)
	//Comprobamos que la edad más la fecha de nacimiento no sea una fecha futura
	if birthDatePlusAge.After(time.Now()) {
		return false
	}
	difference := time.Now().Sub(birthDatePlusAge)
	//Comprobamos que la diferencia entre la fecha de nacimiento y nacimiento+edad no sea más de un año
	if int64(difference.Hours()/24/365) > 0 {
		return false
	}
	return true
}
