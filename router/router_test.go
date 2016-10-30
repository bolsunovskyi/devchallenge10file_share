package router

import "testing"

func TestGetRouter(t *testing.T) {
	if router := GetRouter(); router == nil {
		t.Error("Unable to create router")
	}
}
