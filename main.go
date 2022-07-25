package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Question1 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions1 []Question1

type Question2 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions2 []Question2

type Question3 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions3 []Question3

type Question4 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions4 []Question4

type Question5 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions5 []Question5

type Question6 struct {
	Id    string `json:"Id"`
	Ques  string `json:"ques"`
	List1 string `json:"list1"`
	List2 string `json:"list2"`
	List3 string `json:"list3"`
	List4 string `json:"list4"`
}

var Questions6 []Question6

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello API")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllQuestions1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions1)
}

func returnAllQuestions2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions2)
}

func returnAllQuestions3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions3)
}

func returnAllQuestions4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions4)
}

func returnAllQuestions5(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions5)
}

func returnAllQuestions6(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions6)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/questions1", returnAllQuestions1)
	myRouter.HandleFunc("/questions2", returnAllQuestions2)
	myRouter.HandleFunc("/questions3", returnAllQuestions3)
	myRouter.HandleFunc("/questions4", returnAllQuestions4)
	myRouter.HandleFunc("/questions5", returnAllQuestions5)
	myRouter.HandleFunc("/questions6", returnAllQuestions6)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Questions1 = []Question1{
		Question1{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question1{Id: "1", Ques: "Blah2", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question1{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question1{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question1{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question1{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions2 = []Question2{
		Question2{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question2{Id: "1", Ques: "Blah2", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question2{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question2{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question2{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question2{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions3 = []Question3{
		Question3{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question3{Id: "1", Ques: "Blah2", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question3{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question3{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question3{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question3{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions4 = []Question4{
		Question4{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question4{Id: "1", Ques: "Blah2", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question4{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question4{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question4{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question4{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions5 = []Question5{
		Question5{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "1", Ques: "Blah2", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions6 = []Question6{
		Question6{Id: "1", Ques: "Little interest or pleasure in doing things", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "1", Ques: "Hey Neeraj", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "1", Ques: "Blah3", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "1", Ques: "Blah4", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "1", Ques: "Blah5", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "1", Ques: "Blah6", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	handleRequests()
}
