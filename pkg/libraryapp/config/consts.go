package config

// default setting values
const (
	SETTING_KEY_LOAN_PERIOD           = "LOAN_PERIOD"
	SETTING_DEFAULT_LOAN_PERIOD       = 7 // days
	SETTING_KEY_MAX_LOAN_PER_USER     = "MAX_LOAN_PER_USER"
	SETTING_DEFAULT_MAX_LOAN_PER_USER = 3 // active loan limit
	SETTING_KEY_FINE_PER_DAY          = "FINE_PER_DAY"
	SETTING_DEFAULT_FINE_PER_DAY      = 1 // unit currency
)

// api default values
const (
	API_DEFAULT_LIMIT = 10
	API_DEFAULT_SKIP  = 0
)

// environment variable keys
const (
	ENV_KEY_DB_URL  = "DB_URL"
	ENV_KEY_APP_ENV = "APP_ENV"
)
