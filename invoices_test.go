package invoicexpress

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestInvoiceTaxUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		wantID          int
		wantName        string
		wantValue       float64
		wantErrContains string
	}{
		{
			name:      "value as integer number",
			input:     `{"id":1,"name":"VAT","value":23}`,
			wantID:    1,
			wantName:  "VAT",
			wantValue: 23,
		},
		{
			name:      "value as decimal number",
			input:     `{"id":2,"name":"Reduced","value":6.5}`,
			wantID:    2,
			wantName:  "Reduced",
			wantValue: 6.5,
		},
		{
			name:      "value as numeric string",
			input:     `{"id":3,"name":"VAT","value":"23"}`,
			wantID:    3,
			wantName:  "VAT",
			wantValue: 23,
		},
		{
			name:      "value as trimmed numeric string",
			input:     `{"id":4,"name":"VAT","value":" 17.25 "}`,
			wantID:    4,
			wantName:  "VAT",
			wantValue: 17.25,
		},
		{
			name:      "value as null becomes zero",
			input:     `{"id":5,"name":"VAT","value":null}`,
			wantID:    5,
			wantName:  "VAT",
			wantValue: 0,
		},
		{
			name:      "value as empty string becomes zero",
			input:     `{"id":6,"name":"VAT","value":""}`,
			wantID:    6,
			wantName:  "VAT",
			wantValue: 0,
		},
		{
			name:            "invalid numeric string returns error",
			input:           `{"id":7,"name":"VAT","value":"abc"}`,
			wantErrContains: "invalid tax value",
		},
		{
			name:            "invalid type returns error",
			input:           `{"id":8,"name":"VAT","value":true}`,
			wantErrContains: "invalid tax value type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got InvoiceTax
			err := json.Unmarshal([]byte(tt.input), &got)

			if tt.wantErrContains != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErrContains)
				}
				if !strings.Contains(err.Error(), tt.wantErrContains) {
					t.Fatalf("expected error containing %q, got %q", tt.wantErrContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.ID != tt.wantID {
				t.Fatalf("ID: got %d, want %d", got.ID, tt.wantID)
			}
			if got.Name != tt.wantName {
				t.Fatalf("Name: got %q, want %q", got.Name, tt.wantName)
			}
			if got.Value != tt.wantValue {
				t.Fatalf("Value: got %v, want %v", got.Value, tt.wantValue)
			}
		})
	}
}
