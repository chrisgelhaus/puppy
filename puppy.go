package puppy

import (
	"github.com/chrisgelhaus/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Version() string {
	return "Version 1.0.0"
}

func Version12() string {
	return "Version 1.2.0"
}
