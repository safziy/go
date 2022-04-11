package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func round(val string, places int32) string {
	value, err := decimal.NewFromString(val)
	if err != nil {
		panic(fmt.Sprintf("can not convet string:%s to decimal", val))
	}
	return value.Round(places).String()
}

func ceil(val string) string {
	value, err := decimal.NewFromString(val)
	if err != nil {
		panic(fmt.Sprintf("can not convet string:%s to decimal", val))
	}
	return value.Ceil().String()
}

func floor(val string) string {
	value, err := decimal.NewFromString(val)
	if err != nil {
		panic(fmt.Sprintf("can not convet string:%s to decimal", val))
	}
	return value.Floor().String()
}

func main() {
	fmt.Println(round("3.24", 1))
	fmt.Println(round("3.25", 1))
	fmt.Println(round("3.249999", 1))
	fmt.Println(round("3.2455555", 1))
	fmt.Println(round("3.2500001", 1))
	fmt.Println(round("-3.24", 1))
	fmt.Println(round("-3.25", 1))
	fmt.Println(round("-3.249999", 1))
	fmt.Println(round("-3.2455555", 1))
	fmt.Println(round("-3.2511111", 1))

	fmt.Println(ceil("3.24"))
	fmt.Println(ceil("3.01"))
	fmt.Println(ceil("3.99"))
	fmt.Println(ceil("-3.24"))
	fmt.Println(ceil("-3.01"))
	fmt.Println(ceil("-3.99"))

	fmt.Println(floor("3.24"))
	fmt.Println(floor("3.01"))
	fmt.Println(floor("3.99"))
	fmt.Println(floor("-3.24"))
	fmt.Println(floor("-3.01"))
	fmt.Println(floor("-3.99"))
}
