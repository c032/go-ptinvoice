package ptinvoice_test

import (
	"testing"

	"github.com/c032/go-ptinvoice"
)

func TestParseMessage_Invoice(t *testing.T) {
	const input = `A:123456789*B:999999990*C:PT*D:FT*E:N*F:20191231*G:FTAB2019/0035*H:CSDF7T5H-0035*I1:PT*I2:12000.00*I3:15000.00*I4:900.00*I5:50000.00*I6:6500.00*I7:80000.00*I8:18400.00*J1:PT-AC*J2:10000.00*J3:25000.56*J4:1000.02*J5:75000.00*J6:6750.00*J7:100000.00*J8:18000.00*K1:PT-MA*K2:5000.00*K3:12500.00*K4:625.00*K5:25000.00*K6:3000.00*K7:40000.00*K8:8800.00*L:100.00*M:25.00*N:64000.02*O:513600.58*P:100.00*Q:kLp0*R:9999*S:TB;PT00000000000000000000000;513500.58`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "123456789",
		ptinvoice.FieldCodeB:  "999999990",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "FT",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20191231",
		ptinvoice.FieldCodeG:  "FTAB2019/0035",
		ptinvoice.FieldCodeH:  "CSDF7T5H-0035",
		ptinvoice.FieldCodeI1: "PT",
		ptinvoice.FieldCodeI2: "12000.00",
		ptinvoice.FieldCodeI3: "15000.00",
		ptinvoice.FieldCodeI4: "900.00",
		ptinvoice.FieldCodeI5: "50000.00",
		ptinvoice.FieldCodeI6: "6500.00",
		ptinvoice.FieldCodeI7: "80000.00",
		ptinvoice.FieldCodeI8: "18400.00",
		ptinvoice.FieldCodeJ1: "PT-AC",
		ptinvoice.FieldCodeJ2: "10000.00",
		ptinvoice.FieldCodeJ3: "25000.56",
		ptinvoice.FieldCodeJ4: "1000.02",
		ptinvoice.FieldCodeJ5: "75000.00",
		ptinvoice.FieldCodeJ6: "6750.00",
		ptinvoice.FieldCodeJ7: "100000.00",
		ptinvoice.FieldCodeJ8: "18000.00",
		ptinvoice.FieldCodeK1: "PT-MA",
		ptinvoice.FieldCodeK2: "5000.00",
		ptinvoice.FieldCodeK3: "12500.00",
		ptinvoice.FieldCodeK4: "625.00",
		ptinvoice.FieldCodeK5: "25000.00",
		ptinvoice.FieldCodeK6: "3000.00",
		ptinvoice.FieldCodeK7: "40000.00",
		ptinvoice.FieldCodeK8: "8800.00",
		ptinvoice.FieldCodeL:  "100.00",
		ptinvoice.FieldCodeM:  "25.00",
		ptinvoice.FieldCodeN:  "64000.02",
		ptinvoice.FieldCodeO:  "513600.58",
		ptinvoice.FieldCodeP:  "100.00",
		ptinvoice.FieldCodeQ:  "kLp0",
		ptinvoice.FieldCodeR:  "9999",
		ptinvoice.FieldCodeS:  "TB;PT00000000000000000000000;513500.58",
	}

	parsedFields, err := ptinvoice.Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := parsedFields.Len(), expectedFields.Len(); got != want {
		t.Fatalf("len(%s) = %#v; want %#v", "parsedFields", got, want)
	}

	for key, _ := range expectedFields {
		if got, want := parsedFields.Get(key), expectedFields.Get(key); got != want {
			t.Errorf("parsedFields[%#v] = %#v; want %#v", key, got, want)
		}
	}
}

func TestParseMessage_SimplifiedInvoice(t *testing.T) {
	t.Skip("not implemented")
}

func TestParseMessage_ProFormaInvoice(t *testing.T) {
	t.Skip("not implemented")
}

func TestParseMessage_TransportDocument(t *testing.T) {
	t.Skip("not implemented")
}

func TestParseMessage_InvoiceWithForeignTaxRate(t *testing.T) {
	t.Skip("not implemented")
}

func TestParseMessage_TaxAmendingDebitNote(t *testing.T) {
	t.Skip("not implemented")
}

func TestParseMessage_InvoiceWithVatMarginRegime(t *testing.T) {
	t.Skip("not implemented")
}
