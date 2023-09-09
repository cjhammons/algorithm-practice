package twocrystalballproblem

import (
	"testing"
)

func TestTwoCrystalBallHappy(t *testing.T) {
	arr := []bool{false, false, false, false, false, false, true, true, true, true, true}
	if twoCrystalBall(arr) != 6 {
		t.Errorf("Did not find where balls break")
	}
}
