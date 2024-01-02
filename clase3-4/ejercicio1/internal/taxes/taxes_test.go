package taxes

import (
	"testing"
)

func TestGetTax_Under(t *testing.T) {
	// Arrange
	salary := 49000.0
	expected := 0.0

	// Act
	result := GetTax(salary)

	// Assert
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
		return
	}
	t.Log("success")

}

func TestGetTax_InBetween(t *testing.T) {
	// Arrange
	salary := 100000.0
	expected := 17000.0

	// Act
	result := GetTax(salary)

	// Assert
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
		return
	}
	t.Log("success")

}

func TestGetTax_Over(t *testing.T) {
	// Arrange
	salary := 151000.0
	expected := 40770.0

	// Act
	result := GetTax(salary)

	// Assert
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
		return
	}
	t.Log("success")

}
