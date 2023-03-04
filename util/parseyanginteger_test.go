package util_test

import (
	"github.com/openconfig/ygot/util"
	"testing"
)

func TestParseYangUint(t *testing.T) {
	tests := []struct {
		desc    string
		expr    string
		bits    int
		wantVal uint64
		wantErr bool
	}{
		{
			desc:    "100",
			expr:    "100",
			bits:    32,
			wantVal: 100,
			wantErr: false,
		},
		{
			desc:    "+100",
			expr:    "+100",
			bits:    32,
			wantVal: 100,
			wantErr: false,
		},
		{
			desc:    "-100",
			expr:    "-100",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "0xFF",
			expr:    "0xFF",
			bits:    32,
			wantVal: 255,
			wantErr: false,
		},
		{
			desc:    "+0xFF",
			expr:    "+0xFF",
			bits:    32,
			wantVal: 255,
			wantErr: false,
		},
		{
			desc:    "-0xFF",
			expr:    "-0xFF",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "+ by itself",
			expr:    "+",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "empty string",
			expr:    "",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "foobar",
			expr:    "foobar",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotVal, gotErr := util.ParseYangUint(tt.expr, tt.bits)
			if gotVal != tt.wantVal {
				t.Errorf("got value: %d, wanted: %d\n", gotVal, tt.wantVal)
			}
			if tt.wantErr && gotErr == nil {
				t.Errorf("failed to return error where error is expected\n")
			}
			if !tt.wantErr && gotErr != nil {
				t.Errorf("unexpected error: %q", gotErr.Error())
			}
		})
	}
}

func TestParseYangInt(t *testing.T) {
	tests := []struct {
		desc    string
		expr    string
		bits    int
		wantVal int64
		wantErr bool
	}{
		{
			desc:    "100",
			expr:    "100",
			bits:    32,
			wantVal: 100,
			wantErr: false,
		},
		{
			desc:    "+100",
			expr:    "+100",
			bits:    32,
			wantVal: 100,
			wantErr: false,
		},
		{
			desc:    "-100",
			expr:    "-100",
			bits:    32,
			wantVal: -100,
			wantErr: false,
		},
		{
			desc:    "0xFF",
			expr:    "0xFF",
			bits:    32,
			wantVal: 255,
			wantErr: false,
		},
		{
			desc:    "+0xFF",
			expr:    "+0xFF",
			bits:    32,
			wantVal: 255,
			wantErr: false,
		},
		{
			desc:    "-0xFF",
			expr:    "-0xFF",
			bits:    32,
			wantVal: -255,
			wantErr: false,
		},
		{
			desc:    "+ by itself",
			expr:    "+",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "empty string",
			expr:    "",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
		{
			desc:    "foobar",
			expr:    "foobar",
			bits:    32,
			wantVal: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotVal, gotErr := util.ParseYangInt(tt.expr, tt.bits)
			if gotVal != tt.wantVal {
				t.Errorf("got value: %d, wanted: %d\n", gotVal, tt.wantVal)
			}
			if tt.wantErr && gotErr == nil {
				t.Errorf("failed to return error where error is expected\n")
			}
			if !tt.wantErr && gotErr != nil {
				t.Errorf("unexpected error: %q", gotErr.Error())
			}
		})
	}
}
