package xj2go

import "testing"

func Test_toProperCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{"toProperCase", args{"read_count"}},
		{"toProperCase", args{"read_id"}},
		{"toProperCase", args{"readIdUrl"}},
		{"toProperCase", args{"readIdUrl_ip_xss"}},
		{"toProperCase", args{"readIdUrl_ip_xssCpu"}},
		{"toProperCase", args{"id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toProperCase(tt.args.str)
		})
	}
}
