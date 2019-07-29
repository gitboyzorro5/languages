package main 

import  (
 
 "net/http"
 "log"
 "github.com/gorilla/mux"
 "encoding/json"
)


type Person struct{

	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`

	Adress *Adress `json:"address,omitempty"`
}

type Adress struct {

	City string `json:"city,omitempty"`
    State string `state:"city,omitempty"`

}   

var people []Person


func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){

	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){

params := mux.Vars(req)

for _, item := range people{

	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
}
json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){

params := mux.Vars(req)

var person Person

_ = json.NewDecoder(req.Body).Decode(&person)

person.ID = params["id"]

people = append(people,person)
json.NewEncoder(w).Encode(people)



}


func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){

params := mux.Vars(req){

	for index, item := range people {

		if  item.ID == params["id"] {
			people = append(people[:index], people[index+1]...)

			break
		}
	}
}
}

func main(){

	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Leonard", Lastname: "magut" , Adress: &Adress{City: "Nairobi", State: "Kenya"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "heather" })
	router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",GetPeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}",DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345",router))


}