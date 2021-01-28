package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name   string `json:"product"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

func findExpensiveProduct(newProduct *Product, product *Product) {
	if newProduct.Price > product.Price {
		*product = *newProduct
		return
	}
	if newProduct.Price == product.Price && newProduct.Rating > product.Rating {
		*product = *newProduct
		return
	}
}

func readCSV(file *os.File, p *Product) {
	parser := csv.NewReader(file)
	if _, err := parser.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		product, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		Name := product[0]
		Price, err := strconv.Atoi(product[1])
		if err != nil {
			log.Fatal(err)
		}
		Rating, err := strconv.Atoi(product[2])
		if err != nil {
			log.Fatal(err)
		}
		findExpensiveProduct(&Product{Name, Price, Rating}, p)
	}
}

func readJSON(file *os.File, p *Product) {
	dec := json.NewDecoder(file)

	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		var newProduct Product
		err := dec.Decode(&newProduct)
		if err != nil {
			log.Fatal(err)
		}

		findExpensiveProduct(&newProduct, p)
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("One input file expected")
	}
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	findProduct := new(Product)
	if strings.HasSuffix(file.Name(), ".json") {
		readJSON(file, findProduct)
	} else if strings.HasSuffix(file.Name(), ".csv") {
		readCSV(file, findProduct)
	} else {
		log.Fatal("I can handle only \".json\" and \".csv\"")
	}

	fmt.Println(*findProduct)
}
