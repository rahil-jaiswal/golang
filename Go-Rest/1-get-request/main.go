package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// type foodReview struct {
// 	Page       int `json:"page"`
// 	PerPage    int `json:"per_page"`
// 	Total      int `json:"total"`
// 	TotalPages int `json:"total_pages"`
// 	Data       []struct {
// 		City          string `json:"city"`
// 		Name          string `json:"name"`
// 		EstimatedCost int    `json:"estimated_cost"`
// 		UserRating    struct {
// 			AverageRating float64 `json:"average_rating"`
// 			Votes         int     `json:"votes"`
// 		} `json:"user_rating"`
// 		ID int `json:"id"`
// 	} `json:"data"`
// }

type Hotel struct {
	Data []struct {
		UserRating struct {
			Votes int `json:"votes"`
		} `json:"user_rating"`
	} `json:"data"`
}

func main() {
	url := "https://jsonmock.hackerrank.com/api/food_outlets?city=Seattle&estimated_cost=110"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//fmt.Println(response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))
	//x := foodReview{}

	//x := Hotel{}
	x := gjson.Get(string(content), "data.0.user_rating.votes")
	// err = json.Unmarshal(content, &x)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(x.String())

}
