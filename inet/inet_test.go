package inet

import "testing"

func TestValidIPAddress(t *testing.T) {
	isValid := IsValidIPAddress("127.0.0.1")
	if isValid {
		t.Log("Is valid IP address")
		return
	}
	t.Error("Is invalid IP address")
}

func TestInvalidIPAddress(t *testing.T) {
	isValid := IsValidIPAddress("127.0.0.1234")
	if !isValid {
		t.Log("Is invalid IP address")
		return
	}
	t.Error("Is valid IP address")
}
