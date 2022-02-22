package handler

import (
	"encoding/json"
	"fmt"
	"github.com/RMS/database/helper"
	"github.com/RMS/middleware"
	"github.com/RMS/models"
	"net/http"
)

func CreateResturants(w http.ResponseWriter, r *http.Request) {

	var rest models.Resturants
	//var role = "user"
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		fmt.Println("err")
	}

	creators_id := middleware.GetUserFromContext(r)
	user_id, newerr := helper.New_resturants(rest.Resturants_name, rest.Address, rest.Address_lat, rest.Address_lnt, creators_id)
	fmt.Println(user_id)
	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(user_id)
	if err2 != nil {
		fmt.Println("err")
	}

}

func CreateDishes(w http.ResponseWriter, r *http.Request) {

	var dish models.Dishes
	//var role = "user"
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		fmt.Println("err")
	}

	creators_id := middleware.GetUserFromContext(r)
	dish_id, newerr := helper.New_dish(dish.Dishes_name, dish.Price, dish.Restaurant_id, creators_id)
	fmt.Println(dish_id)
	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(dish_id)
	if err2 != nil {
		fmt.Println("err")
	}

}

func Allrestaurants(w http.ResponseWriter, r *http.Request) {

	role_id := middleware.GetUserFromContext(r)

	allrest, err := helper.Allrest(role_id)
	if err != nil {
		fmt.Println("helper allrest error ")
		panic(err)
	}
	err = json.NewEncoder(w).Encode(allrest)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

}

func Alldishes(w http.ResponseWriter, r *http.Request) {

	role_id := middleware.GetUserFromContext(r)

	alldish, role, rest_name, err := helper.Alldish(role_id)
	if err != nil {
		fmt.Println("helper allrest error ")
		panic(err)
	}
	//w.Write([]byte(fmt.Sprintf("Created by ", role)))
	//	w.Write([]byte(fmt.Sprintf("Resturant", rest_name)))

	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

	err = json.NewEncoder(w).Encode(rest_name)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}
	err = json.NewEncoder(w).Encode(alldish)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

}
