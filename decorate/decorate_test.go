// file: decorate_test.go
// author: TaoZer
// Date: 2022/1/26
// Description: decoration pattern test

package decorate

import (
	"fmt"
	"testing"
)

func TestDecorate(t *testing.T) {
	mt := NewMilkTea()
	bb := AddBubble(mt)
	bbs := AddBubble(bb)
	fmt.Println(bbs.Description())
}