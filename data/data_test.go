package data_test

import (
	"fmt"
	"testing"

	"github.com/marstr/units/data"
)

func ExampleData() {
	capacity := 2 * data.Kilobyte
	fmt.Println(capacity)
	fmt.Printf("%d\n", capacity)
	// Output:
	// 2KB
	// 2048
}

func Example_sliceAllocation() {
	buf := make([]byte, 0, 10*data.Megabyte)
	fmt.Println(cap(buf))
	// Output:
	// 10485760
}

func TestData_String(t *testing.T) {
	testcases := []struct {
		Size     data.Data
		Expected string
	}{
		{6*data.Megabyte + 512*data.Kilobyte, "6.5MB"},
		{6 * data.Megabyte, "6MB"},
		{data.MaxData, "16EB"},
		{data.Exabyte + 512*data.Petabyte, "1.5EB"},
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
