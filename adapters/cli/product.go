package cli

import (
	"fmt"

	"github.com/codeedu/go-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string, produtId string,
	productName string,
	price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enabled":
		product, err := service.Get(produtId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disabled":
		product, err := service.Get(produtId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	default:
		res, err := service.Get(produtId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s \nName: %s \nPrice: %f \nStatus: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())

	}
	return result, nil

}
