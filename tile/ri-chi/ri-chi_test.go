package ri_chi

import (
	"fmt"
	"testing"
)

func TestRichi1(t *testing.T) {
	richi := Richi{}
	array := richi.ReturnTilesSelfArray()
	fmt.Println()
	for _, i2 := range array{
		i2.PrintCard()
	}
}
