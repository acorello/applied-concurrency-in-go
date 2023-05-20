package fixtures

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/applied-concurrency-in-go/models/product"
)

const productInputPath string = "./input/products.csv"

// importProducts imports the start position of the products DB
func ImportProducts(addProduct func(product.Id, product.Product)) error {
	const expectedFields = 5
	productRecords, err := readCsv(productInputPath)
	if err != nil {
		return err
	}

	for _, productRecord := range productRecords {
		if len(productRecord) != expectedFields {
			continue
		}
		id := product.Id(productRecord[0])
		stock, err := strconv.Atoi(productRecord[2])
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(productRecord[4], 64)
		if err != nil {
			continue
		}
		addProduct(id, product.Product{
			ID:    id,
			Name:  fmt.Sprintf("%s(%s)", productRecord[1], productRecord[3]),
			Stock: stock,
			Price: price,
		})
	}
	return nil
}

// the format of the csv file is hardcoded so we can take some
// error handling liberties for the sake of brevity
func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
