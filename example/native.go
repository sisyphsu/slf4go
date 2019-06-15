package main

import (
	"github.com/sisyphsu/slf4go"
	"github.com/sisyphsu/slf4go/example/modules"
)

// doesn't need initialize

// use slf4go everywhere
func main() {
	logger := slf4go.GetLogger("main")
	logger.Debugf("I want %s", "Cycle Import")
	logger.Errorf("please support it, in %02d second!", 1)
	modules.Login()
}