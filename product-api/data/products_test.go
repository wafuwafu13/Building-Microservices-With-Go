package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "wafu",
		Price: 1.00,
		SKU: "avs-urrw-eef",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}