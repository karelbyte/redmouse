package db

import (
	"elpuertodigital/redmouse/models"

	"github.com/bxcodec/faker/v3"
)

func MeasureSeed() {

	measure := models.Measure{}

	for range [15]int{} {
		err := faker.FakeData(&measure)
		if err == nil {
			//Conect().Create(&measure)
		} 
	}	
}