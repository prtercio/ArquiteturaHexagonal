package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestGetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	actualID := product.GetID()
	require.Equal(t, product.ID, actualID)
}

func TestGetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"

	actualName := product.GetName()
	require.Equal(t, product.Name, actualName)
}

func TestGetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = "ENABLED"

	actualStatus := product.GetStatus()
	require.Equal(t, product.Status, actualStatus)
}

func TestGetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	actualPrice := product.GetPrice()
	require.Equal(t, product.Price, actualPrice)
}
