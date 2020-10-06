package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/test", testDatabaseConnection)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello database connection testing World")
}

//dataSourceName is the post data key for the connection string
func testDatabaseConnection(w http.ResponseWriter, r *http.Request) {

	// Double check it's a post request being made
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Print("You must use POST to send the data.")
		fmt.Fprintf(w, "invalid_http_method You must use POST to send the data.")
		return
	}

	//get the data in the POST
	r.ParseForm()
	dataSourceName := r.Form.Get("dataSourceName")

	//Start recording time
	start := time.Now()

	//create connection to DB
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Failed to open connection to Database.", err)
		return
	}

	//record how long the connection took in microseconds
	elapsed := time.Since(start)
	elaspedInSeconds := float64(elapsed) / float64(time.Second)

	//display
	fmt.Printf("Successfully opened connection to the database in %f seconds.\n", elaspedInSeconds)

	//Pinging the DB
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB functionality appears in tact. Now closing connection...")

	db.Close()

}
