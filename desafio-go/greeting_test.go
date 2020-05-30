package main

import "testing"

func TestGreeting(t *testing.T) {

	result := greeting("Code.education Rocks!")

	if result == "<b>Code.education Rocks!</b>" {
		t.Logf("greeting('Code.education Rocks!') sucesso")
	} else {
		t.Errorf("greeting('Code.education Rocks!') erro")
	}

}
