package xj2go

import (
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		nodes *[]leafNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max test",
			args: args{
				nodes: &[]leafNode{
					{
						path: "a.b.c.d.e.f.g.h",
					},
					{
						path: "a.b.c.d.e",
					},
					{
						path: "a.b.c",
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.nodes); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathExists(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    bool
		wantErr bool
	}{
		{
			name:    "not existed directory",
			path:    "./temp",
			want:    false,
			wantErr: false,
		},
		{
			name:    "existed directory",
			path:    "./testjson",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pathExists(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("pathExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toProperCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"toProperCase", args{"read_count"}, "ReadCount"},
		{"toProperCase", args{"read_id"}, "ReadID"},
		{"toProperCase", args{"readIdUrl"}, "ReadIDURL"},
		{"toProperCase", args{"readIdUrl_ip_xss"}, "ReadIDURLIPXSS"},
		{"toProperCase", args{"readIdUrl_ip_xssCpu"}, "ReadIDURLIPXssCPU"},
		{"toProperCase", args{"id"}, "ID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toProperCase(tt.args.str); got != tt.want {
				t.Errorf("toProperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toProperType(t *testing.T) {
	tests := []struct {
		name string
		val  interface{}
		want string
	}{
		{
			name: "string",
			val:  "this is a test",
			want: "string",
		},
		{
			name: "int",
			val:  1,
			want: "int",
		},
		{
			name: "time",
			val:  "2017-10-31T11:59:17+08:00",
			want: "time.Time",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toProperType(tt.val); got != tt.want {
				t.Errorf("toProperType() = %v, want %v", got, tt.want)
			}
		})
	}
}
