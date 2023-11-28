package service

import (
	"encoding/csv"
	"fin_test/model"
	"os"
	"log"
	"fmt"
	"time"
	"strconv"
)

func GetTickInfo(code string, year, month, day, hour int) (*model.Tick, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	file, err := os.Open("order_books.csv") 
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	startAt := time.Date(year, time.Month(month), day, hour, 0, 0, 0, jst)
	fmt.Println(startAt)

	var tickValues []int
	for i, v := range rows {
		if  i == 0 {
			continue
		}
		parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", v[0])
		if  err != nil {
			log.Fatal(err)
		}

		if parsedTime.After(startAt) && parsedTime.Before(startAt.Add(1 * time.Hour)) && v[1] == code {
			value, _ :=  strconv.Atoi(v[2])
		tickValues = append(tickValues, value)
		}
	}

	tick := model.TickNomarize(tickValues)
	return tick, nil
}
