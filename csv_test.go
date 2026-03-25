package tiktokads

import (
	"os"
	"testing"
)

func Test_convertCsv(t *testing.T) {
	r, err := os.Open("tests/catalog-report.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	rows, err := convertCsv[CatalogProductStatus](r)
	assertNotEmptyResult(t, rows, err)

}
