package classes

import (
	"github.com/dkadio/dysr/util"
	"log"
	"testing"
)

func TestStore(t *testing.T) {
	c, _ := util.LoadConfig("./util")
	s := NewBoltAdapter(c)
	s.putValueFor("foo", "bar")
	v := s.getValueFor("foo")
	if v != "bar" {
		t.Fail()
	}
	log.Println("key:", "foo --> ", "bar ==", v)
}
