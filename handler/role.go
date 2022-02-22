package handler

import (
	"encoding/json"
	"fmt"
	"github.com/RMS/database/helper"
	"github.com/RMS/middleware"

	"github.com/RMS/models"
	"net/http"
)

func AdminSignup(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	var role = "admin"
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		fmt.Println("err")
	}

	adminid, newerr := helper.Newadmin(admin.Roleid, admin.Name, admin.Email, admin.Password, role)

	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(adminid)
	if err2 != nil {
		fmt.Println("err")
	}
}

func SubadminSignup(w http.ResponseWriter, r *http.Request) {
	var subadmin models.Subadmin
	var role = "subadmin"
	err := json.NewDecoder(r.Body).Decode(&subadmin)
	if err != nil {
		fmt.Println("err")
	}

	creators_id := middleware.GetUserFromContext(r)
	subadminid, newerr := helper.Newsubadmin(subadmin.Roleid, subadmin.Name, subadmin.Email, subadmin.Password, creators_id, role)

	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(subadminid)
	if err2 != nil {
		fmt.Println("err")
	}
}

func UserSignup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var role = "user"
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("err")
	}

	creators_id := middleware.GetUserFromContext(r)
	user_id, newerr := helper.Newuser(user.Roleid, user.Name, user.Email, user.Password, creators_id, role)
	fmt.Println(user_id)
	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(user_id)
	if err2 != nil {
		fmt.Println("err")
	}

	err = helper.Address(user_id, user.Address, user.Address_lat, user.Address_lnt)
	if err != nil {
		panic(err)
	}
}

/*
func Subsadmin_Signup(w http.ResponseWriter, r *http.Request) {
	var user models.Subadmin
	var role = "subadmin"

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("err")
	}
	creators_id := middleware.GetUserFromContext(r)
	subadminid, newerr := helper.Newsubadmin(user.Name, user.Email, user.Password, role, creators_id)

	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(subadminid)
	if err2 != nil {
		fmt.Println("err")
	}
}
*/

func Login(writer http.ResponseWriter, request *http.Request) {
	var Cred models.Loginrole
	err := json.NewDecoder(request.Body).Decode(&Cred)
	if err != nil {
		fmt.Println("err")
	}
	loginUser, loginErr := helper.Login(Cred.Email, Cred.Password)
	if loginErr != nil {
		fmt.Println("err")
	}

	err2 := json.NewEncoder(writer).Encode(loginUser)
	if err2 != nil {
		fmt.Println("err")
	}

}

func ResetPassword(w http.ResponseWriter, request *http.Request) {
	var cred models.FogetPassword
	err := json.NewDecoder(request.Body).Decode(&cred)
	if err != nil {
		fmt.Println(err)
	}
	NewErr := helper.ForgetPass(cred.ID, cred.Email, cred.Password)
	if NewErr != nil {
		fmt.Println(NewErr)
	}
}

func Allusers(w http.ResponseWriter, r *http.Request) {

	userid := middleware.GetUserFromContext(r)
	all_users, err := helper.Showuser(userid)
	if err != nil {
		fmt.Println("helper all user error")
		panic(err)
		return
	}
	err = json.NewEncoder(w).Encode(all_users)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

}

func Allsubadmins(w http.ResponseWriter, r *http.Request) {

	userid := middleware.GetUserFromContext(r)

	all_subadmin, err := helper.Showsubadmin(userid)
	if err != nil {
		fmt.Println("helper subadmin error ")
		panic(err)
	}
	err = json.NewEncoder(w).Encode(all_subadmin)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

}

func Distance(w http.ResponseWriter, r *http.Request) {

	var restname models.Rest

	err := json.NewDecoder(r.Body).Decode(&restname)
	if err != nil {
		fmt.Println("err")
	}

	userid := middleware.GetUserFromContext(r)

	dist, err := helper.Calculate(userid, restname.Restname)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(dist)
	if err != nil {
		panic(err)
	}
}
