package main

import "testing"

type testCases struct {
	A, B, Expected int
}

//只要傳入測試用例即可測試
func createTestCases(t *testing.T, cases *testCases) {
	t.Helper() //marks the calling function as a test helper function
	if ans := Mul(cases.A, cases.B); ans != cases.Expected {
		t.Fatalf("%d * %d expected %d, but %d got.", cases.A, cases.B, cases.Expected, ans)
	}
}

func TestMul(t *testing.T) {
	createTestCases(t, &testCases{1, 8, 8})
	createTestCases(t, &testCases{-2, -2, 4})
	createTestCases(t, &testCases{A: 2, B: 0, Expected: 0})
}
