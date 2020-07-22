package userdefinetype

import (
	"errors"
)

type Date struct {
	day   int
	month int
	year  int
}

func (value *Date) SetDay(day int) error {
	if day > 31 || day < 1 {
		return errors.New("Invalid day")
	}
	value.day = day
	return nil
}

func (value *Date) SetMonth(month int) error {
	if month > 12 || month < 1 {
		return errors.New("Invalid month")
	}
	value.month = month
	return nil
}

func (value *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year")
	}

	value.year = year
	return nil
}
