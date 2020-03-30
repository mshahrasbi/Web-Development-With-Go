// 1. Parse this CSV file, putting two fields from the contents of the CSV
//	file into a data structure.
// 2. Parse an html template, then pass the data from step 1 into the CSV
//	template; have the html template display the CSV data in a web page.

package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// comments
type Record struct {
	Date time.Time
	Open float64
}

func processCSV(res http.ResponseWriter, req *http.Request) {
	records := prs("table.csv")

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []Record {

	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records
}

func main() {
	http.HandleFunc("/", processCSV)
	http.ListenAndServe(":8080", nil)
}
