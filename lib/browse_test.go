package lib

import (
	"testing"
)

func TestOpenUrl(t *testing.T) {
	url := "https://www.baidu.com"
	err := Open(url)
	if err != nil {
		t.Fatal(err)
	}
}
