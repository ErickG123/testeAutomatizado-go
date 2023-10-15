package tax

import (
	"testing"
)

// Criando uma função para fazer o teste
// t *testing.T => Testes Comuns
func TestCalculateTax(t *testing.T) {
	// Rodar Test: go test .
	// Rodar Test: go test -v

	// Cobertura de Código: go test -coverprofile="coverage.out"
	// O coverage verifica se testas tudo que dava

	// Verificar a Cobertura: go tool cover -html="coverage.out"

	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// Teste de Benchmark
// b *testing.B => Teste de Benchmark
func BenchmarkCalculateTax(b *testing.B) {
	// Rodar o Teste: go test -bench=.
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// Fuzzing: Teste de Mutação
// f *testing.F => Teste de Fuzzing
func FuzzCalculateTax(f *testing.F) {
	// Rodar o Teste: go test -fuzz=.
	// Seeding
	seed := []float64{0, -1, -2, -2.5, 500, 1000, 1500, 10000, 20000}
	for _, amount := range seed {
		// Adicionando os valores ao Fuzzing para fazer o teste (seeding)
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
