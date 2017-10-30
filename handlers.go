package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func queryParams(values url.Values) (*Params, error) {
	params := &Params{}

	params.String1 = values.Get("String1")
	params.String2 = values.Get("String2")

	int1, err := strconv.Atoi(values.Get("Int1"))
	if err != nil {
		return nil, err
	}
	params.Int1 = int1

	int2, err := strconv.Atoi(values.Get("Int2"))
	if err != nil {
		return nil, err
	}
	params.Int2 = int2

	limit, err := strconv.Atoi(values.Get("Limit"))
	if err != nil {
		return nil, err
	}
	params.Limit = limit

	return params, nil
}

//FizzbuzzHandler retrives data from the http GET request,
//build the json answer and send it back
func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	params, err := queryParams(values)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := Fizzbuzz(params.String1, params.String2, params.Int1, params.Int2, params.Limit)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

//ExampleHandler is a mockup endpoint thats returns a fizzbuzz with default values
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	res, err := Fizzbuzz("fizz", "buzz", 3, 5, 100)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
