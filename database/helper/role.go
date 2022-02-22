package helper

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/RMS/database"
	"math"

	"github.com/RMS/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//NewUser save the data of new user

func Newadmin(roleid int, name, email, password, role string) (int, error) {
	SQL := `INSERT INTO users(role_id,name,email,password,role)VALUES($1,$2,$3,$4,$5) RETURNING role_id`
	var id int
	pass := HashPassword(password)
	err := database.Data.Get(&id, SQL, roleid, name, email, pass, role)
	if err != nil {
		return 404, err
	}
	return id, nil
}

func Newsubadmin(roleid int, name, email, password string, creators_id int, role string) (int, error) {

	SQL := `SELECT name, role FROM users WHERE role_id=$1`
	var creatorsname, check_role string
	err := database.Data.QueryRowx(SQL, creators_id).Scan(&creatorsname, &check_role)
	if err != nil {
		return 404, err
	}

	if check_role == "admin" {

		SQL = `INSERT INTO users(role_id,name,email,password,role,created_by)VALUES($1,$2,$3,$4,$5,$6) RETURNING role_id`
		var id int
		pass := HashPassword(password)
		err = database.Data.Get(&id, SQL, roleid, name, email, pass, role, creatorsname)
		if err != nil {
			return 404, err
		}
		return id, nil

	} else {
		return 404, nil
	}

}

func Newuser(roleid int, name, email, password string, creators_id int, role string) (int, error) {

	SQL := `SELECT name, role FROM users WHERE role_id=$1`
	var creatorsname, check_role string
	err := database.Data.QueryRowx(SQL, creators_id).Scan(&creatorsname, &check_role)
	if err != nil {
		return 404, err
	}

	if check_role == "admin" || check_role == "subadmin" {

		SQL = `INSERT INTO users(role_id,name,email,password,role,created_by)VALUES($1,$2,$3,$4,$5,$6) RETURNING role_id`
		var id int
		pass := HashPassword(password)
		err = database.Data.Get(&id, SQL, roleid, name, email, pass, role, creatorsname)
		if err != nil {
			return 404, err
		}

		return id, nil
	} else {
		return 404, nil
	}

}

func Address(user_id int, add string, add1, add2 float64) error {
	var id int
	SQL := `INSERT INTO addresses(role_id,address,address_lat,address_lnt)VALUES($1,$2,$3,$4) RETURNING role_id`
	err := database.Data.Get(&id, SQL, user_id, add, add1, add2)
	if err != nil {
		return err
	}

	return nil
}

//HashPassword helps to encrypt the password
func HashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(pass)
}

//CheckHashPassword  verify the password of user
func CheckHashPassword(password, hashpass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

//Login function takes email and password, returns userID
func Login(email, password string) (string, error) {
	SQL := `SELECT role_id,password FROM users WHERE email=$1`
	var id, hashpass string
	err := database.Data.QueryRowx(SQL, email).Scan(&id, &hashpass)
	if err != nil {
		return "err", err
	}
	pass, passErr := CheckHashPassword(password, hashpass)
	if pass != true && passErr != nil {
		return "err", passErr
	}
	//fmt.Println(userid)

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("token error : +v", err)
	}

	return ss, nil

}

func ForgetPass(userid int, email, password string) error {
	SQL := `SELECT id FROM users WHERE email=$1`
	var id int
	err := database.Data.QueryRowx(SQL, email).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	if userid == id {
		newSQL := `UPDATE users SET password=$1 where id=$2`
		pass := HashPassword(password)
		_, newErr := database.Data.Exec(newSQL, pass, userid)
		if newErr != nil {
			fmt.Println(newErr)
		}
	} else {
		return errors.New("WRONG CREDENTIALS")
	}

	return nil
}

func Showuser(userid int) ([]models.User, error) {

	SQL := `SELECT id,name,email from users  `
	user := make([]models.User, 0)
	err := database.Data.Select(&user, SQL)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return user, nil
}

func Showsubadmin(userid int) ([]models.Admin, error) {

	SQL := `SELECT  role FROM users WHERE role_id=$1`
	var check_role string
	err := database.Data.QueryRowx(SQL, userid).Scan(&check_role)
	if err != nil {
		return nil, err
	}

	if check_role == "admin" {

		SQL := `SELECT id,name,email from users`
		user := make([]models.Admin, 0)
		err := database.Data.Select(&user, SQL)
		if err != nil {
			panic(err)
		}
		if err == sql.ErrNoRows {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func Calculate(userid int, restname string) (float64, error) {

	SQL := `SELECT address_lat,address_lnt FROM addresses WHERE role_id=$1`
	var x, y float64
	err := database.Data.QueryRowx(SQL, userid).Scan(&x, &y)
	if err != nil {
		return 404, err
	}

	SQL = `SELECT address_lat,address_lng FROM restaurants WHERE restaurant_name=$1`
	var xi, yi float64
	err = database.Data.QueryRowx(SQL, restname).Scan(&xi, &yi)
	if err != nil {
		return 404, err
	}

	ans := Distance(x, y, xi, yi)

	return ans, err
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
