package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/subosito/gotenv"
)

var db *sql.DB //database
var err error
var (
	id        int
	firstname string
	lastname  string
	email     string
	position  string
	country   string
)

type Staff_summary struct {
	Id        int    `json:"id"`
	Firstname string `json:"info"`
	LastName  string `json:"info"`
	Email     string `json:"info"`
	Position  string `json:"info"`
	Location  string `json:"info"`
	Country   string `json:"info"`
}

// type Field map[string]Staff_summary

// 	LastName  string
// 	Email     string
// 	Position  string
// 	Location  string
// 	Country   string
// 	Code      string
// 	Capital   string
// Tags pq.StringArray 'gorm:"type:varchar(64)[]"'
// }

// type staffs struct {
// 	Staffs []staff_summary
// }

func main() {
	var router = mux.NewRouter()
	router.Use(commonMiddleware)

	initDB()

	defer db.Close()

	// http.HandleFunc("/api/shuayb/get_all", indexHandler)
	// router.HandleFunc("/api/shuayb/get_all", Get_all).Methods("GET")
	router.HandleFunc("/api/shuayb/get_by_country", Get_by_country).Methods("GET")
	router.HandleFunc("/api/shuayb/new_staff", Create_new_staff).Methods("POST")
	router.HandleFunc("/api/shuayb/update_info/{id}", Update_staff_info).Methods("PATCH")
	router.HandleFunc("/api/shuayb/delete_staff/{id}", Delete_staff).Methods("DELETE")
	// http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
		panic(err)
	}
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	st := staffs{}
// 	err := Get_all(&st)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	out, err := json.Marshal(st)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	fmt.Fprintf(w, string(out))
// }

// func Get_all(w http.ResponseWriter, r *http.Request) {
// 	c, err := db.Query("SELECT * FROM employee_info")
// 	// fmt.Println(c)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer c.Close()

// 	// fds := json.Marshal(c)
// 	masterData := make(map[string][]interface{})
// 	// staffs := []*Staff_summary{}
// 	for i := c.Next(); true; {
// 		// var id int
// 		// var info string

// 		// err = c.Scan(&id, &info)
// 		err = c.Scan(&id, &firstname, &lastname, &email, &position, &country)
// 		if err != nil {
// 			log.Fatal(err)

// 		} else {
// 			masterData = append(masterData, c[i])
// 		}
// 		return masterData
// 		// log.Println(c)
// 		// fmt.Fprintf(w, "%d", id)
// 		// fmt.Fprintf(w, staffs)

// 	}
// 	err = c.Err()

// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(masterData)
// 	w.Header().Add("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(masterData)
// }

func Get_by_country(w http.ResponseWriter, r *http.Request) {

	// func (d *DbDao) makeStructJSON(queryText string, w http.ResponseWriter) error {

	// returns rows *sql.Rows
	rows, err := db.Query("SELECT * FROM employee_info")
	if err != nil {
		panic(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	masterData := make(map[string][]interface{})

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		for i, v := range values {

			x := v.([]byte)

			//NOTE: FROM THE GO BLOG: JSON and GO - 25 Jan 2011:
			// The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays; it will happily unmarshal any valid JSON blob into a plain interface{} value. The default concrete Go types are:
			//
			// bool for JSON booleans,
			// float64 for JSON numbers,
			// string for JSON strings, and
			// nil for JSON null.

			// if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
			// 	masterData[columns[i]] = append(masterData[columns[i]], nx)
			if b, ok := strconv.ParseBool(string(x)); ok == nil {
				masterData[columns[i]] = append(masterData[columns[i]], b)
			} else if "string" == fmt.Sprintf("%T", string(x)) {
				masterData[columns[i]] = append(masterData[columns[i]], string(x))
			} else {
				fmt.Printf("Failed on if for type %T of %v\n", x, x)
			}

		}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(masterData)

	if err != nil {
		panic(err)
	}
	panic(err)

}

func Create_new_staff(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Ho")
}

func Update_staff_info(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hu")
}

func Delete_staff(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hy")
}
