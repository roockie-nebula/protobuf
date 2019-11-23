package duration

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// IsNull will return if the current duration is null
func (m *Duration) IsNull() bool {

	if m == nil {
		return true
	}

	// we use IsNotNull instead of IsNull to make sure that a duration is
	// initialized as null value
	return !m.IsNotNull
}

// Set will set the duration to the given duration
func (m *Duration) Set(value time.Duration) {

	if m == nil {
		*m = Duration{}
	}

	m.Nanos = value.Nanoseconds()
	m.IsNotNull = true
}

// SetNull will clear the duration
func (m *Duration) SetNull() {

	if m == nil {
		return
	}

	m.Nanos = 0
	m.IsNotNull = false
}

// Duration returns a golang duration object
func (m *Duration) Duration() time.Duration {

	if m.IsNull() {
		return 0
	}

	return time.Duration(m.Nanos)
}

// Scan implements the Scanner interface of the database driver
func (m *Duration) Scan(value interface{}) error {

	// initialize duration if pointer is nil
	if m == nil {
		*m = Duration{}
	}

	// convert the interface to a duration type
	dbDuration, ok := value.(time.Duration)

	if ok {
		m.Nanos = dbDuration.Nanoseconds()
		m.IsNotNull = true
		return nil
	}

	m.Nanos = 0
	m.IsNotNull = false
	return nil
}

// Value implements the db driver Valuer interface
func (m Duration) Value() (driver.Value, error) {

	if m.IsNull() {
		return nil, nil
	}

	return time.Duration(m.Nanos), nil
}

// ImplementsGraphQLType is required by the graphql custom scalar interface
// this defines the name used in the schema to declare a null duration type
func (m *Duration) ImplementsGraphQLType(name string) bool {
	return name == "Duration"
}

// UnmarshalGraphQL is required by the graphql custom scalar interface
// this wraps the null duration
func (m *Duration) UnmarshalGraphQL(input interface{}) error {
	// initialize duration if pointer is nil
	if m == nil {
		*m = Duration{}
	}

	switch input := input.(type) {
	case Duration:
		m.IsNotNull = input.IsNotNull
		m.Nanos = input.Nanos
		return nil
	case time.Duration:
		m.Nanos = input.Nanoseconds()
		m.IsNotNull = true
		return nil
	case string:
		duration, err := time.ParseDuration(input)
		if err != nil {
			return err
		}

		m.Set(duration)
		return nil
	case nil:
		m.SetNull()
		return nil
	default:
		fmt.Printf("%T\n", input)
		fmt.Println(input)
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON will return the content as json value, this is also called
// by graphql to generate the response
func (m Duration) MarshalJSON() ([]byte, error) {

	if m.IsNull() {
		return []byte("null"), nil
	}

	// TODO: format the timestamp in iso compatible duration format

	return []byte(m.Duration().String()), nil
}

// UnmarshalJSON is used to convert the json representation into a duration
func (m *Duration) UnmarshalJSON(input []byte) error {

	// trim the leading and trailing quotes from the duration
	cleanInput := bytes.Trim(input, "\"")

	// convert to string
	asString := string(cleanInput)

	if asString == "null" {
		m.SetNull()
		return nil
	}

	duration, err := time.ParseDuration(asString)
	if err != nil {
		return err
	}

	m.Set(duration)
	return nil
}
