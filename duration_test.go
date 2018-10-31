package duration

import (
	"testing"
	"time"
)

func Test_FormatFuzzy(t *testing.T) {
	t.Log(FormatFuzzy(time.Hour * 10000))
}
