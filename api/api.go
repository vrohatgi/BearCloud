package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

//Declare a global array of Credentials
//See credentials.go

var users []Credentials = []Credentials{}

func RegisterRoutes(router *mux.Router) error {

	/*

		Fill out the appropriate get methods for each of the requests, based on the nature of the request.

		Think about whether you're reading, writing, or updating for each request


	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)

	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	cookie, err := request.Cookie("access_token")
	if err != nil {
		fmt.Fprintf(response, "")
		return
	}
	fmt.Fprintf(response, cookie.Value)
	return
}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	userID := request.URL.Query().Get("userID")
	fmt.Fprintf(response, userID)
	return
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(response, creds.Username + "\n" + creds.Password)
	return

}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/
	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}
	users = append(users, creds)
	response.WriteHeader(201)
	return


}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array

		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == creds.Username {
			fmt.Fprintf(response, strconv.Itoa(i))
			return
		}
	}
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return

}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == creds.Username {
			fmt.Fprintf(response, users[i].Password)
			return
		}
	}

	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == creds.Username {
			users[i].Password = creds.Password
			return
		}
	}

	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/
	creds := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == creds.Username && users[i].Password == creds.Password {
			users = remove(users, i)
			return
		}
	}

	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func remove(users []Credentials, index int) []Credentials {
	return append(users[:index], users[index+1:]...)
}
