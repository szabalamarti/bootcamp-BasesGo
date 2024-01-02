package salary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSalary(t *testing.T) {
	t.Run("case 01: category A and 100 minutes", func(t *testing.T) {
		// Arrange
		category := "A"
		minutes := 100
		expected := 1666

		// Act
		result := GetSalary(minutes, category)

		// Assert
		require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
		t.Log("success")
	})

	t.Run("case 02: category B and 100 minutes", func(t *testing.T) {
		// Arrange
		category := "B"
		minutes := 100
		expected := 3000

		// Act
		result := GetSalary(minutes, category)

		// Assert
		require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
		t.Log("success")
	})

	t.Run("case 03: category C and 100 minutes", func(t *testing.T) {
		// Arrange
		category := "C"
		minutes := 100
		expected := 7500

		// Act
		result := GetSalary(minutes, category)

		// Assert
		require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
		t.Log("success")
	})
}
