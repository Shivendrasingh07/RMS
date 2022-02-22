package models

type Admin struct {
	Roleid   int    `db:"role_id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type Subadmin struct {
	Roleid   int    `db:"role_id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type User struct {
	Roleid      int     `db:"role_id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Email       string  `db:"email" json:"email"`
	Password    string  `db:"password" json:"password"`
	Address     string  `db:"address" json:"address"`
	Address_lat float64 `db:"address_lat" json:"address_lat"`
	Address_lnt float64 `db:"address_lng" json:"address_lnt"`
}

type Resturants struct {
	Creators_id     int    `db:"creators_id" json:"creators_id"`
	Resturants_name string `db:"restaurant_name" json:"restaurant_name"`

	Address     string  `db:"address" json:"address"`
	Address_lat float64 `db:"address_lat" json:"address_lat"`
	Address_lnt float64 `db:"address_lng" json:"address_lnt"`
}

type Rest struct {
	//Restname string `db:"restaurant_name" json:"restaurant_name"`
	Restname string ` json:"restaurant_name"`
}

type All_resturants struct {
	//	Creators_id     int    `db:"creators_id" json:"creators_id"`
	Resturants_name string `db:"restaurant_name" json:"restaurant_name"`

	Address     string  `db:"address" json:"address"`
	Address_lat float64 `db:"address_lat" json:"address_lat"`
	Address_lnt float64 `db:"address_lng" json:"address_lnt"`
}

type Dishes struct {
	Restaurant_id int    `db:"restaurant_id" json:"restaurant_id"`
	Creators_id   int    `db:"creators_id" json:"creators_id"`
	Dishes_name   string `db:"dishes_name" json:"dishes"`
	Price         int    `db:"price" json:"price"`
}

type Showdishes struct {
	Dishes_name string `db:"dishes_name" json:"dishes"`
	Price       int    `db:"price" json:"price"`
}

type Loginrole struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type FogetPassword struct {
	Email    string `db:"email" json:"email"`
	ID       int    `db:"id" json:"id"`
	Password string `db:"password" json:"password"`
}

type Cordinates struct {
	x float64
	y float64
}
type Cordinatesi struct {
	xi float64
	yi float64
}
