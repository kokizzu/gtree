package gtree

import (
	"testing"
)

type want struct {
	Name      string
	hierarchy uint
	index     uint
}

const fixedIndex uint = 1

func TestTabStrategy_Generate(t *testing.T) {
	// ref: https://zenn.dev/kimuson13/articles/go_table_driven_test
	tests := map[string]struct {
		row  string
		want want
	}{
		"root/hierarchy=1": {"- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		"child/hierarchy=2": {"	- aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
		"child/hierarchy=2/tab on the way": {"	- aaa	bb", want{Name: "aaa	bb", hierarchy: 2, index: fixedIndex}},
		"invalid/hierarchy=0/prefix space": {" - aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/prefix chars": {"xx- aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/no hyphen":    {"xx aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/tab only": {"			", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node := (*tabStrategy)(nil).generate(tt.row, fixedIndex)
			if node.name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}

func TestTwoSpacesStrategy_Generate(t *testing.T) {
	tests := map[string]struct {
		row  string
		want want
	}{
		"root/hierarchy=1":                     {"- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		"child/hierarchy=2":                    {"  - aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/space only":       {"  ", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node := (*twoSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if node.name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}

func TestFourSpacesStrategy_Generate(t *testing.T) {
	tests := map[string]struct {
		row  string
		want want
	}{
		"root/hierarchy=1":                     {"- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		"child/hierarchy=2":                    {"    - aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
		"child/hierarchy=3":                    {"        - aaa    bb", want{Name: "aaa    bb", hierarchy: 3, index: fixedIndex}},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
		"invalid/hierarchy=0/space only":       {"    ", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node := (*fourSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if node.name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}