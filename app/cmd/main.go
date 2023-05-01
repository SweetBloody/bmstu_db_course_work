package main

import (
	driverHandler "app/internal/pkg/driver/delivery/http"
	driverRepository "app/internal/pkg/driver/repository/postgresql"
	"fmt"
	"log"
	"net/http"

	driverUsecase "app/internal/pkg/driver/usecase"

	middleware "app/internal/app/middleware"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//type Driver struct {
//	DriverId        int    `db:"driver_id"`
//	DriverName      string `db:"driver_name"`
//	DriverCountry   string `db:"driver_country"`
//	DriverBirthDate string `db:"driver_birth_date"`
//}

//func RouteHandle(w http.ResponseWriter, r *http.Request) {
//	params := "user=postgresql dbname=formula1 password=postgresql host=localhost port=1000 sslmode=disable"
//	db, err := sqlx.Connect("postgresql", params)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	driver := models.Driver{}
//	err = db.Get(&driver, "select * from drivers where driver_id = 0")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Fprintf(w, "%#v", driver)
//}

func main() {
	params := "user=postgres dbname=formula1 password=postgres host=localhost port=1000 sslmode=disable"
	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driverRepo := driverRepository.NewPsqlDriverRepository(db)

	driverUsecase := driverUsecase.NewDriverUsecase(driverRepo)

	m := mux.NewRouter()

	driverHandler.NewDriverHandler(m, driverUsecase)

	mMiddleware := middleware.AccessLogMiddleware(m)

	fmt.Println("starting server at :5259")
	http.ListenAndServe(":5259", mMiddleware)
}
