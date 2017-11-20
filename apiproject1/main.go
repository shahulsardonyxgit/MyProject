package main

//importing mux for HTTP request,json,log
import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

// Create a structure called person with 4 attributes
type Person struct{
  ID string `json:"ID,omitempty"`
  Firstname string `json:"firstName,omitempty"`
  Lastname string `json:"lastName,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct{
  City string `json:"city,omitempty"`
  State string `json:state,omitempty`
}

//global variable of type Person
var people []Person

//function to get all the person details
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
 json.NewEncoder(w).Encode(people)
}

//function to find a particular person with id
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
    params :=mux.Vars(req)
    for _, item:=range people{
      if item.ID==params["id"]{
        json.NewEncoder(w).Encode(item)
        return
      }
    }
 json.NewEncoder(w).Encode(&Person{})
}

//function to create a new person
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
    params:=mux.Vars(req)
    var person Person
    _=json.NewDecoder(req.Body).Decode(&person)
    person.ID=params["id"]
    people=append(people,person)
    json.NewEncoder(w).Encode(people)
}

//function to delete an existing person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){
     params:=mux.Vars(req)
     for index, item:=range(people){
       if item.ID==params["id"]{
         people=append(people[:index],people[index+1:]...)
         break
       }
     }
   json.NewEncoder(w).Encode(people)
}

//main function to call http
func main(){
   router:=mux.NewRouter()
   people=append(people,Person{ID:"1",Firstname:"Shahul",Lastname:"Hameed",Address: &Address{City:"Bangalore",State:"Karnataka"}})
   people=append(people,Person{ID:"2",Firstname:"Sohail",Lastname:"Hameed",Address: &Address{City:"Bangalore",State:"Karnataka"}})
   people=append(people,Person{ID:"3",Firstname:"ShahRukh",Lastname:"Khan",Address: &Address{City:"Bangalore",State:"Karnataka"}})
   router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
   router.HandleFunc("/people/{id}",GetPersonEndpoint).Methods("GET")
   router.HandleFunc("/people/{id}",CreatePersonEndpoint).Methods("POST")
   router.HandleFunc("/people/{id}",DeletePersonEndpoint).Methods("DELETE")
   log.Fatal(http.ListenAndServe(":12345",router))

}
