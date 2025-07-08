package contract

// Rule definitions with their corresponding error messages
var (
	// Acceptance rules
	Accepted   = ValidationRule{Name: "accepted", Message: "The :attribute must be accepted"}
	AcceptedIf = ValidationRule{
		Name:    "accepted_if",
		Message: "The :attribute must be accepted when :param0 is :param1",
	}
	AcceptedUnless = ValidationRule{
		Name:    "accepted_unless",
		Message: "The :attribute must be accepted unless :param0 is :param1",
	}
	AcceptedWith = ValidationRule{
		Name:    "accepted_with",
		Message: "The :attribute must be accepted when :param0 is present",
	}
	AcceptedWithout = ValidationRule{
		Name:    "accepted_without",
		Message: "The :attribute must be accepted when :param0 is not present",
	}
	Declined   = ValidationRule{Name: "declined", Message: "The :attribute must be declined"}
	DeclinedIf = ValidationRule{
		Name:    "declined_if",
		Message: "The :attribute must be declined when :param0 is :param1",
	}
	DeclinedUnless = ValidationRule{
		Name:    "declined_unless",
		Message: "The :attribute must be declined unless :param0 is :param1",
	}
	DeclinedWith = ValidationRule{
		Name:    "declined_with",
		Message: "The :attribute must be declined when :param0 is present",
	}
	DeclinedWithout = ValidationRule{
		Name:    "declined_without",
		Message: "The :attribute must be declined when :param0 is not present",
	}
	// Boolean rules
	Boolean = ValidationRule{Name: "boolean", Message: "The :attribute must be true or false"}
	// Comparison rules
	Gt         = ValidationRule{Name: "gt", Message: "The :attribute must be greater than :param0"}
	Lt         = ValidationRule{Name: "lt", Message: "The :attribute must be less than :param0"}
	Min        = ValidationRule{Name: "min", Message: "The :attribute must be at least :param0"}
	Max        = ValidationRule{Name: "max", Message: "The :attribute may not be greater than :param0"}
	Size       = ValidationRule{Name: "size", Message: "The :attribute must be :param0"}
	Same       = ValidationRule{Name: "same", Message: "The :attribute and :param0 must match"}
	Different  = ValidationRule{Name: "different", Message: "The :attribute and :param0 must be different"}
	StartsWith = ValidationRule{
		Name:    "starts_with",
		Message: "The :attribute must start with one of the following: :param0",
	}
	EndsWith = ValidationRule{Name: "ends_with", Message: "The :attribute must end with one of the following: :param0"}
	// Conditional rules
	Required   = ValidationRule{Name: "required", Message: "The :attribute field is required"}
	RequiredIf = ValidationRule{
		Name:    "required_if",
		Message: "The :attribute field is required when :param0 is :param1",
	}
	RequiredUnless = ValidationRule{
		Name:    "required_unless",
		Message: "The :attribute field is required unless :param0 is :param1",
	}
	RequiredWith = ValidationRule{
		Name:    "required_with",
		Message: "The :attribute field is required when :param0 is present",
	}
	RequiredWithout = ValidationRule{
		Name:    "required_without",
		Message: "The :attribute field is required when :param0 is not present",
	}
	RequiredWithAll = ValidationRule{
		Name:    "required_with_all",
		Message: "The :attribute field is required when :param0 is present",
	}
	RequiredWithoutAll = ValidationRule{
		Name:    "required_without_all",
		Message: "The :attribute field is required when none of :param0 are present",
	}
	// Control rules
	Filled    = ValidationRule{Name: "filled", Message: "The :attribute field must be filled"}
	Nullable  = ValidationRule{Name: "nullable", Message: "The :attribute field is nullable"}
	Present   = ValidationRule{Name: "present", Message: "The :attribute field must be present in the input"}
	Sometimes = ValidationRule{Name: "sometimes", Message: "The :attribute field is sometimes required"}
	Bail      = ValidationRule{Name: "bail", Message: "Stop validator on first failure"}
	// Database rules
	Unique = ValidationRule{Name: "unique", Message: "The :attribute has already been taken"}
	Exists = ValidationRule{Name: "exists", Message: "The selected :attribute is invalid"}
	// Date rules
	Date         = ValidationRule{Name: "date", Message: "The :attribute is not a valid date"}
	After        = ValidationRule{Name: "after", Message: "The :attribute must be a date after :param0"}
	AfterOrEqual = ValidationRule{
		Name:    "after_or_equal",
		Message: "The :attribute must be a date after or equal to :param0",
	}
	Before        = ValidationRule{Name: "before", Message: "The :attribute must be a date before :param0"}
	BeforeOrEqual = ValidationRule{
		Name:    "before_or_equal",
		Message: "The :attribute must be a date before or equal to :param0",
	}
	DateEquals = ValidationRule{Name: "date_equals", Message: "The :attribute must be a date equal to :param0"}
	DateFormat = ValidationRule{Name: "date_format", Message: "The :attribute does not match the format :param0"}
	// numeric rules
	Numeric    = ValidationRule{Name: "numeric", Message: "The :attribute must be a number"}
	NumericGt  = ValidationRule{Name: "gt", Message: "The :attribute must be greater than :param0"}
	NumericGte = ValidationRule{Name: "gte", Message: "The :attribute must be greater than or equal to :param0"}
	NumericLt  = ValidationRule{Name: "lt", Message: "The :attribute must be less than :param0"}
	NumericLte = ValidationRule{Name: "lte", Message: "The :attribute must be less than or equal to :param0"}
	Integer    = ValidationRule{Name: "integer", Message: "The :attribute must be an integer"}
	Decimal    = ValidationRule{Name: "decimal", Message: "The :attribute must have :param0 decimal places"}
	MultipleOf = ValidationRule{Name: "multiple_of", Message: "The :attribute must be a multiple of :param0"}
	// File rules
	File  = ValidationRule{Name: "file", Message: "The :attribute must be a file"}
	Image = ValidationRule{Name: "image", Message: "The :attribute must be an image"}
	Mimes = ValidationRule{Name: "mimes", Message: "The :attribute must be a file of type: :param0"}
	// special rules
	ActiveURL = ValidationRule{Name: "active_url", Message: "The :attribute must be a valid URL"}
	Confirmed = ValidationRule{Name: "confirmed", Message: "The :attribute confirmation does not match"}
	In        = ValidationRule{Name: "in", Message: "The selected :attribute is invalid"}
	IP        = ValidationRule{Name: "ip", Message: "The :attribute must be a valid IP address"}
	JSON      = ValidationRule{Name: "json", Message: "The :attribute must be a valid JSON string"}
	List      = ValidationRule{Name: "list", Message: "The :attribute must be a list of values"}
	Map       = ValidationRule{Name: "map", Message: "The :attribute must be a map"}
	NotIn     = ValidationRule{Name: "not_in", Message: "The selected :attribute is invalid"}
	Password  = ValidationRule{Name: "password", Message: "The :attribute must be at least :param0 characters long"}
	Regex     = ValidationRule{Name: "regex", Message: "The :attribute format is invalid"}
	// String rules
	Alpha     = ValidationRule{Name: "alpha", Message: "The :attribute may only contain letters"}
	Alphanum  = ValidationRule{Name: "alphanum", Message: "The :attribute may only contain letters and numbers"}
	AlphaDash = ValidationRule{
		Name:    "alpha_dash",
		Message: "The :attribute may only contain letters, numbers, dashes and underscores",
	}
	Email     = ValidationRule{Name: "email", Message: "The :attribute must be a valid email address"}
	String    = ValidationRule{Name: "string", Message: "The :attribute must be a string"}
	URL       = ValidationRule{Name: "url", Message: "The :attribute must be a valid URL"}
	UUID      = ValidationRule{Name: "uuid", Message: "The :attribute must be a valid UUID"}
	Lowercase = ValidationRule{Name: "lowercase", Message: "The :attribute must be lowercase"}
	Uppercase = ValidationRule{Name: "uppercase", Message: "The :attribute must be uppercase"}
	ASCII     = ValidationRule{Name: "ascii", Message: "The :attribute must only contain ASCII characters"}
	Ulid      = ValidationRule{Name: "ulid", Message: "The :attribute must be a valid ULID"}
	Slug      = ValidationRule{Name: "slug", Message: "The :attribute must be a valid slug"}
	// Auth rules
	CurrentPassword = ValidationRule{Name: "current_password", Message: "The :attribute is incorrect"}
	// Conditional prohibit rules
	Prohibited   = ValidationRule{Name: "prohibited", Message: "The :attribute field is prohibited"}
	ProhibitedIf = ValidationRule{
		Name:    "prohibited_if",
		Message: "The :attribute field is prohibited when :other is :value",
	}
	ProhibitedUnless = ValidationRule{
		Name:    "prohibited_unless",
		Message: "The :attribute field is prohibited unless :other is in :values",
	}
	Prohibits = ValidationRule{
		Name:    "prohibits",
		Message: "The :attribute field prohibits :other from being present",
	}
	// Additional string rules
	DoesntStartWith = ValidationRule{
		Name:    "doesnt_start_with",
		Message: "The :attribute must not start with one of the following: :values",
	}
	DoesntEndWith = ValidationRule{
		Name:    "doesnt_end_with",
		Message: "The :attribute must not end with one of the following: :values",
	}
	// Additional numeric rules
	Gte = ValidationRule{Name: "gte", Message: "The :attribute must be greater than or equal to :value"}
	Lte = ValidationRule{Name: "lte", Message: "The :attribute must be less than or equal to :value"}
)

// ValidationRule defines a struct that holds both the rule name and its error message
type ValidationRule struct {
	Name    string
	Message string
}

// Rule defines a single validator rule following Interface Segregation Principle.
type Rule interface {
	// Name returns the rule identifier
	Name() string

	// Validate performs the validator logic
	Validate(ctx RuleContext) error
}

// RuleContext provides validator context data for rules.
type RuleContext interface {
	// Field returns the field name being validated
	Field() string

	// Value returns the value being validated
	Value() any

	// Parameters returns rule parameters
	Parameters() []string

	// Data returns all data being validated
	Data() map[string]any

	// Attribute returns custom field name for messages
	Attribute(field string) string
}

// RuleFactory creates validator rules.
type RuleFactory interface {
	Create(name string, params []string) (Rule, error)
	Available() []string
}

// RuleCreator is a function type for creating rules
type RuleCreator func(parameters []string) (Rule, error)

// RuleRegistry manages rule registration and retrieval
type RuleRegistry interface {
	Register(name string, creator RuleCreator) error
	Get(name string) (RuleCreator, bool)
	List() []string
	Clear()
}
