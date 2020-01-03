package kmp

import (
	"testing"
)

func TestViolent(t *testing.T) {
	arrS := "bacbabababcbab"
	arrP := "ababc"
	ViolentMatch(arrS, arrP)
}

func TestGetNext(t *testing.T) {
	arr := "ababa"
	getNexts(arr)
}
