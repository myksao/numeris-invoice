package utils

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type PaginationReq struct {
	Offset string `form:"offset" binding:"required,number,gte=0"`
	Limit  string `form:"limit" binding:"required,number,gte=1,required_with=offset"`
}

// Function to convert a time.Time object to a cron expression
func ConvertToCron(dueDate time.Time) string {
	// Extract individual components from the due date
	second := dueDate.Second()
	minute := dueDate.Minute()
	hour := dueDate.Hour()
	dayOfMonth := dueDate.Day()
	month := dueDate.Month()
	weekday := dueDate.Weekday()

	// Convert weekday from time.Weekday to cron format (0-6, where Sunday = 0)
	cronWeekday := (int(weekday) + 1) % 7

	// Create a cron expression string in the format: "second minute hour day month weekday"
	cronExpression := fmt.Sprintf("%d %d %d %d %d %d", second, minute, hour, dayOfMonth, int(month), cronWeekday)

	return cronExpression
}

func CustomValidationError(err error) []any {
	// validate := validator.New(validator.WithRequiredStructEnabled())

	errorMessage := make([]any, 0)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			if err.Tag() == "required" {
				errorMessage = append(errorMessage, fmt.Sprintf("Field %s is required", err.Field()))
			}

			if err.Value() != "" {
				errorMessage = append(errorMessage, fmt.Sprintf("Field %s must be a valid %s", err.Field(), err.Tag()))
			}

			if err.Tag() == "oneof" {
				errorMessage = append(errorMessage, fmt.Sprintf("Field %s must be one of %s", err.Field(), err.Param()))
			}
		}

		// from here you can create your own error messages in whatever language you wish
		return errorMessage
	}
	return errorMessage
}

// HashPassword hashes a plain-text password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plain-text password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
