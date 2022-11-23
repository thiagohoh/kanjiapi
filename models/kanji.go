package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Kanji struct {
	gorm.Model
	Reading   string `json:"reading" validade:"nonzero"`
	Meaning   string `json:"meaning" validate:"nonzero"`
	Kanji     string `json:"kanji" validade:"nonzero"`
	JlptlvlID uint
}

func KanjiValidation(kanji *Kanji) error {
	if err := validator.Validate(kanji); err != nil {
		return err
	}

	return nil
}
