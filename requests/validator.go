package requests

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mizanalyst/mizanalyst/database"

	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register custom "exists" validation: validates that a record exists in the given table.column
	// Usage tag: validate:"exists=users.email"
	validate.RegisterValidation("exists", func(fl validator.FieldLevel) bool {
		param := fl.Param()
		parts := strings.SplitN(param, ".", 2)
		if len(parts) != 2 {
			return false
		}

		table, column := parts[0], parts[1]
		value := fl.Field().String()

		db := database.GetDB()
		var count int64
		db.Table(table).Where(fmt.Sprintf("%s = ?", column), value).Count(&count)
		return count > 0
	})

	// Register custom "unique" validation: validates that no record exists in the given table.column
	// Usage tag: validate:"unique=users.email"
	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		param := fl.Param()
		parts := strings.SplitN(param, ".", 2)
		if len(parts) != 2 {
			return false
		}

		table, column := parts[0], parts[1]
		value := fl.Field().String()

		db := database.GetDB()
		var count int64
		db.Table(table).Where(fmt.Sprintf("%s = ?", column), value).Count(&count)
		return count == 0
	})
}

// Validate runs struct validation and returns a map of field -> error message.
// Returns nil if validation passes.
func Validate(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		errs[field] = formatValidationError(e)
	}

	return errs
}

// GetValidator returns the underlying validator instance for custom registrations.
func GetValidator() *validator.Validate {
	return validate
}

// SetDB allows overriding the DB used for exists/unique validations (useful for testing).
// This is a no-op; the validators use database.GetDB() directly.
// Kept for API compatibility if future needs arise.
func SetDB(_ *gorm.DB) {}

// formatValidationError returns a human-readable error message for a validation error.
func formatValidationError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s characters", e.Param())
	case "exists":
		return "Record does not exist"
	case "unique":
		return "This value is already taken"
	default:
		return fmt.Sprintf("Failed on '%s' validation", e.Tag())
	}
}
