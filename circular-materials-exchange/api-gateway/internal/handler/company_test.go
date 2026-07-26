package handler

import (
	companypb "api-gateway/internal/pb/company"
	"reflect"
	"testing"
)

func TestCompanyJSONPreservesCertificationsArray(t *testing.T) {
	want := []string{"ISO 9001", "GRS"}
	payload := companyJSON(&companypb.Company{Certifications: want})

	got, ok := payload["certifications"].([]string)
	if !ok {
		t.Fatalf("certifications must be []string, got %T", payload["certifications"])
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("certifications = %#v, want %#v", got, want)
	}
}
