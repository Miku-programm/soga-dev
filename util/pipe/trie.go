package pipe

import "strings"

// YhrM38Rda retains the 24-byte original RTTI layout.
type YhrM38Rda struct {
	xMLHHG1nh   map[string]*YhrM38Rda
	HJKQN58GQYL any
}

// O45rZY8aJQtn has one root pointer, read by Search at offset zero. Its original
// field name was not retained; root names that observed field in restored source.
type O45rZY8aJQtn struct{ root *YhrM38Rda }

// NewEmptyTrie expresses the original static root plus init-time makemap used
// by LpKMLQuaAsk (0x5029418 -> 0x5029420 -> 0x5099b68).
// This helper name is new; it does not claim an additional original symbol.
func NewEmptyTrie() *O45rZY8aJQtn {
	return &O45rZY8aJQtn{root: &YhrM38Rda{xMLHHG1nh: make(map[string]*YhrM38Rda)}}
}

// Pbem1Ia: 0x191be60. A leading empty label is allowed here, but Search rejects it.
func Pbem1Ia(domain string) ([]string, bool) {
	if strings.HasSuffix(domain, ".") {
		return nil, false
	}
	labels := strings.Split(domain, ".")
	if len(labels) == 1 {
		if labels[0] == "" {
			return nil, false
		}
		return labels, true
	}
	for _, label := range labels[1:] {
		if label == "" {
			return nil, false
		}
	}
	return labels, true
}

// Search: 0x191bf60. Names are case-sensitive and trailing dots are invalid.
func (t *O45rZY8aJQtn) Search(domain string) *YhrM38Rda {
	labels, ok := Pbem1Ia(domain)
	if !ok || labels[0] == "" {
		return nil
	}
	node := t.a0_RMMKxLWS(t.root, labels)
	if node == nil || node.HJKQN58GQYL == nil {
		return nil
	}
	return node
}

// a0_RMMKxLWS: 0x191c020. Exact labels outrank one-label '*', then an empty
// label catches the remaining suffix. Branches without data permit fallback.
func (t *O45rZY8aJQtn) a0_RMMKxLWS(root *YhrM38Rda, labels []string) *YhrM38Rda {
	if len(labels) == 0 {
		return root
	}
	tail := labels[:len(labels)-1]
	if node := root.xMLHHG1nh[labels[len(labels)-1]]; node != nil {
		if result := t.a0_RMMKxLWS(node, tail); result != nil && result.HJKQN58GQYL != nil {
			return result
		}
	}
	if node := root.xMLHHG1nh["*"]; node != nil {
		if result := t.a0_RMMKxLWS(node, tail); result != nil && result.HJKQN58GQYL != nil {
			return result
		}
	}
	return root.xMLHHG1nh[""]
}
