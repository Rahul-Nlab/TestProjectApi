package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testproject/business"

	"github.com/gorilla/mux"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Handlers struct {
	user business.Business
}

func New(user business.Business) Handlers {
	return Handlers{
		user: user,
	}
}

func (h Handlers) GetUsersRequest(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("User GET served")
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Ge`t("Origin"))
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	
	vars := mux.Vars(r)
	id := vars["id"]
	
	userStruct, e := h.user.GetUsers(id)
	
	if e != "" {
		w.Write([]byte(e))
		return
	}
	
	json.NewEncoder(w).Encode(userStruct)
}

func (h Handlers) GetAddressesRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Addresses GET served")

	vars := mux.Vars(r)
	id := vars["id"]

	addressStruct, e := h.user.AddressOfUser(id)

	if e != "" {
		w.Write([]byte(e))
		return
	}

	json.NewEncoder(w).Encode(addressStruct)

// NOW, THIS HAS TO BE TRANSFERED TO THE BUSINESS PORTION
// 	w.Write([]byte(str))
	
// 	fmt.Println("Getting Address Data...")
// 	rows, err := h.Db.Query("SELECT * FROM Addresses")
// 	CheckErr(err)
// 	var tempAddressStruct []Addresses
// 	for rows.Next() {
// 		var addressId, pincode int
// 		var street, area, city string

// 		err = rows.Scan(&addressId, &street, &area, &pincode, &city)
// 		CheckErr(err)

// 		tempAddressStruct = append(tempAddressStruct, Addresses{A_id: addressId, Street: street, Area: area, Pincode: pincode, City: city})
// 	}
}

func (h Handlers) CreateUserRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, e := ioutil.ReadAll(r.Body)
	if e != nil {
		w.Write([]byte("Please enter a valid json input!"))
	}

	str := h.user.CreateUsers(id, reqBody)
	w.Write([]byte(str))
	fmt.Println("User POST served")
}

func (h Handlers) ChangeUserRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, e := ioutil.ReadAll(r.Body)
	if e != nil {
		w.Write([]byte("Please enter a valid json input!"))
	}

	str := h.user.ChangeUser(id, reqBody)

	w.Write([]byte(str))
	fmt.Println("User PUT served")
}

func (h Handlers) DeleteUserRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	str := h.user.DeleteUser(id)

	w.Write([]byte(str))
	fmt.Println("User DELETE served")
}

// func (h HttpHandlers) CreateAddresses(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Reading the request...")
// 	reqBody, e := ioutil.ReadAll(r.Body)
// 	CheckErr(e)
// 	var newAddress Addresses
// 	json.Unmarshal(reqBody, &newAddress)
// 	h.Db.Exec("INSERT INTO Addresses(street, area , pincode, city) VALUES (newAddress);")
// 	// fmt.Fprintf(w, "%+v", string(reqBody))

// 	json.NewEncoder(w).Encode(newAddress)
// 	fmt.Println("Successfully Added!")
// }
