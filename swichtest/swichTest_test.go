package swichtest

import "testing"

func TestSwitchTest(t *testing.T) {
	SwitchTest(60)
	SwitchTest(70)
	SwitchTest(10)
	SwitchTest(-1)
	SwitchTest(100)

	SwitchTests(10,20,"+")
	SwitchTests(10,20,"-")
	SwitchTests(10,20,"*")
	SwitchTests(10,20,"/")
}


