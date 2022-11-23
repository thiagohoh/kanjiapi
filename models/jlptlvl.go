package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Jlptlvl struct {
	gorm.Model
	Level  string  `json:"level" validate:"nonzero"`
	Kanjis []Kanji `json:"kanjis"`
}

func JlptValidation(j *Jlptlvl) error {

	if err := validator.Validate(j); err != nil {
		return err
	}

	return nil
}
