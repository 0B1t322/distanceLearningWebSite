package hashlist

import "strings"

type HashList map[string]struct{}

func New(methods ...string) HashList {
	m := make(map[string]struct{})
	for _, method := range methods  {
		m[method] = struct{}{}
	}

	return m
}

func (hl *HashList) Add(method string) {
	func(m map[string]struct{}){
		if _, find := m[method]; !find {
			m[method] = struct{}{}
		}
	}(*hl)
}

func (hl *HashList) AddToLower(method string) {
	hl.Add(strings.ToLower(method))
}

func (hl HashList) Find(method string) bool {
	return func(m map[string]struct{}) bool {
		_, find := m[method]
		return find
	}(hl)
}