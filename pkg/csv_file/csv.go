package csv_file

import (
	"encoding/csv"
	"log"
	"os"
	"web_scraper/internal/scraping/domain/model"
)

type CsvFile struct {
	CsvFile *os.File
}

func New() *CsvFile {
	file, err := os.Create("./result.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvFile := CsvFile{CsvFile: file}
	return &csvFile
}

func (c *CsvFile) WriteContents(allRows []model.Content, path *string) {
	if path != nil {
		var err error
		c.CsvFile, err = os.Open(*path)
		if err != nil {
			log.Fatalln(err)
		}
	}
	csvWriter := csv.NewWriter(c.CsvFile)
	for _, c := range allRows {
		csvWriter.Write(c.ToSlice())
	}
	csvWriter.Flush()
}

func (c *CsvFile) Close() error {
	return c.CsvFile.Close()
}
