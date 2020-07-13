package main

import (
	"fmt"
	"goApiRest/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbCnn := database.InitDB()

	defer dbCnn.Close()
	fmt.Println(dbCnn)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
