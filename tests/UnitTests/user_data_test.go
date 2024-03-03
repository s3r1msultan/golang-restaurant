package UnitTests

import (
	"final/models"
	"testing"
)

func TestUserValidation(t *testing.T) {
	// Test case 1: Valid user
	validUser := models.User{Email: "test@example.com", FirstName: "John", LastName: "Doe"}
	if err := models.ValidateUser(validUser); err != nil {
		t.Errorf("ValidateUser failed for valid user: %v", err)
	}

	// Test case 2: Empty email
	emptyEmailUser := models.User{FirstName: "Alice", LastName: "Smith"}
	if err := models.ValidateUser(emptyEmailUser); err == nil || err.Error() != "email is required" {
		t.Errorf("ValidateUser failed for empty email: expected 'email is required' error")
	}

	// Test case 3: Empty first name
	emptyFirstNameUser := models.User{Email: "test@example.com", LastName: "Johnson"}
	if err := models.ValidateUser(emptyFirstNameUser); err == nil || err.Error() != "first name is required" {
		t.Errorf("ValidateUser failed for empty first name: expected 'first name is required' error")
	}
	// Test case 4: Empty last name
	emptyLastNameUser := models.User{Email: "test@example.com", FirstName: "Emily"}
	if err := models.ValidateUser(emptyLastNameUser); err == nil || err.Error() != "last name is required" {
		t.Errorf("ValidateUser failed for empty last name: expected 'last name is required' error")
	}

}
