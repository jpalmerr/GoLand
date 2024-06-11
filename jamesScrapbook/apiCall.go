package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// https://jsonmock.hackerrank.com/api/food_outlets?city=seattle

type RestaurantResponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	City          string     `json:"city"`
	Name          string     `json:"name"`
	EstimatedCost int        `json:"estimated_cost"`
	UserRating    UserRating `json:"user_rating"`
	ID            int        `json:"id"`
}

type UserRating struct {
	AverageRating float64 `json:"average_rating"`
	Votes         int     `json:"votes"`
}

func getRestaurantResponse(city string) (RestaurantResponse, error) {
	/**
	The defer keyword in Go is used to schedule a function call to be executed after the surrounding function returns.
	In this case, resp.Body.Close() is being deferred.

	*/
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s", city)
	// similar concept to Either data type -if req is successful err is nil. Otherwise err not nill and can deal with
	resp, err := http.Get(url)
	if err != nil {
		return RestaurantResponse{}, err
	}
	defer resp.Body.Close() // closing the resource

	/**
	Using defer to schedule the Close() call ensures that the response body is reliably closed,
	even in the presence of errors or early returns
	*/

	var data RestaurantResponse
	// resp.Body -> response body stream received from the server after making the HTTP request
	// json.NewDecoder -> creates a new json.Decoder instance
	// .Decode(data) -> Decode(&data) is a method on the json.Decoder instance that decodes the JSON & stores in variable
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return RestaurantResponse{}, err // return an empty struct (for safe usage later) and the error
	}

	return data, nil
}

func bestRestaurant(city string, cost int32) string {
	data, err := getRestaurantResponse(city)
	if err != nil {
		// Handle the error
		return "error!"
	}

	var bestRestaurant Data
	var maxRating float64
	maxRating = 0.0

	for _, restaurant := range data.Data {
		if restaurant.EstimatedCost <= int(cost) && restaurant.UserRating.AverageRating > maxRating {
			maxRating = restaurant.UserRating.AverageRating
			bestRestaurant = restaurant
		} else if restaurant.EstimatedCost <= int(cost) && restaurant.UserRating.AverageRating == maxRating && restaurant.EstimatedCost < bestRestaurant.EstimatedCost {
			bestRestaurant = restaurant
		}
	}

	if maxRating == 0.0 {
		return "no restaurant found" // No restaurant found
	}

	return bestRestaurant.Name
}

func main() {
	fmt.Printf(bestRestaurant("Seattle", 100))
}
