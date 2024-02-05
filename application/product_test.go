package application_test

import (
	"testing"

	"github.com/IsaqueAmorim/hexagonal-architecture/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Enable()
	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Price = 0
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Price = 10
	product.Status = application.DISABLED

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid status"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	id := uuid.NewV4().String()
	product.Id = id

	require.Equal(t, id, product.GetId())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"

	require.Equal(t, "Product Test", product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_NewProduct(t *testing.T) {
	product, err := application.NewProduct(uuid.NewV4().String(), "Product Test", 10, application.ENABLED)
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
}
