//go:build example
// +build example

package main

import (
	"fmt"
	"time"

	scgvalidator "github.com/next-trace/scg-validator/validator"
)

func main() {
	fmt.Println("=== SCG-Validator Comprehensive Feature Demo ===")
	fmt.Println()

	// Demo 1: Basic Validation
	basicValidationDemo()

	// Demo 2: Custom Messages and Attributes
	customMessagesDemo()

	// Demo 3: Advanced Validation Rules
	advancedRulesDemo()

	// Demo 4: Bail Rule Demonstration
	bailRuleDemo()

	// Demo 5: Complex Data Validation
	complexDataDemo()

	// Demo 6: Request Isolation Demo
	requestIsolationDemo()

	fmt.Println("\n=== All Demos Completed Successfully! ===")
}

// basicValidationDemo demonstrates basic validation functionality
func basicValidationDemo() {
	fmt.Println("1. Basic Validation Demo")
	fmt.Println("========================")

	validator := scgvalidator.New()

	// Example user registration data
	userData := map[string]interface{}{
		"name":     "John Doe",
		"email":    "john@example.com",
		"age":      25,
		"website":  "https://example.com",
		"password": "securepass123",
	}

	// Define validation rules
	rules := map[string]string{
		"name":     "required|alpha",
		"email":    "required|email",
		"age":      "required|numeric|min:18|max:100",
		"website":  "url",
		"password": "required|min:8",
	}

	result := validator.ValidateWithResult(userData, rules)

	if result.IsValid() {
		fmt.Println("✅ User registration data is valid!")
	} else {
		fmt.Println("❌ Validation failed:")
		printErrors(result)
	}
	fmt.Println()
}

// customMessagesDemo demonstrates custom messages and attributes
func customMessagesDemo() {
	fmt.Println("2. Custom Messages and Attributes Demo")
	fmt.Println("=====================================")

	validator := scgvalidator.New()

	// Set custom messages
	validator.SetCustomMessage("required", "The :attribute field is absolutely required!")
	validator.SetCustomMessage("email", "Please provide a valid :attribute address")
	validator.SetCustomMessage("min", "The :attribute must be at least :param0 characters long")

	// Set custom attributes
	validator.SetCustomAttribute("email", "Email Address")
	validator.SetCustomAttribute("password", "Password")
	validator.SetCustomAttribute("name", "Full Name")

	// Test data with validation errors
	invalidData := map[string]interface{}{
		"name":     "",
		"email":    "invalid-email",
		"password": "123",
	}

	rules := map[string]string{
		"name":     "required",
		"email":    "required|email",
		"password": "required|min:8",
	}

	result := validator.ValidateWithResult(invalidData, rules)
	fmt.Println("Custom message validation result:")
	printErrors(result)
	fmt.Println()
}

// advancedRulesDemo demonstrates various advanced validation rules
func advancedRulesDemo() {
	fmt.Println("3. Advanced Validation Rules Demo")
	fmt.Println("=================================")

	validator := scgvalidator.New()

	// Comprehensive test data
	testData := map[string]interface{}{
		"username":    "john_doe123",
		"email":       "john@example.com",
		"phone":       "+1234567890",
		"age":         25,
		"score":       85.5,
		"is_active":   true,
		"tags":        []string{"developer", "golang"},
		"website":     "https://johndoe.dev",
		"ip_address":  "192.168.1.1",
		"uuid":        "550e8400-e29b-41d4-a716-446655440000",
		"json_config": `{"theme": "dark", "lang": "en"}`,
		"birth_date":  "1998-05-15",
	}

	// Advanced validation rules (using available rules)
	rules := map[string]string{
		"username":    "required|alpha_num|min:3|max:20",
		"email":       "required|email",
		"phone":       "required|min:10|max:15",
		"age":         "required|numeric|between:18,65",
		"score":       "numeric|between:0,100",
		"is_active":   "boolean",
		"tags":        "required",
		"website":     "url",
		"ip_address":  "required",
		"uuid":        "required|min:36|max:36",
		"json_config": "required",
		"birth_date":  "required",
	}

	result := validator.ValidateWithResult(testData, rules)

	if result.IsValid() {
		fmt.Println("✅ All advanced validations passed!")
	} else {
		fmt.Println("❌ Advanced validation failed:")
		printErrors(result)
	}
	fmt.Println()
}

// bailRuleDemo demonstrates the bail rule functionality
func bailRuleDemo() {
	fmt.Println("4. Bail Rule Demo")
	fmt.Println("=================")

	validator := scgvalidator.New()

	// Data that will fail multiple rules
	testData := map[string]interface{}{
		"email": "invalid",
	}

	// Without bail - will show all validation errors
	fmt.Println("Without bail rule:")
	rules1 := map[string]string{
		"email": "required|email|min:10",
	}
	result1 := validator.ValidateWithResult(testData, rules1)
	printErrors(result1)

	// With bail - will stop at first failure
	fmt.Println("With bail rule:")
	rules2 := map[string]string{
		"email": "bail|required|email|min:10",
	}
	result2 := validator.ValidateWithResult(testData, rules2)
	printErrors(result2)
	fmt.Println()
}

// complexDataDemo demonstrates validation of complex nested data
func complexDataDemo() {
	fmt.Println("5. Complex Data Validation Demo")
	fmt.Println("===============================")

	validator := scgvalidator.New()

	// Complex user profile data
	profileData := map[string]interface{}{
		"user_id":          12345,
		"username":         "johndoe",
		"email":            "john@example.com",
		"profile_complete": true,
		"last_login":       time.Now().Format("2006-01-02"),
		"preferences":      `{"theme": "dark", "notifications": true}`,
		"social_links":     []string{"https://twitter.com/johndoe", "https://github.com/johndoe"},
		"bio":              "Software developer passionate about Go",
		"location":         "San Francisco, CA",
		"website":          "https://johndoe.dev",
	}

	// Complex validation rules (using available rules)
	rules := map[string]string{
		"user_id":          "required|numeric|min:1",
		"username":         "required|alpha_num|min:3|max:20",
		"email":            "required|email",
		"profile_complete": "boolean",
		"last_login":       "required",
		"preferences":      "required",
		"social_links":     "required",
		"bio":              "max:500",
		"location":         "max:100",
		"website":          "url",
	}

	result := validator.ValidateWithResult(profileData, rules)

	if result.IsValid() {
		fmt.Println("✅ Complex profile data validation passed!")
		fmt.Printf("   Validated %d fields successfully\n", len(profileData))
	} else {
		fmt.Println("❌ Complex validation failed:")
		printErrors(result)
	}
	fmt.Println()
}

// requestIsolationDemo demonstrates request isolation between validator instances
func requestIsolationDemo() {
	fmt.Println("6. Request Isolation Demo")
	fmt.Println("=========================")

	// Create two separate validator instances
	validator1 := scgvalidator.New()
	validator2 := scgvalidator.New()

	// Set different custom messages for each validator
	validator1.SetCustomMessage("required", "Validator 1: This field is required!")
	validator2.SetCustomMessage("required", "Validator 2: Field cannot be empty!")

	testData := map[string]interface{}{
		"field": "",
	}
	rules := map[string]string{
		"field": "required",
	}

	// Validate with first validator
	result1 := validator1.ValidateWithResult(testData, rules)
	fmt.Println("Validator 1 result:")
	printErrors(result1)

	// Validate with second validator
	result2 := validator2.ValidateWithResult(testData, rules)
	fmt.Println("Validator 2 result:")
	printErrors(result2)

	fmt.Println("✅ Request isolation working correctly - different custom messages!")
	fmt.Println()
}

// printErrors is a helper function to print validation errors
func printErrors(result interface{}) {
	if errorResult, ok := result.(interface{ Errors() map[string][]string }); ok {
		for field, messages := range errorResult.Errors() {
			for _, message := range messages {
				fmt.Printf("  - %s: %s\n", field, message)
			}
		}
	}
}
