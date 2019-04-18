package main

import (
	"encoding/json"
	"fmt"

	"database/sql"
	"log"
	"net/http"
	"os"

	"io/ioutil"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var err error
var db *sql.DB
var e = echo.New()

type Staff_struct struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json: "lastname"`
	Email     string `json : "email"`
	Position  string `json : "country"`
	Country   string `json : "age"`
}

type Staffs_struct struct {
	Staffs_struct []Staff_struct `json:"staffs_info"`
}

func main() {
	initDB()
	get_all_staffs()
	create_new_staff()
	update_staff_info()
	delete_staff_info()
	namee()

	defer db.Close()

	e.Start(":8000")

}

func initDB() {
	// router := mux.NewRouter()
	gotenv.Load()

	username := os.Getenv("username")
	password := os.Getenv("password")
	database := os.Getenv("database")
	hostname := os.Getenv("hostname")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", hostname, username, database, password) //Build connection string
	// fmt.Println(dbUri)
	db, err = sql.Open("postgres", dbUri)
	// fmt.Println(db)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
}

func get_all_staffs() {
	e.GET("/get_all_staffs", func(c echo.Context) error {
		sqlStatement := "SELECT * FROM employee_info"
		rows, err := db.Query(sqlStatement)

		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		}

		result := Staffs_struct{}

		for rows.Next() {
			staff_struct := Staff_struct{}
			err2 := rows.Scan(&staff_struct.Id, &staff_struct.Firstname, &staff_struct.Lastname, &staff_struct.Email, &staff_struct.Position, &staff_struct.Country)
			// Exit if we get an error
			if err2 != nil {
				return err2
			}
			result.Staffs_struct = append(result.Staffs_struct, staff_struct)
		}
		fmt.Println(result.Staffs_struct)
		return c.JSON(http.StatusCreated, result)

		//return c.String(http.StatusOK, "ok")
	})
}

func create_new_staff() {
	e.POST("/new_staff", func(c echo.Context) error {
		u := new(Staff_struct)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO employee_info (firstname, lastname, email, position, country) VALUES ($1, $2, $3, $4, $5)"
		res, err := db.Query(sqlStatement, u.Firstname, u.Lastname, u.Email, u.Position, u.Country)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})
}

func update_staff_info() {
	e.PATCH("/update_staff_info/:id", func(c echo.Context) error {
		u := new(Staff_struct)
		if err := c.Bind(u); err != nil {
			return err
		}

		sqlStatement := "UPDATE employee_info SET firstname=$1, lastname=$2, email=$3, position=$4, country=$5 WHERE id=$6"
		res, err := db.Query(sqlStatement, u.Firstname, u.Lastname, u.Email, u.Position, u.Country, u.Id)

		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}

		return c.String(http.StatusOK, string(u.Id))

	})
}

func delete_staff_info() {
	e.DELETE("/delete_staff_info/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM employee_info WHERE id = $1"
		res, err := db.Query(sqlStatement, id)
		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusOK, "Deleted")
		}
		return c.String(http.StatusOK, id+"Delete")
	})

}

func namee() {
	e.GET("/name", func(c echo.Context) error {
		s := Staff_struct{}

		sqlStatement := "INSERT INTO employee_info (firstname, lastname, email, position, country) VALUES ($1, $2, $3, $4, $5)"
		res, err := db.Query(sqlStatement, u.Firstname, u.Lastname, u.Email, u.Position, u.Country)

		defer c.Request().Body.Close()

		b, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, &s)
		if err != nil {
			panic(err)
		}
		return c.String(http.StatusOK, "ok")
	})

}

// var posttedCheckbook PostedCheckbook
// err = json.Unmarshal(body, &postedCheckbook)
// if err != nil {
// log.WithFields(logFields).Errorf("can't parse request body:%s", err.Error())

// w.WriteHeader(http.StatusBadRequest)
// fmt.Fprint(w, util.ParseError("parse_error", fmt.Sprintf("can't parse request body: %s", err.Error()), ""))

// return
// }

// body, err := ioutil.ReadAll(r.Body)
// if err != nil {
// log.WithFields(logFields).Errorf("can't read request body: %s", err.Error())

// w.WriteHeader(http.StatusInternalServerError)
// fmt.Fprint(w, util.ParseError("server_error", fmt.Sprintf("can't read request body: %s", err.Error()), ""))

// return
// }

// err = r.Body.Close()
// if err != nil {
// log.WithFields(logFields).Errorf("can't close request: %s", err.Error())

// w.WriteHeader(http.StatusInternalServerError)
// fmt.Fprint(w, util.ParseError("server_error", fmt.Sprintf("can't close request: %s", err.Error()), ""))

// return
