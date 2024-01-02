package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimum(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3, 4, 5}
	expected := 1.0

	// Act
	result := Minimum(values...)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestMaximum(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3, 4, 5}
	expected := 5.0

	// Act
	result := Maximum(values...)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestAverage(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3, 4, 5}
	expected := 3.0

	// Act
	result := Average(values...)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
