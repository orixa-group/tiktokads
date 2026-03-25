package tiktokads

import "testing"

func TestGetCatalogStatus(t *testing.T) {
	initTestSession()

	rows, err := GetCatalogStatus("7489850869144485895", "7615239103906449169")
	assertNotEmptyResult(t, rows, err)
	for _, row := range rows {
		if len(row.ProductId) == 0 {
			t.Error("missing product id")
		}
		if len(row.IssueTitle) == 0 {
			t.Error("missing issue title")
		}
		if len(row.IssueSeverity) == 0 {
			t.Error("missing issue severity")
		}
	}
}
