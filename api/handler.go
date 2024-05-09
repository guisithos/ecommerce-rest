package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type productRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type productResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type envelope map[string]interface{}

func (app *application) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (app *application) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var productReq productRequest
	var payload jsonResponse

	err := app.readJSON(w, r, &productReq)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	var product = models.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
	}

	_, err = app.models.Product.Insert(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload = jsonResponse{
		Error:   false,
		Message: "Product successfully created",
		Data:    envelope{"product": product.Name},
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)

}

func (app *application) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	product, err := app.models.Product.GetOneById(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var productResponse = productResponse{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Product successfully obtained",
		Data:    envelope{"product": productResponse},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productReq productRequest
	var payload jsonResponse

	err := app.readJSON(w, r, &productReq)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.models.Product.GetOneById(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var product = models.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		UpdatedAt:   time.Now(),
		Id:          productID,
	}

	_, err = app.models.Product.Update(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload = jsonResponse{
		Error:   false,
		Message: "Product successfully updated",
		Data:    envelope{"product": product.Name},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) AllProducts(w http.ResponseWriter, r *http.Request) {
	var products models.Product
	all, err := products.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Data successfully obtained",
		Data:    envelope{"products": all},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.Product.DeleteByID(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Product successfully deleted",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}
