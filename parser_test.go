package ptinvoice_test

import (
	"testing"

	"github.com/c032/go-ptinvoice"
)

func TestParseMessage_Invoice(t *testing.T) {
	const input = `A:123456789*B:999999990*C:PT*D:FT*E:N*F:20191231*G:FT AB2019/0035*H:CSDF7T5H-0035*I1:PT*I2:12000.00*I3:15000.00*I4:900.00*I5:50000.00*I6:6500.00*I7:80000.00*I8:18400.00*J1:PT-AC*J2:10000.00*J3:25000.56*J4:1000.02*J5:75000.00*J6:6750.00*J7:100000.00*J8:18000.00*K1:PT-MA*K2:5000.00*K3:12500.00*K4:625.00*K5:25000.00*K6:3000.00*K7:40000.00*K8:8800.00*L:100.00*M:25.00*N:64000.02*O:513600.58*P:100.00*Q:kLp0*R:9999*S:TB;PT00000000000000000000000;513500.58`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "123456789",
		ptinvoice.FieldCodeB:  "999999990",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "FT",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20191231",
		ptinvoice.FieldCodeG:  "FT AB2019/0035",
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
	const input = `A:123456789*B:999999990*C:PT*D:FS*E:N*F:20190812*G:FS CDVF/12345*H:CDF7T5HD-12345*I1:PT*I7:0.65*I8:0.15*N:0.15*O:0.80*Q:YhGV*R:9999*S:NU;0.80`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "123456789",
		ptinvoice.FieldCodeB:  "999999990",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "FS",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20190812",
		ptinvoice.FieldCodeG:  "FS CDVF/12345",
		ptinvoice.FieldCodeH:  "CDF7T5HD-12345",
		ptinvoice.FieldCodeI1: "PT",
		ptinvoice.FieldCodeI7: "0.65",
		ptinvoice.FieldCodeI8: "0.15",
		ptinvoice.FieldCodeN:  "0.15",
		ptinvoice.FieldCodeO:  "0.80",
		ptinvoice.FieldCodeQ:  "YhGV",
		ptinvoice.FieldCodeR:  "9999",
		ptinvoice.FieldCodeS:  "NU;0.80",
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

func TestParseMessage_ProFormaInvoice(t *testing.T) {
	const input = `A:500000000*B:123456789*C:PT*D:PF*E:N*F:20190123*G:PF G2019CB/145789*H:HB6FT7RV-145789*I1:PT*I2:12345.34*I3:12532.65*I4:751.96*I5:52789.00*I6:6862.57*I7:32425.69*I8:7457.91*N:15072.44*O:125165.12*Q:r/fY*R:9999`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "500000000",
		ptinvoice.FieldCodeB:  "123456789",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "PF",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20190123",
		ptinvoice.FieldCodeG:  "PF G2019CB/145789",
		ptinvoice.FieldCodeH:  "HB6FT7RV-145789",
		ptinvoice.FieldCodeI1: "PT",
		ptinvoice.FieldCodeI2: "12345.34",
		ptinvoice.FieldCodeI3: "12532.65",
		ptinvoice.FieldCodeI4: "751.96",
		ptinvoice.FieldCodeI5: "52789.00",
		ptinvoice.FieldCodeI6: "6862.57",
		ptinvoice.FieldCodeI7: "32425.69",
		ptinvoice.FieldCodeI8: "7457.91",
		ptinvoice.FieldCodeN:  "15072.44",
		ptinvoice.FieldCodeO:  "125165.12",
		ptinvoice.FieldCodeQ:  "r/fY",
		ptinvoice.FieldCodeR:  "9999",
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

func TestParseMessage_TransportDocument(t *testing.T) {
	const input = `A:500000000*B:123456789*C:PT*D:GT*E:N*F:20190720*G:GT G234CB/50987*H:GTVX4Y8B-50987*I1:0*N:0.00*O:0.00*Q:5uIg*R:9999`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "500000000",
		ptinvoice.FieldCodeB:  "123456789",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "GT",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20190720",
		ptinvoice.FieldCodeG:  "GT G234CB/50987",
		ptinvoice.FieldCodeH:  "GTVX4Y8B-50987",
		ptinvoice.FieldCodeI1: "0",
		ptinvoice.FieldCodeN:  "0.00",
		ptinvoice.FieldCodeO:  "0.00",
		ptinvoice.FieldCodeQ:  "5uIg",
		ptinvoice.FieldCodeR:  "9999",
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

func TestParseMessage_InvoiceWithForeignTaxRate(t *testing.T) {
	const input = `A:123456789*B:4443332215*C:FR*D:FT*E:N*F:20190526*G:ABC BNH/4561*H:DK5ZJ2HN-4561*I1:FR*I7:100.00*I8:20.00*N:20.00*O:120.00*Q:YJRE*R:9999`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "123456789",
		ptinvoice.FieldCodeB:  "4443332215",
		ptinvoice.FieldCodeC:  "FR",
		ptinvoice.FieldCodeD:  "FT",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20190526",
		ptinvoice.FieldCodeG:  "ABC BNH/4561",
		ptinvoice.FieldCodeH:  "DK5ZJ2HN-4561",
		ptinvoice.FieldCodeI1: "FR",
		ptinvoice.FieldCodeI7: "100.00",
		ptinvoice.FieldCodeI8: "20.00",
		ptinvoice.FieldCodeN:  "20.00",
		ptinvoice.FieldCodeO:  "120.00",
		ptinvoice.FieldCodeQ:  "YJRE",
		ptinvoice.FieldCodeR:  "9999",
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

func TestParseMessage_TaxCorrectionDebitNote(t *testing.T) {
	const input = `A:123456789*B:500000000*C:PT*D:ND*E:N*F:20190216*G: M1F KLG/6145*H: RQD8L6DG-6145*I1:PT-MA*I6:26.50*N:26.50*O:26.50*Q:h1rB*R:9999`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "123456789",
		ptinvoice.FieldCodeB:  "500000000",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "ND",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20190216",
		ptinvoice.FieldCodeG:  " M1F KLG/6145",
		ptinvoice.FieldCodeH:  " RQD8L6DG-6145",
		ptinvoice.FieldCodeI1: "PT-MA",
		ptinvoice.FieldCodeI6: "26.50",
		ptinvoice.FieldCodeN:  "26.50",
		ptinvoice.FieldCodeO:  "26.50",
		ptinvoice.FieldCodeQ:  "h1rB",
		ptinvoice.FieldCodeR:  "9999",
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

func TestParseMessage_InvoiceIssuedUnderVATMarginScheme(t *testing.T) {
	const input = `A:500000000*B:123456789*C:PT*D:FT*E:N*F:20191124*G:NF 19A/789145*H: JL9DS4TT-789145*I1:PT-AC*I7:50.00*I8:9.00*L:1000.00*N:9.00*O:1059.00*Q:d8/K*R:9999`

	expectedFields := ptinvoice.Fields{
		ptinvoice.FieldCodeA:  "500000000",
		ptinvoice.FieldCodeB:  "123456789",
		ptinvoice.FieldCodeC:  "PT",
		ptinvoice.FieldCodeD:  "FT",
		ptinvoice.FieldCodeE:  "N",
		ptinvoice.FieldCodeF:  "20191124",
		ptinvoice.FieldCodeG:  "NF 19A/789145",
		ptinvoice.FieldCodeH:  " JL9DS4TT-789145",
		ptinvoice.FieldCodeI1: "PT-AC",
		ptinvoice.FieldCodeI7: "50.00",
		ptinvoice.FieldCodeI8: "9.00",
		ptinvoice.FieldCodeL:  "1000.00",
		ptinvoice.FieldCodeN:  "9.00",
		ptinvoice.FieldCodeO:  "1059.00",
		ptinvoice.FieldCodeQ:  "d8/K",
		ptinvoice.FieldCodeR:  "9999",
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
