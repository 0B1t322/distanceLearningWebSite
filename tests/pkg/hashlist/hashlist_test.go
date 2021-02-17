package hashlist_test

import (
	"strings"
	"testing"

	"github.com/0B1t322/distanceLearningWebSite/pkg/hashlist"
)

func TestFunc_Find(t *testing.T) {
	h := hashlist.New("addcourse")

	method := "/coursesservice/addcourse"
	i := strings.LastIndex(method,"/")
	t.Log(method[i+1:])
	find := h.Find(method[i+1:])
	if !find {
		t.FailNow()
	}
}