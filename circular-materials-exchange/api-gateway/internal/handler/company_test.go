package handler

import (
	companypb "api-gateway/internal/pb/company"
	"encoding/json"
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

func TestCompanyJSONUsesEmptyArrayForMissingCertifications(t *testing.T) {
	payload := companyJSON(&companypb.Company{})

	got, ok := payload["certifications"].([]string)
	if !ok {
		t.Fatalf("certifications must be []string, got %T", payload["certifications"])
	}
	if len(got) != 0 {
		t.Fatalf("certifications = %#v, want empty array", got)
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if _, ok := decoded["certifications"].([]interface{}); !ok {
		t.Fatalf("certifications JSON must be an array, got %s", encoded)
	}
}
