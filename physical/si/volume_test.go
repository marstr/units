package si

import (
	"testing"
)

func TestVolume(t *testing.T) {
	const tolerance = Millilitre / 1000000
	const fromLitre = 2.5 * Litre
	const fromMillilitre = 2500 * Millilitre

	if fromLitre == fromMillilitre {
		t.Logf("fromLiter and fromMilliliter were equal.")
	} else {
		t.Logf("fromLiter and fromMilliliter were different.")
	}

	close := fromMillilitre.Within(fromLitre, 1*Millilitre)
	if !close {
		t.Logf("%v and %v were NOT within %v of one another", fromLitre, fromMillilitre, tolerance)
		t.Fail()
	} else if testing.Verbose() {
		t.Logf("%v and %v were within %v of one another", fromLitre, fromMillilitre, tolerance)
	}
}
