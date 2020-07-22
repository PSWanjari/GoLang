package userdefinetype

import (
	"errors"
)

type Date struct {
	day   int
	month int
	year  int
}

//---------Setters--------------------
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

//----------getters-----------
func (value *Date) Day() int {
	return value.day
}
func (value *Date) Month() int {
	return value.month
}
func (value *Date) Year() int {
	return value.year
}
