package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tax struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	PaymentDue  int
}

func main() {
	taxes := []Tax{}

	data, err := os.ReadFile(`\\qumulo\BLIS\ScottGetzFiles\Mi-Tek\mitekOne.json`)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &taxes)
	if err != nil {
		fmt.Println(err)
		return
	}

	taxData, err := grabPaymentInfo("pdaberk@blogs.com", taxes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", taxData)

	// var counter int

	// for {
	// 	counter = counter + 1
	// 	if counter == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Hello World")
	// 	time.Sleep(3 * time.Second)
	// }

	// for i := 0; i < len(t); i++ {
	// 	t[i].PhoneNumber = "1111"
	// 	fmt.Println(t[i].PhoneNumber)
	// }

	// for i, v := range t {
	// 	t[i].PhoneNumber = "1111"
	// 	fmt.Println(v.PhoneNumber)
	// }

	// for _, v := range t {
	// 	fmt.Println(v.PhoneNumber)
	// }
}

func grabPaymentInfo(s string, t []Tax) (Tax, error) {
	var taxInfo Tax
	var found bool

	for _, v := range t {
		if v.Email == s {
			taxInfo = v
			found = true
		}
	}

	if !found {
		return taxInfo, fmt.Errorf("%s Email Not Found", s)
	}

	return taxInfo, nil
}
