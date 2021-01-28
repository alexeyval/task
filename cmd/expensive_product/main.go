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
	"task/config"
)

type Product struct {
	Name   string `json:"product"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

type ExpensiveProduct struct {
	Product
}

func (expProduct *ExpensiveProduct) FindExpensiveProduct(newProducts []Product) {
	for _, newProduct := range newProducts {
		if newProduct.Price > expProduct.Price ||
			newProduct.Price == expProduct.Price &&
				newProduct.Rating > expProduct.Rating {
			expProduct.Product = newProduct
		}
	}
}

func readCSV(file *os.File, expProduct *ExpensiveProduct) {
	parser := csv.NewReader(file)
	if _, err := parser.Read(); err != nil {
		log.Fatal(err)
	}

	newProducts := make([]Product, 0, config.BufferProducts)
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

		newProducts = append(newProducts, Product{Name, Price, Rating})
		if len(newProducts) == config.BufferProducts {
			expProduct.FindExpensiveProduct(newProducts)
			newProducts = newProducts[:0]
		}
	}
	expProduct.FindExpensiveProduct(newProducts)
	newProducts = newProducts[:0]
}

func readJSON(file *os.File, expProduct *ExpensiveProduct) {
	dec := json.NewDecoder(file)

	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	newProducts := make([]Product, 0, config.BufferProducts)
	for dec.More() {
		var newProduct Product
		err := dec.Decode(&newProduct)
		if err != nil {
			log.Fatal(err)
		}

		newProducts = append(newProducts, newProduct)
		if len(newProducts) == config.BufferProducts {
			expProduct.FindExpensiveProduct(newProducts)
			newProducts = newProducts[:0]
		}
	}
	expProduct.FindExpensiveProduct(newProducts)
	newProducts = newProducts[:0]

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

	expProduct := ExpensiveProduct{}
	if strings.HasSuffix(file.Name(), ".json") {
		readJSON(file, &expProduct)
	} else if strings.HasSuffix(file.Name(), ".csv") {
		readCSV(file, &expProduct)
	} else {
		log.Fatal("I can handle only \".json\" and \".csv\"")
	}

	fmt.Println(expProduct.Product)
}
