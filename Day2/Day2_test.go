package main

import "testing"

func TestGoodIncr(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 2 5 6 7"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestGoodDecr(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 9 8 5 2"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestIncrEq(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 2 3 3 4 5"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestDecrEq(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 9 9 8 7 6 5"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestGoodIncrDecr(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 4 6 5 7 8"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestGoodDecrIncr(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 7 8 6 5 4"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestGoodIncrDoubleSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 2 1 3 4 5"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestGoodDecrDoubleSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"2 1 2"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestBadIncrDoubleEq(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 2 2 3 4 5 5"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestBadDecrDoubleEq(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 9 8 8 7 6 6 5"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestBadIncrDoubleSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 4 6 5 8 7 9 10"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}

func TestBadDecrDoubleSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 7 8 5 2 3 1"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}

func TestBadDecrDoubleSwitch2(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 7 4 6 5 2 3"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}

func TestBadMultiSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"2 1 2 1"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}

func TestBadWithRecurse(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 2 3 4 6 4 7 8"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestBadIncrGoodSwitch(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"1 3 7 6 8 9"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}

func TestSampleData1(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"9 8 10 11 14 14 17 18"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}

func TestSampleData2(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"2 4 4 6 9 10 13 11"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestSampleData3(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"10 8 10 13 15 17 14 14"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestSampleData4(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"85 82 86 88 86"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestSampleData5(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"85 89 91 91 93 95 96"})
	exp := 0

	if act != exp {
		t.Error("Expected unsafe report")
	}
}
func TestSampleData6(t *testing.T) {
	act := isSafeWithOneRemoval([]string{"23 26 25 23 20"})
	exp := 1

	if act != exp {
		t.Error("Expected safe report")
	}
}
