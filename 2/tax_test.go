package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func TestCalculateTax(t *testing.T) {
	// Utilizando a lib testify
	// Testando a função para o valor 1000
	tax, err := CalculateTax(1000.0)

	// Verifica se o "err" veio em branco
	assert.Nil(t, err)
	// Verifica se a "tax" é igual a 10
	assert.Equal(t, 10.0, tax)

	// Testando a função para o valor 0
	tax, err = CalculateTax(0.0)
	// Se der erro, ele retorna uma mensagem
	assert.Error(t, err, "amount must be greater than 0")
	// Garantindo que o "tax" é 0
	assert.Equal(t, 0.0, tax)
	// Verificando se a mensagem de erro contem algo
	assert.Contains(t, err.Error(), "greater than 0")
}

// Testando o Mock
func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	// Salvo com Sucesso
	// Só posso chamar essa função com o valor de 10 1 vez (.Once())
	repository.On("SaveTax", 10.0).Return(nil).Once()

	// Erro ao Salvar
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	// Verifica se as chamadas foram feitas corretamente
	repository.AssertExpectations(t)

	// Verificando se o Mock foi chamado 2 vezes
	// dentro da função
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
