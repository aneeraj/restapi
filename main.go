package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
    "github.com/gorilla/mux"
)

type Question struct {
    Id      string    `json:"Id"`
    Ques   string `json:"ques"`
    List1 string `json:"list1"`
    List2 string `json:"list2"`
}

var Questions []Question

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello API")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllQuestions(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllQuestions")
    json.NewEncoder(w).Encode(Questions)
}

func returnSingleQuestion(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, question := range Questions {
        if question.Id == key {
            json.NewEncoder(w).Encode(question)
        }
    }
}


func createNewQuestion(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Question struct
    // append this to our Questions array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var question Question 
    json.Unmarshal(reqBody, &question)
    // update our global Questions array to include
    // our new Question
    Questions = append(Questions, question)

    json.NewEncoder(w).Encode(question)
}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, question := range Questions {
        if question.Id == id {
            Questions = append(Questions[:index], Questions[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/questions", returnAllQuestions)
    myRouter.HandleFunc("/question", createNewQuestion).Methods("POST")
    myRouter.HandleFunc("/question/{id}", deleteQuestion).Methods("DELETE")
    myRouter.HandleFunc("/question/{id}", returnSingleQuestion)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    Questions = []Question{
        Question{Id: "1", Ques: "Hello", List1: "Question Description", List2: "Question Content"},
        Question{Id: "2", Ques: "Hello 2", List1: "Question Description", List2: "Question Content"},
    }
    handleRequests()
}