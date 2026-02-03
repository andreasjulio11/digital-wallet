package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// logika perhitungan (Unit Test)
func TestTransactionLogic(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		initialBalance := 100000
		withdrawAmount := 40000

		expected := 60000
		result := initialBalance - withdrawAmount

		assert.Equal(t, expected, result)
	})

	t.Run("Periksa Kembali Saldo", func(t *testing.T) {
		initialBalance := 20000
		withdrawAmount := 50000

		isAllowed := initialBalance >= withdrawAmount

		assert.False(t, isAllowed, "Seharusnya tidak boleh narik kalau saldo kurang")
	})
}
