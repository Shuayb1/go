package main

import (
    "fmt"

    "database/sql"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/labstack/echo"
    _ "github.com/lib/pq"
    "github.com/subosito/gotenv"
    "github.com/labstack/echo/middleware"
    "github.com/dgrijalva/jwt-go"
)

var err error
var db *sql.DB
var e = echo.New()
var mySigningKey = []byte(os.Getenv("secrete_key"))
var username = os.Getenv("username")
var password = os.Getenv("password")
var database = os.Getenv("database")
var hostname = os.Getenv("hostname")

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
    e := echo.New()

    gr := e.Group("/token_required")
    
    gr.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningMethod: "HS256",
        SigningKey: []byte(os.Getenv("secrete_key")),
    }))

    initDB()
    e.GET("/get_token", get_token)
    gr.GET("/get_all_staffs",get_all_staffs)
    gr.POST("/new_staff", create_new_staff)
    gr.PATCH("/update_staff_info/:id", update_staff_info)
    gr.DELETE("/delete_staff_info/:id", delete_staff_info)

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

func get_all_staffs(c echo.Context) error {

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
}

func create_new_staff(c echo.Context) error {
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
}

func update_staff_info(c echo.Context) error {
    u := new(Staff_struct)

    id := c.Param("id")

    if err := c.Bind(u); err != nil {
        return err
    }

    sqlStatement := "UPDATE employee_info SET firstname=$1, lastname=$2, email=$3, position=$4, country=$5 WHERE id=$6"
    res, err := db.Query(sqlStatement, u.Firstname, u.Lastname, u.Email, u.Position, u.Country, id)

    if err != nil {
        fmt.Println(err)
        //return c.JSON(http.StatusCreated, u);
    } else {
        fmt.Println(res)
        return c.JSON(http.StatusCreated, "res")
    }

    return c.String(http.StatusOK, "res")
}

func delete_staff_info(c echo.Context) error {
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
}

func get_token(c echo.Context) error {
    // Create the token
    raw_token := jwt.New(jwt.SigningMethodHS256)

    // Create a map to store our claims
    claims := raw_token.Claims.(jwt.MapClaims)
    
    //  Set token claims
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // Sign the token with our secret
    token, err := raw_token.SignedString(mySigningKey)
    
    if err != nil{
        return err
    }
    return c.JSON(http.StatusOK, map[string]string{
        "token": token})
}
