package loader

import (
	"encoding/csv"
	"os"
	"strconv"
	"tickets/domain/model"
)

func LoadTickets() []model.Ticket {

	file, err := os.Open("data/tickets.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		panic(err)
	}

	tickets := make([]model.Ticket, len(records))
	for t, record := range records {
		price, err := strconv.Atoi(record[5])
		if err != nil {
			panic(err)
		}
		tickets[t] = model.Ticket{
			ID:      record[0],
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}
	}

	return tickets
}
