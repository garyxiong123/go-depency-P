package P

import (
	M "github.com/garyxiong123/go-depency-C/C"

	"testing"
)

func Go_dependency_B_V1() {
	M.Go_dependency_C_V1()
	println("go_dependency_B_V1")
}

func TestDiffrentTag(t *testing.T) {
	Go_dependency_B_V1()
}
