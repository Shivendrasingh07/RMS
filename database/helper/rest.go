package helper

import (
	"database/sql"
	"github.com/RMS/database"
	"github.com/RMS/models"
)

func New_resturants(Resturants_name, Address string, Address_lat, Address_lnt float64, creators_id int) (int, error) {

	SQL := `SELECT name, role FROM users WHERE role_id=$1`
	var creatorsname, check_role string
	err := database.Data.QueryRowx(SQL, creators_id).Scan(&creatorsname, &check_role)
	if err != nil {
		return 404, err
	}

	if check_role == "admin" || check_role == "subadmin" {

		SQL = `INSERT INTO restaurants( restaurant_name,  address, address_lat, address_lng, creators_id)VALUES($1,$2,$3,$4,$5) RETURNING creators_id`
		var id int

		err = database.Data.Get(&id, SQL, Resturants_name, Address, Address_lat, Address_lnt, creators_id)
		if err != nil {
			return 404, err
		}

		return id, nil
	} else {
		return 404, nil
	}

}

func New_dish(Dishes_name string, price, rest_id, creators_id int) (int, error) {

	SQL := `SELECT  role FROM users WHERE role_id=$1`
	var Check_role string
	err := database.Data.QueryRowx(SQL, creators_id).Scan(&Check_role)
	if err != nil {
		return 404, err
	}

	if Check_role == "admin" || Check_role == "subadmin" {

		SQL = `INSERT INTO dishes(dishes_name, price, restaurant_id, creators_id)VALUES($1,$2,$3,$4) RETURNING id`
		var id int

		err = database.Data.Get(&id, SQL, Dishes_name, price, rest_id, creators_id)
		if err != nil {
			return 404, err
		}

		return id, nil
	} else {
		return 404, nil
	}

}

func Allrest(role_id int) ([]models.All_resturants, error) {

	SQL := `SELECT  restaurant_name, address,address_lat,address_lng from restaurants WHERE creators_id=$1 and archived_at isnull `
	rest := make([]models.All_resturants, 0)
	err := database.Data.Select(&rest, SQL, role_id)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return rest, nil
}

func Alldish(role_id int) ([]models.Showdishes, string, string, error) {

	SQL := `SELECT  restaurant_name FROM restaurants WHERE creators_id=$1`
	var resturant_name string
	err := database.Data.QueryRowx(SQL, role_id).Scan(&resturant_name)
	if err != nil {
		return nil, "", "", err
	}

	SQL = `SELECT  role FROM users WHERE role_id=$1`
	var Check_role string
	err = database.Data.QueryRowx(SQL, role_id).Scan(&Check_role)
	if err != nil {
		return nil, "", "", err
	}

	if Check_role == "admin" || Check_role == "subadmin" {

		SQL = `SELECT  dishes_name,price from dishes WHERE creators_id=$1 and archived_at isnull `
		rest := make([]models.Showdishes, 0)
		err = database.Data.Select(&rest, SQL, role_id)
		if err != nil {
			panic(err)
		}
		if err == sql.ErrNoRows {
			return nil, "", "", err
		}
		return rest, Check_role, resturant_name, nil
	} else {
		return nil, "", "", nil
	}
}
