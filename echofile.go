package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Ram struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Bird struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func getcat(c echo.Context) error {
	catname := c.QueryParam("name")
	cattype := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("This is a %s of %s", catname, cattype))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catname,
			"type": cattype,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": " do well",
	})

	// return c.String(http.StatusOK, fmt.Sprintf("This is a %s of %s", catname, cattype))
}

func addcat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed reading %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Thisisis")
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is d cat %v", cat)
	return c.String(http.StatusOK, "THANKS")
}

func addram(c echo.Context) error {
	ram := Ram{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&ram)
	if err != nil {

		log.Printf("ddsdssd", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf(" blablabla")
	return c.String(http.StatusOK, "THANKS")
}

func addbird(c echo.Context) error {
	bird := Bird{}

	err := c.Bind(&bird)

	if err != nil {

		log.Printf("ddsdssd", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf(" blablabla")
	return c.String(http.StatusOK, "THANKS")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "WE are the admin")

}

func main() {
	e := echo.New()

	//g := e.Group("/admin", middleware.Logger()) //u can add middleware like this or

	//to group endpoints
	g := e.Group("/admin")

	//this logs the server interaction   //can do this to add middleware too
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	//basic authetication for admin enpoint
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool,error) {
		//check in d db

		if username == "ubanquity" && password == "ubanquity" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/main", mainAdmin)

	// e.GET("/cat/:data", getcat, middleware.Logger()) // u can add middleware like this too
	g.GET("/cat/:data", getcat)

	e.POST("/cats", addcat)
	e.POST("/rams", addram)
	e.POST("/birds", addbird)

	e.Start(":8000")
}

// defer c.Request().Body.Close()
// b, err := ioutil.ReadAll(c.Request().Body)

// defer c.Request().Body.Close()
// err := json.NewDecoder(c.Request().Body).Decode(&ram

// err := c.Bind(&bird)
