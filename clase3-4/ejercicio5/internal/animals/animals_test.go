package animals

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDog(t *testing.T) {
	// Arrange
	value := 5
	expected := 50.0

	// Act
	result := animalDog(value)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestCat(t *testing.T) {
	// Arrange
	value := 5
	expected := 25.0

	// Act
	result := animalCat(value)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestHamster(t *testing.T) {
	// Arrange
	value := 5
	expected := 1.25

	// Act
	result := animalHamster(value)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestTarantula(t *testing.T) {
	// Arrange
	value := 5
	expected := 0.75

	// Act
	result := animalTarantula(value)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
