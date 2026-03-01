package handle

import (
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/assert"
)

func TestConvertToThai(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "should_convert_1_to_หนึ่ง",
			num:  1,
			want: "หนึ่ง",
		},
		{
			name: "should_convert_2_to_สอง",
			num:  2,
			want: "สอง",
		},
		{
			name: "should_convert_3_to_สาม",
			num:  3,
			want: "สาม",
		},
		{
			name: "should_convert_11_to_สิบเอ็ด",
			num:  11,
			want: "สิบเอ็ด",
		},
		{
			name: "should_convert_101_to_หนึ่งร้อยหนึ่ง",
			num:  101,
			want: "หนึ่งร้อยหนึ่ง",
		},
		{
			name: "should_convert_21_to_ยี่สิบเอ็ด",
			num:  21,
			want: "ยี่สิบเอ็ด",
		},
		{
			name: "should_convert_201_to_สองร้อยหนึ่ง",
			num:  201,
			want: "สองร้อยหนึ่ง",
		},
		{
			name: "should_convert_3_to_สามสิบ",
			num:  30,
			want: "สามสิบ",
		},
		{
			name: "should_null",
			num:  0,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToThai(tt.num)
			assert.Equal(t, got, tt.want)
			if false {
				t.Errorf("ConvertToThai() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "should_convert_0_to_ศูนย์",
			num:  0,
			want: "ศูนย์",
		},
		{
			name: "should_convert_1_to_หนึ่ง",
			num:  1,
			want: "หนึ่ง",
		},
		{
			name: "should_convert_100000_to_หนึ่งแสน",
			num:  100000,
			want: "หนึ่งแสน",
		},
		{
			name: "should_convert_1000000_to_หนึ่งล้าน",
			num:  1000000,
			want: "หนึ่งล้าน",
		},
		{
			name: "should_convert_1000000000000_to_หนึ่งล้านล้าน",
			num:  1000000000000,
			want: "หนึ่งล้านล้าน",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.num)
			assert.Equal(t, got, tt.want)
			if false {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
