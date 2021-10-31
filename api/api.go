package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func GetApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1. Performing Http Get...")

	respon, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		log.Fatalln(err)
	}

	defer respon.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(respon.Body)

	//conver respon body to string
	// bodyString := string(bodyBytes)
	// fmt.Println("API response as String:\n" + bodyString)

	// conver respon body to struct
	var todoStruct Todos
	json.Unmarshal(bodyBytes, &todoStruct)
	json.NewEncoder(w).Encode(todoStruct)
	// fmt.Printf("Api Response as Struct %v\n", todoStruct)

}

func GetApiId(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + Id)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	var todoStruct Todo
	json.Unmarshal(bodyByte, &todoStruct)
	json.NewEncoder(w).Encode(todoStruct)

}

func PostAPI() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// convert response to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// convert response to struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%v\n", todoStruct)

}

func PutApi() {
	fmt.Println("3. Performing Http Put...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}

	jsonReq, err := json.Marshal(todo)

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// convert respon to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// convert response tp struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)

	fmt.Printf("API responses struct %v \n", todoStruct)

}

func Delete() {
	fmt.Println("4. Performing Http Delete...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// change respon to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
