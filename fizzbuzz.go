package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Params struct {
	String1 string `json:"string1,omitempty"`
	String2 string `json:"string2,omitempty"`
	Int1    int    `json:"int1"`
	Int2    int    `json:"int2"`
	Limit   int    `json:"limit"`
}

//Fizzbuzz takes two strings and three integers and build the fizzbuzz based on those parameters
func Fizzbuzz(string1, string2 string, int1, int2, limit int) ([]string, error) {

	if limit < 1 {
		return nil, errors.New("limit must be > to 1")
	}

	if int1 < 1 || int2 < 1 {
		return nil, errors.New("multiples arguments must be > to 1")
	}

	chain := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			chain = append(chain, fmt.Sprintf("%s%s", string1, string2))
		} else if i%int1 == 0 {
			chain = append(chain, string1)
		} else if i%int2 == 0 {
			chain = append(chain, string2)
		} else {
			chain = append(chain, strconv.Itoa(i))
		}
	}

	return chain, nil
}
