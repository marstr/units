package data_test

import (
	"fmt"
	"github.com/marstr/units/data"
	"testing"
)

func ExampleData() {
	mySize := 2 * data.Kilobyte

	fmt.Println(mySize)
	fmt.Println(mySize / data.Byte)

	// Output:
	// 16384
	// 2048
}

func TestData_MarshalText(t *testing.T) {
	testcases := []struct {
		Size     data.Data
		Expected string
	}{
		{6*data.Megabyte + 512*data.Kilobyte, "6.5MB"},
		{6 * data.Megabyte, "6MB"},
		{4 * data.Bit, "0.5B"},
		{data.DataMax, "2EB"},
		{1.5 * data.Exabyte, "1.5EB"},
	}

	for _, tc := range testcases {
		t.Run(tc.Expected, func(t *testing.T) {
			if got := tc.Size.String(); got != tc.Expected {
				t.Logf("\n\tgot: \t%q\n\twant:\t%q", got, tc.Expected)
				t.Fail()
			}
		})
	}
}
