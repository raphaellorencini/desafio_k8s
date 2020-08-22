package main

import "testing"

func TestGreetingsFunc(t *testing.T) {
	result := greetings("Code.education Rocks!")
	expected := "<b>Code.education Rocks!</b>";
    if result != "<b>Code.education Rocks!</b>" {
       t.Errorf("Func greetings() error: expected %s, got: %s.", expected, result)
    }
}