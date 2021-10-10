package ptinvoice

import (
	"fmt"
	"strings"
)

type ParserError struct {
	error

	FieldNumber int
}

func Parse(input string) (*Fields, error) {
	parsedFields := &Fields{}

	rawFields := strings.Split(input, "*")
	for i, rawField := range rawFields {
		rawFieldParts := strings.SplitN(rawField, ":", 2)
		if len(rawFieldParts) != 2 {
			err := ParserError{
				error:       fmt.Errorf("missing separator in field %d", i),
				FieldNumber: i,
			}

			return nil, err
		}

		code := FieldCode(rawFieldParts[0])
		value := rawFieldParts[1]

		fieldLength, isValidCode := FieldLengthsByCode[code]
		if !isValidCode {
			err := ParserError{
				error:       fmt.Errorf("invalid code %s in field %d", string(code), i),
				FieldNumber: i,
			}

			return nil, err
		}

		if valueLength := len(value); valueLength > fieldLength {
			err := ParserError{
				error:       fmt.Errorf("value length exceeded for field %d; wanted %d characters but got %d instead", i, valueLength, fieldLength),
				FieldNumber: i,
			}

			return nil, err
		}

		parsedFields.Set(code, value)
	}

	return parsedFields, nil
}
