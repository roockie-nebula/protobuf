package nullfloat

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// IsNull will return if the current null float is null
func (nf *NullFloat) IsNull() bool {

	if nf == nil {
		return true
	}

	// we use IsNotNull instead of IsNull to make sure that a float is
	// initialized as null value
	return nf.IsNotNull == false
}

// Set will set the null float to the given value
func (nf *NullFloat) Set(value float64) {

	if nf == nil {
		*nf = NullFloat{}
	}

	nf.Float = value
	nf.IsNotNull = true

}

// SetNull will set the null float to null
func (nf *NullFloat) SetNull() {

	if nf == nil {
		*nf = NullFloat{}
	}

	nf.Float = 0
	nf.IsNotNull = false

}

// Scan implements the Scanner interface of the database driver
func (nf *NullFloat) Scan(value interface{}) error {

	// if the null float is nil, initialize it
	if nf == nil {
		*nf = NullFloat{}
	}

	// if the value is nil, reset the data of the null float
	if value == nil {

		nf.Float = 0
		nf.IsNotNull = false
		return nil

	}

	// create a sql NullFloat64 to use the not exported convertAssign-method
	// of the golang sql package
	sqlFloat := sql.NullFloat64{}

	// scan the value, using the sql package
	err := sqlFloat.Scan(value)
	if err != nil {
		return err
	}
	nf.IsNotNull = true
	nf.Float = sqlFloat.Float64
	return nil
}

// Value implements the db driver Valuer interface
func (nf NullFloat) Value() (driver.Value, error) {
	if nf.IsNull() {
		return nil, nil
	}
	return nf.Float, nil
}

// ImplementsGraphQLType is required by the graphql custom scalar interface
// this defines the name used in the schema to declare a null float type
func (nf *NullFloat) ImplementsGraphQLType(name string) bool {
	return name == "Float"
}

// UnmarshalGraphQL is required by the graphql custom scalar interface
// this wraps the null float
func (nf *NullFloat) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {

	case NullFloat:
		nFloat := input
		nf.IsNotNull = nFloat.IsNotNull
		nf.Float = nFloat.Float
		return nil

	case float32:
		nf.Float = float64(input)
		nf.IsNotNull = true
		return nil

	case float64:
		nf.Float = input
		nf.IsNotNull = true
		return nil

	default:
		fmt.Printf("%T\n", input)
		fmt.Println(input)
		return fmt.Errorf("wrong type")
	}
}
