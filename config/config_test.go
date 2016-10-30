package config

import "testing"

func TestRead(t *testing.T) {
	res := Read("")
	if res {
		t.Error("Wrong behavoiur")
	}

	res = Read("../")
	if !res {
		t.Error("Wrong behavoiur")
	}
}
