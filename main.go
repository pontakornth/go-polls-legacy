package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pontakornth/go-polls/handlers"
	"github.com/pontakornth/go-polls/repository"
	"github.com/pontakornth/go-polls/services"

	"github.com/spf13/viper"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "not found"})
}
func main() {
	initTimezone()
	initConfig()
	connection := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.database"))
	db, err := sqlx.Open("pgx", connection)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to the database")
	questionRepo, choiceRepo := repository.NewQuestionRepository(db), repository.NewChoiceRepository(db)
	pollsService := services.NewPollsService(questionRepo, choiceRepo)
	pollsHandler := handlers.NewPollsHandler(pollsService)
	http.HandleFunc("/", handle)
	http.HandleFunc("/polls", pollsHandler.GetAllQuestions)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("POLLS")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
