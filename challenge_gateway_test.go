package challengelib

import (
	"fmt"
	"testing"
)

func Test_getChallenge(t *testing.T) {
	s := FindChallenge("easy", "backend")

	fmt.Println(s)
}
