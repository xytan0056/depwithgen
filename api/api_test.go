package api

import (
	"testing"
	c "github.com/xytan0056/depwithgen/.gen/client"
)

func TestSomeTest(t *testing.T) {
	t.Logf("testing on %s", c.Client("007"))
}