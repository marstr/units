package si_test

import (
	"fmt"
	"testing"

	"github.com/marstr/units/physical/imperial"
	"github.com/marstr/units/physical/si"
	"github.com/marstr/units/physical/us"
)

func TestVolume_Within(t *testing.T) {
	const tolerance = si.Millilitre / 1000000
	const fromLitre = 2.5 * si.Litre
	const fromMillilitre = 2500 * si.Millilitre

	if fromLitre == fromMillilitre {
		t.Logf("fromLiter and fromMilliliter were equal.")
	} else {
		t.Logf("fromLiter and fromMilliliter were different.")
	}

	close := fromMillilitre.Within(fromLitre, 1*si.Millilitre)
	if !close {
		t.Logf("%v and %v were NOT within %v of one another", fromLitre, fromMillilitre, tolerance)
		t.Fail()
	} else if testing.Verbose() {
		t.Logf("%v and %v were within %v of one another", fromLitre, fromMillilitre, tolerance)
	}
}

func ExampleVolume_imperialVsUS() {
	const americanBeer = us.Pint
	const britishBeer = imperial.Pint

	delta := britishBeer - americanBeer
	fmt.Printf("Brits get %v (%0.1f%%) more beer per pint than Americans.\n", delta, 100*delta/americanBeer)
	// Output:
	// Brits get 0.10L (20.1%) more beer per pint than Americans.
}
