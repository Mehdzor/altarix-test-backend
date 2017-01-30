package services

import (
	"altarix_test/model"
	"altarix_test/database"
	"altarix_test/socket"
)

func ListProducts() *[]model.Product {
	var list []model.Product
	database.DataSource.Find(&list)
	return &list
}

func AddProduct(p *model.Product) {
	database.DataSource.Create(p)
	socket.ActiveSocket.Send(model.Event{Product: p, Name: model.Sold})
}

func DeleteProduct(id int) *model.ApiError {
	var count int
	p := model.Product{ID: id}
	database.DataSource.First(&p).Count(&count)
	if count == 0 {
		err := model.ApiErrorForReason(model.ProductNotExistsOrSold)
		return &err
	}
	database.DataSource.Delete(&p)
	socket.ActiveSocket.Send(model.Event{Product: &p, Name: model.Bought})
	return nil
}