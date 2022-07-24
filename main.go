package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
    "github.com/gorilla/mux"
)

type Questions struct {
    Id      string    `json:"Id"`
    Ques   string `json:"Question"`
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

    for _, Question := range Questions {
        if Question.Id == key {
            json.NewEncoder(w).Encode(Question)
        }
    }
}


func createNewQuestion(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Question struct
    // append this to our Questions array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var Question Question 
    json.Unmarshal(reqBody, &Question)
    // update our global Questions array to include
    // our new Question
    Questions = append(Questions, Question)

    json.NewEncoder(w).Encode(Question)
}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, Question := range Questions {
        if Question.Id == id {
            Questions = append(Questions[:index], Questions[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/Questions", returnAllQuestions)
    myRouter.HandleFunc("/Question", createNewQuestion).Methods("POST")
    myRouter.HandleFunc("/Question/{id}", deleteQuestion).Methods("DELETE")
    myRouter.HandleFunc("/Question/{id}", returnSingleQuestion)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Questions = []Question{
        Question{Id: "1", Title: "Hello", Desc: "Question Description", Content: "Question Content"},
        Question{Id: "2", Title: "Hello 2", Desc: "Question Description", Content: "Question Content"},
    }
    handleRequests()
}