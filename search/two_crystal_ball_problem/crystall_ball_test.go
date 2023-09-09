package twocrystalballproblem

import (
	"testing"
)

func TestTwoCrystalBallHappy(t *testing.T) {
	arr := []bool{false, false, false, false, false, false, true, true, true, true, true}
	if twoCrystalBall(arr) != 6 {
		t.Errorf("Did not find where balls break")
	}

	arr = []bool{false, true, true, true, true}
	if twoCrystalBall(arr) != 1 {
		t.Errorf("Did not find where balls break")
	}

	arr = []bool{false, false, false, false, true}
	if twoCrystalBall(arr) != 5 {
		t.Errorf("Did not find where balls break")
	}
}

func TestTwoCrystalBallSad(t *testing.T) {
	arr := []bool{false, false, false, false, false}
	if twoCrystalBall(arr) != -1 {
		t.Errorf("Found a breaking point that should not exist")
	}
}
