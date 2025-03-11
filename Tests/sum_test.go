package Tests

import "testing"


func testSum(t *testing.T) {
	// To fail without a message
	t.Fail()

	// To fail with a message
	t.Fatalf("failed")

	
}