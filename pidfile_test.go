package pidfile

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

var (
	badPath = "bad/path/test.pid"
	okPath  = "test.pid"
)

func TestNew(t *testing.T) {
	_, err := New(badPath)
	if err == nil {
		t.Fail()
	}

	p, err := New(okPath)
	if err != nil {
		t.Fail()
	}

	defer p.Remove()

	b, err := ioutil.ReadFile(p.Path)
	if err != nil {
		t.Fail()
	}

	i, err := strconv.Atoi(string(b))
	if err != nil {
		t.Fail()
	}

	if i != os.Getpid() {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	p, err := New(okPath)
	if err != nil {
		t.Fail()
	}

	err = p.Remove()
	if err != nil {
		t.Fail()
	}

	_, err = os.Stat(p.Path)
	if os.IsExist(err) {
		t.Fail()
	}
}
