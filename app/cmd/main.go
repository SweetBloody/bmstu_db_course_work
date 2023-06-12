package main

import (
	"fmt"
	"log"
	"net/http"

	authHandler "app/internal/pkg/auth/delivery/http"
	driverHandler "app/internal/pkg/driver/delivery/http"
	grandPrixHandler "app/internal/pkg/grand_prix/delivery/http"
	qualHandler "app/internal/pkg/qual_result/delivery/http"
	raceHandler "app/internal/pkg/race_result/delivery/http"
	teamHandler "app/internal/pkg/team/delivery/http"
	trackHandler "app/internal/pkg/track/delivery/http"
	userHandler "app/internal/pkg/user/delivery/http"

	driverRepository "app/internal/pkg/driver/repository/postgresql"
	grandPrixRepository "app/internal/pkg/grand_prix/repository/postgresql"
	qualRepository "app/internal/pkg/qual_result/repository/postgresql"
	raceRepository "app/internal/pkg/race_result/repository/postgresql"
	teamRepository "app/internal/pkg/team/repository/postgresql"
	trackRepository "app/internal/pkg/track/repository/postgresql"
	userRepository "app/internal/pkg/user/repository/postgresql"

	driverUsecase "app/internal/pkg/driver/usecase"
	grandPrixUsecase "app/internal/pkg/grand_prix/usecase"
	qualUsecase "app/internal/pkg/qual_result/usecase"
	raceUsecase "app/internal/pkg/race_result/usecase"
	teamUsecase "app/internal/pkg/team/usecase"
	trackUsecase "app/internal/pkg/track/usecase"
	userUsecase "app/internal/pkg/user/usecase"

	middleware "app/internal/app/middleware"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	params := "user=postgres dbname=formula1 password=postgres host=localhost port=1000 sslmode=disable"
	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driverRepo := driverRepository.NewPsqlDriverRepository(db)
	teamRepo := teamRepository.NewPsqlTeamRepository(db)
	trackRepo := trackRepository.NewPsqlTrackRepository(db)
	gpRepo := grandPrixRepository.NewPsqlGPRepository(db)
	raceRepo := raceRepository.NewPsqlRaceResultRepository(db)
	qualRepo := qualRepository.NewPsqlQualResultRepository(db)
	userRepo := userRepository.NewPsqlUserRepository(db)

	driverUcase := driverUsecase.NewDriverUsecase(driverRepo)
	teamUcase := teamUsecase.NewTeamUsecase(teamRepo)
	trackUcase := trackUsecase.NewTrackUsecase(trackRepo)
	gpUcase := grandPrixUsecase.NewGrandPrixUsecase(gpRepo)
	raceUcase := raceUsecase.NewRaceResultUsecase(raceRepo)
	qualUcase := qualUsecase.NewQualResultUsecase(qualRepo)
	userUcase := userUsecase.NewUserUsecase(userRepo)

	m := mux.NewRouter()

	driverHandler.NewDriverHandler(m, driverUcase)
	teamHandler.NewTeamHandler(m, teamUcase)
	trackHandler.NewTrackHandler(m, trackUcase)
	grandPrixHandler.NewDriverHandler(m, gpUcase, raceUcase, qualUcase)
	raceHandler.NewRaceResultHandler(m, raceUcase)
	qualHandler.NewQualResultHandler(m, qualUcase)
	authHandler.NewAuthHandler(m, userUcase)
	userHandler.NewUserHandler(m, userUcase)

	mMiddleware := middleware.LogMiddleware(m)

	fmt.Println("starting server at :5259")
	http.ListenAndServe(":5259", mMiddleware)
}
