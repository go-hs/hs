package deck

import (
	"fmt"
	"testing"
)

func TestError_Wrap(t *testing.T) {
	var custom ErrorMessage = "Test Error"
	err := custom.Wrap(fmt.Errorf("causer"))
	t.Log(err.Error())
}
