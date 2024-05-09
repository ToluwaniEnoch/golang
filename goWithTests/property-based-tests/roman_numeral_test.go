package main

import "testing"

func TestRomanNumeral(t *testing.T) {

	cases := []struct {
		Description string
		Number int
		Want string
	} {
		{Description: "1 gets converted to I", Number: 1, Want: "I"},
		{Description: "2 gets converted to II", Number: 2, Want: "II"},
		{Description: "3 gets converted to III", Number: 3, Want: "III"},
		{Description: "4 gets converted to IV", Number: 4, Want: "IV"},
		{Description: "5 gets converted to V", Number: 5, Want: "V"},
		{Description: "6 gets converted to VI", Number: 6, Want: "VI"},
		{Description: "7 gets converted to VII", Number: 7, Want: "VII"},
		{Description: "8 gets converted to VII", Number: 8, Want: "VIII"},
		{Description: "9 gets converted to IX", Number: 9, Want: "IX"},
		{Description: "10 gets converted to X", Number: 10, Want: "X"},
	}
	
	for _, test := range cases {
		t.Run(test.Description, func (t *testing.T)  {
			got := ConvertToRoman(test.Number)

			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}

}