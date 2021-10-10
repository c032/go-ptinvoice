package ptinvoice

import (
	"fmt"
	"sort"
)

type FieldCode string

const (
	FieldCodeA  FieldCode = "A"
	FieldCodeB  FieldCode = "B"
	FieldCodeC  FieldCode = "C"
	FieldCodeD  FieldCode = "D"
	FieldCodeE  FieldCode = "E"
	FieldCodeF  FieldCode = "F"
	FieldCodeG  FieldCode = "G"
	FieldCodeH  FieldCode = "H"
	FieldCodeI1 FieldCode = "I1"
	FieldCodeI2 FieldCode = "I2"
	FieldCodeI3 FieldCode = "I3"
	FieldCodeI4 FieldCode = "I4"
	FieldCodeI5 FieldCode = "I5"
	FieldCodeI6 FieldCode = "I6"
	FieldCodeI7 FieldCode = "I7"
	FieldCodeI8 FieldCode = "I8"
	FieldCodeJ1 FieldCode = "J1"
	FieldCodeJ2 FieldCode = "J2"
	FieldCodeJ3 FieldCode = "J3"
	FieldCodeJ4 FieldCode = "J4"
	FieldCodeJ5 FieldCode = "J5"
	FieldCodeJ6 FieldCode = "J6"
	FieldCodeJ7 FieldCode = "J7"
	FieldCodeJ8 FieldCode = "J8"
	FieldCodeK1 FieldCode = "K1"
	FieldCodeK2 FieldCode = "K2"
	FieldCodeK3 FieldCode = "K3"
	FieldCodeK4 FieldCode = "K4"
	FieldCodeK5 FieldCode = "K5"
	FieldCodeK6 FieldCode = "K6"
	FieldCodeK7 FieldCode = "K7"
	FieldCodeK8 FieldCode = "K8"
	FieldCodeL  FieldCode = "L"
	FieldCodeM  FieldCode = "M"
	FieldCodeN  FieldCode = "N"
	FieldCodeO  FieldCode = "O"
	FieldCodeP  FieldCode = "P"
	FieldCodeQ  FieldCode = "Q"
	FieldCodeR  FieldCode = "R"
	FieldCodeS  FieldCode = "S"
)

var FieldLengthsByCode = map[FieldCode]int{
	FieldCodeA:  9,
	FieldCodeB:  30,
	FieldCodeC:  12,
	FieldCodeD:  2,
	FieldCodeE:  1,
	FieldCodeF:  8,
	FieldCodeG:  60,
	FieldCodeH:  70,
	FieldCodeI1: 5,
	FieldCodeI2: 16,
	FieldCodeI3: 16,
	FieldCodeI4: 16,
	FieldCodeI5: 16,
	FieldCodeI6: 16,
	FieldCodeI7: 16,
	FieldCodeI8: 16,
	FieldCodeJ1: 5,
	FieldCodeJ2: 16,
	FieldCodeJ3: 16,
	FieldCodeJ4: 16,
	FieldCodeJ5: 16,
	FieldCodeJ6: 16,
	FieldCodeJ7: 16,
	FieldCodeJ8: 16,
	FieldCodeK1: 5,
	FieldCodeK2: 16,
	FieldCodeK3: 16,
	FieldCodeK4: 16,
	FieldCodeK5: 16,
	FieldCodeK6: 16,
	FieldCodeK7: 16,
	FieldCodeK8: 16,
	FieldCodeL:  16,
	FieldCodeM:  16,
	FieldCodeN:  16,
	FieldCodeO:  16,
	FieldCodeP:  16,
	FieldCodeQ:  4,
	FieldCodeR:  4,
	FieldCodeS:  65,
}

var FieldCodes []FieldCode

type Fields map[FieldCode]string

func (f Fields) Get(code FieldCode) string {
	return f[code]
}

func (f Fields) Set(code FieldCode, value string) {
	f[code] = value
}

func (f Fields) Len() int {
	return len(f)
}

func (f Fields) String() string {
	var encoded string

	for _, key := range FieldCodes {
		value, ok := f[key]
		if !ok {
			continue
		}
		if encoded != "" {
			encoded += "*"
		}

		encoded += fmt.Sprintf("%s:%s", key, value)
	}

	return encoded
}

func init() {
	var codes []string
	for key, _ := range FieldLengthsByCode {
		codes = append(codes, string(key))
	}

	sort.Strings([]string(codes))

	FieldCodes = nil
	for _, code := range codes {
		FieldCodes = append(FieldCodes, FieldCode(code))
	}
}
