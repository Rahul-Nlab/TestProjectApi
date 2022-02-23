package main

//THE ORGANISED VERSION
import (
	"fmt"
	"log"
	"net/http"
	"testproject/business"
	"testproject/databases"
	"testproject/handlers"

	muxhan "github.com/gorilla/handlers"
	"github.com/labstack/echo/v4"

	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := databases.Setup()
	defer db.Close()

	// h := handlers.Handlers{
	// 	User: business.New(db),
	// }

	h := handlers.New(business.New(db))

	// router := mux.NewRouter()
	router := echo.New()

	fmt.Println("Up!")

	// router.HandleFunc("/v1/users/", h.GetUsersRequest).Methods(http.MethodGet)
	router.GET("/v1/users/", h.GetUsersRequest())
	router.HandleFunc("/v1/users/{id}/", h.GetUsersRequest).Methods(http.MethodGet)

	router.HandleFunc("/v1/users/", h.CreateUserRequest).Methods(http.MethodPost)
	router.HandleFunc("/v1/users/{id}/", h.CreateUserRequest).Methods(http.MethodPost)

	router.HandleFunc("/v1/users/{id}/", h.ChangeUserRequest).Methods(http.MethodPut)
	router.HandleFunc("/v1/users/{id}/", h.DeleteUserRequest).Methods(http.MethodDelete)

	router.HandleFunc("/v1/users/{id}/addresses/", h.GetAddressesRequest).Methods(http.MethodGet)

	// router.HandleFunc("/v1/addresses/", h.GetAddresses).Methods(http.MethodGet)
	// router.HandleFunc("/v1/addresses/", h.CreateAddresses).Methods(http.MethodPost)

	headersOk := muxhan.AllowedHeaders([]string{"*"})
	originsOk := muxhan.AllowedOrigins([]string{"*"})
	methodsOk := muxhan.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":5434", muxhan.CORS(headersOk, originsOk, methodsOk)(router)))
}
