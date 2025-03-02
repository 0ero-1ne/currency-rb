package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	CurrencyID           uint64  `json:"Cur_ID"`
	Date                 string  `json:"Date"`
	CurrencyAbbreviation string  `json:"Cur_Abbreviation"`
	CurrencyScale        uint64  `json:"Cur_Scale"`
	CurrencyName         string  `json:"Cur_Name"`
	CurrencyOfficialRate float64 `json:"Cur_OfficialRate"`
}
