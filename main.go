package main

import (
	"back-end/modules/core/common"
	"back-end/modules/scheduler"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	dateFormat = "2006_01_02"
)

func init() {
	var err error
	err = godotenv.Load() // can pass the location. default root
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetFlags(log.Ldate | log.Ltime)
	file, err := os.OpenFile("logs/"+time.Now().Format(dateFormat)+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		log.SetOutput(file)
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", common.Login).Methods("POST")
	router.HandleFunc("/auth/forgot", common.Forgot).Methods("POST")

	router.HandleFunc("/api/save/mother", common.SaveMother).Methods("POST")
	router.HandleFunc("/api/save/father", common.SaveFather).Methods("POST")
	router.HandleFunc("/api/save/child", common.SaveChild).Methods("POST")
	router.HandleFunc("/api/save/weight", common.SaveWeight).Methods("POST")
	router.HandleFunc("/api/save/location", common.SaveLocation).Methods("POST")

	router.HandleFunc("/api/get/weight", common.GetWeight).Methods("POST")
	router.HandleFunc("/api/get/children", common.GetChildren).Methods("POST")
	router.HandleFunc("/api/get/mother", common.GetMother).Methods("POST")
	router.HandleFunc("/api/get/child", common.GetChild).Methods("POST")
	router.HandleFunc("/api/get/location", common.GetLocation).Methods("POST")
	router.HandleFunc("/api/get/all/mothers", common.GetAllMothers).Methods("POST")
	router.HandleFunc("/api/get/all/children", common.GetAllChildren).Methods("POST")

	router.HandleFunc("/api/check/nic", common.CheckNic).Methods("POST")

	scheduler.StartScheduler()

	fmt.Println("Starting Server at - http://0.0.0.0:3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", router))
}