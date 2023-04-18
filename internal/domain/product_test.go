package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewProduct_CorrectFields(t *testing.T) {
	product := ddlProduct()

	assert.NotNil(t, product.ID)
	assert.Equal(t, product.Name, "DDL")
	assert.Equal(t, product.Description, "Dulce de leche")
	assert.Equal(t, product.Category, "dulces")
	assert.Equal(t, product.Price.Value, float32(250))
	assert.Equal(t, product.Stock.Amount, 30)
}

func Test_ProductCanConsumeStockAndThatIsDecreased(t *testing.T) {
	prod := ddlProduct()
	stockBeforeConsume := prod.Stock.Amount
	unitsToConsume := 4

	_ = prod.Consume(unitsToConsume)

	assert.Equal(t, prod.Stock.Amount, stockBeforeConsume-unitsToConsume)
}

func Test_ProductCanConsumeStockAndThatIsIncreased(t *testing.T) {
	prod := ddlProduct()
	stockBeforeRestore := prod.Stock.Amount
	unitsToRestore := 4

	_ = prod.Restore(unitsToRestore)

	assert.Equal(t, prod.Stock.Amount, stockBeforeRestore+unitsToRestore)
}

func TestProduct_KnowCannotConsumeStock(t *testing.T) {
	prod := ddlProduct()
	assert.False(t, prod.CanConsume(prod.Stock.Amount+1))
}

func TestProduct_KnowCanConsumeStock(t *testing.T) {
	prod := ddlProduct()

	assert.True(t, prod.CanConsume(prod.Stock.Amount))
}

func ddlProduct() Product {
	id, _ := uuid.Parse("aa2723cf-dd96-11ed-b17f-7c8ae169e5a7")
	price, _ := NewPrice(float32(250))
	stock, _ := NewStock(30)
	return Product{ID: id, Name: "DDL", Description: "Dulce de leche", Category: "dulces", Price: price, Stock: stock, IDSeller: uuid.Nil}
}
