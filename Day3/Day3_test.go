package main

import (
	"testing"
)

func TestMulRegexp(t *testing.T) {
	res := mulRegex("mul(1,2)")
	act := ""
	if res != nil {
		act = res[0]
	}
	exp := "mul(1,2)"
	if act != exp {
		t.Error("Expected", exp, "but found", act)
	}
}

func TestMul(t *testing.T) {
	act := calcMul("mul(42,65)")
	exp := 2730
	if act != exp {
		t.Error("Expected", exp, "but found", act)
	}
}

func TestDoDontMulRegexp(t *testing.T) {
	act := mulDoDontRegex("mul(1,2)garbagedon't()somethingmul(5,6)againdo()mul(7,8)whatdon't()mul(9,10)butdo()thismul(11,12)")
	exp := [][]string{{"mul(1,2)garbagedon't()", "mul(1,2)garbage"},
		{"do()mul(7,8)whatdon't()", "mul(7,8)what"},
		{"do()thismul(11,12)", "thismul(11,12)"}}

	if len(act) != len(exp) {
		t.Error("Did not find the correct number of matches")
	}

	for i, match := range act {
		for j, actMatch := range match {
			if j > 0 && actMatch != "" {
				expMatch := exp[i][1]
				if actMatch != expMatch {
					t.Error("Expected match", expMatch, "but found", actMatch)
				}
			}

		}
	}
}

func TestMulFromRes(t *testing.T) {
	reRes := [][]string{{"mul(1,2)garbagedon't()", "mul(1,2)garbage"},
		{"do()mul(7,8)whatdon't()", "mul(7,8)what"},
		{"do()thismul(11,12)", "thismul(11,12)"}}

	act := calcMulFromRes(reRes)
	exp := 190

	if act != exp {
		t.Error("Expected", exp, "but got", act)
	}
}
