package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(Dbuser, DbPassword, "products_db")
	if err != nil {
		log.Fatal("Error occured while initialising db test")
	}

	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10, 7),
		PRIMARY KEY(id)
	);`

	_, err := a.Db.Exec(createTableQuery)
	if err != nil {
		log.Fatalln(err)
	}
}

func clearTable() {
	a.Db.Exec("DELETE FROM products")
	a.Db.Exec("ALTER TABLE products AUTO_INCREMENT=1")
	// log.Println("ClearTable")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES ('%v', %v, %v)", name, quantity, price)
	_, err := a.Db.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func TestGetProductAPI(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 500)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)

	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}

func TestCreateProductAPI(t *testing.T) {
	clearTable()
	var product = []byte(`{"name":"chair", "quantity":1, "price":100}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))

	req.Header.Set("Content-Type", "application/json")
	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "chair" {
		t.Errorf("Expected name: %v, Got: %v", "chair", m["name"])
	}
	// log.Printf("%T", m["quantity"])
	// marshal converts int to float by default
	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	}
	if m["price"] != 100.0 {
		t.Errorf("Expected price: %v, Got: %v", 100.0, m["price"])
	}
}

func TestDeleteProductAPI(t *testing.T) {
	clearTable()
	addProduct("connector", 10, 10)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProductAPI(t *testing.T) {
	clearTable()
	addProduct("connector", 10, 10)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(req)

	var oldValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &oldValue)

	var product = []byte(`{"name":"connector", "quantity":1, "price":10}`)
	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(product))

	req.Header.Set("Content-Type", "application/json")

	response = sendRequest(req)
	var newValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &newValue)

	if oldValue["id"] != newValue["id"] {
		t.Errorf("Expected id: %v, Got: %v ", newValue["id"], oldValue["id"])
	}

	if oldValue["name"] != newValue["name"] {
		t.Errorf("Expected name: %v, Got: %v ", newValue["name"], oldValue["name"])
	}

	if oldValue["price"] != newValue["price"] {
		t.Errorf("Expected price: %v, Got: %v ", newValue["price"], oldValue["price"])
	}

	if oldValue["quantity"] == newValue["quantity"] {
		t.Errorf("Expected quantity: %v, Got: %v ", newValue["quantity"], oldValue["quantity"])
	}

}
