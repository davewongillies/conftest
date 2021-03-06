package downloader

import "testing"

func TestOCIDetector_Detect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"should detect azurecr",
			"user.azurecr.io/policies:tag",
			"oci://user.azurecr.io/policies:tag",
		},
		{
			"should add latest tag",
			"user.azurecr.io/policies",
			"oci://user.azurecr.io/policies:latest",
		},
		{
			"should detect 127.0.0.1:5000 as most likely being an OCI registry",
			"127.0.0.1:5000/policies:tag",
			"oci://127.0.0.1:5000/policies:tag",
		},
		{
			"should detect 127.0.0.1:5000 as most likely being an OCI registry and tag it properly if no tag is supplied",
			"127.0.0.1:5000/policies",
			"oci://127.0.0.1:5000/policies:latest",
		},
	}
	pwd := "/pwd"
	d := &OCIDetector{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, ok, err := d.Detect(tt.input, pwd)
			if err != nil {
				t.Fatalf("OCIDetector.Detect() error = %v", err)
			}
			if !ok {
				t.Fatal("OCIDetector.Detect() not ok, should have detected")
			}
			if out != tt.expected {
				t.Errorf("OCIDetector.Detect() output = %v, want %v", out, tt.expected)
			}
		})
	}
}
