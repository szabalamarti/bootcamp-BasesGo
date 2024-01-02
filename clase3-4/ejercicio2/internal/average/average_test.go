package average

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	// Arrange
	grades := []int{1, 2, 3, 4, 5}
	expected := 3.0

	// Act
	result := Average(grades...)

	// Assert
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
