package services

import (
	"errors"
	"time"
)

func CalculateDateInterval(startDate, endDate string) (int, error) {
	//layout da data
	const layout = "02/01/2006" //DD/MM/YYYY

	startDateParsed, err1 := time.Parse(layout, startDate)
	endDateParsed, err2 := time.Parse(layout, endDate)

	if err1 != nil || err2 != nil {
		return 0, errors.New("Invalid date format")
	}

	difference := endDateParsed.Sub(startDateParsed)

	//converte a diferença em dias
	return int(difference.Hours() / 24), nil
}
