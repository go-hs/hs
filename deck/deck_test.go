package deck

import "testing"

func TestDecode_ValidString(t *testing.T) {
	d, err := DecodeString("AAECAR8GxwPJBLsFmQfZB/gIDI0B2AGoArUDhwSSBe0G6wfbCe0JgQr+DAA=")
	if err != nil {
		t.Error(err)
	}

	t.Log(d)
}
