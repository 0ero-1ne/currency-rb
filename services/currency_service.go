package services

import (
	"currency/db"
	"currency/models"
	"log"
)

type CurrencyService struct{}

func NewCurrencyService() CurrencyService {
	return CurrencyService{}
}

func (service *CurrencyService) FindAll() []models.Currency {
	var currencies []models.Currency
	db.DB().Find(&currencies)
	return currencies
}

func (service *CurrencyService) FindAllByDay(day string) []models.Currency {
	var currencies []models.Currency
	db.DB().Where("date like ?", day+"%").Find(&currencies)
	return currencies
}

func (service *CurrencyService) SaveMany(currencies []*models.Currency) {
	result := db.DB().Create(&currencies)

	if result.Error != nil {
		log.Printf("Can not save the currencies: %v", result.Error.Error())
	} else {
		log.Printf("%d rows were saved", result.RowsAffected)
	}
}
