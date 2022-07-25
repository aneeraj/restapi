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
		Question2{Id: "2", Ques: "I have visible signs of nervousness such as sweaty palms, shaky hands, and so on right before a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I have butterflies in my stomach before a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I feel nauseated before a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I read through the test and feel that I do not know any of the answers", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I panic before and during a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I make mistakes on easy questions or put answers in the wrong places.", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "My mind goes blank during a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I remember the information that I blanked once I get out of the testing situation", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
		Question2{Id: "2", Ques: "I have trouble sleeping the night before a test", List1: "Not at all", List2: "Sometimes", List3: "More often", List4: "Always"},
	}
	Questions3 = []Question3{
		Question3{Id: "3", Ques: "I am worried about dirt, germs, and viruses. Ex. Fear of getting germs from touching door handles or shaking hands or sitting in certain chairs or seats or fear of getting AIDS.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I wash my hands very often or in a special way to be sure I am not dirty or contaminated.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I fear I will lose control and do something I dont want to do. Ex. Fear of driving into a tree, fear of running over someone, fear of stabbing someone.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "Repeated checking of door locks, the stove, the iron, or electrical outlets before leaving home; repeated checking that one’s cupboard at school is locked, or if one is properly dressed.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I am occupied with morality issues, justice, or what is right or wrong. Ex. Worries about always doing ”the right thing,” having told a lie, or having cheated someone.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I need to pray to cancel bad thoughts or feelings", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "How things are placed or how they are positioned is important to me. It needs to feel “just right” (but isnt associated with magical thinking).", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I have worries that I look peculiar; I am concerned that something is wrong with my looks.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
		Question3{Id: "3", Ques: "I do things that injure my body. Ex: Scratching and tearing the skin, cutting oneself or banging one’s head.", List1: "Not at all", List2: "A little", List3: "Moderately", List4: "Extremely"},
	}
	Questions4 = []Question4{
		Question4{Id: "4", Ques: "Do you have trouble falling asleep?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Do you have any medical conditions that disrupt your sleep?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Do you take anything to help you sleep?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Are you a shift worker or is your sleep schedule irregular?", List1: "Not at all", List2: "Occasionally", List3: "most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Are your legs restless and/or uncomfortable before bed?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Do you think something is wrong with your body?", List1: "Not at all", List2: "Several Days", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Do you feel nervous or worried?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Has anyone said that you stop breathing, gasp, snort, or choke in your sleep?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
		Question4{Id: "4", Ques: "Do you have difficulty staying awake during the day?", List1: "Not at all", List2: "Occasionally", List3: "Most nights/days", List4: "Always"},
	}
	Questions5 = []Question5{
		Question5{Id: "5", Ques: "Have any of your closest relationships been troubled  by a lot of arguments or repeated breakups? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you deliberately hurt yourself physically (e.g., punched yourself, cut yourself, burned yourself)? How about made a suicide attempt? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: " Have you been extremely moody? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you felt very angry a lot of the time? How about often acted in an angry or sarcastic manner? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you often been distrustful of other people? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you chronically felt empty?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you frequently felt unreal or as if things around you were unreal? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you often felt that you had no idea of who you are or that you have no identity? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question5{Id: "5", Ques: "Have you made desperate efforts to avoid feeling abandoned or being abandoned (e.g., repeatedly called someone to reassure yourself that he or she still cared, begged them not to leave you, clung to them physically)? ", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	Questions6 = []Question6{
		Question6{Id: "6", Ques: "You felt so good or so hyper that other people thought you were not your normal self or you were so hyper that you got into trouble?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You were so irritable that you shouted at people or started fights or arguments?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You felt much more self-confident than usual?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You got much less sleep than usual and found you didn’t really miss it?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You were so easily distracted by things around you that you had trouble concentrating or staying on track?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "Spending money got you or your family in trouble?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You were much more social or outgoing than usual, for example, you telephoned friends in the middle of the night?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You did things that were unusual for you or that other people might have thought were excessive, foolish, or risky?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
		Question6{Id: "6", Ques: "You were much more active or did many more things than usual?", List1: "Not at all", List2: "Several Days", List3: "More than half the days", List4: "Nearly every day"},
	}
	handleRequests()
}
