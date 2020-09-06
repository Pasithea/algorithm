package main

import (
	"container/list"
	"errors"
	"fmt"
)

func main() {
	// test data
	x := "I dont't need sex the government fucks me everyday."
	patterns := []string{"fuck", "fucks"}
	m := NewMatcher()
	m.Build(patterns)
	fmt.Println(m.Match(x))
}

type trieNode struct {
	child    [256]*trieNode
	fail     *trieNode
	length   uint32
	isEnding bool
}

func newTrieNode(length uint32) *trieNode {
	return &trieNode{[256]*trieNode{}, nil, length, false}
}

// Matcher matcher
type Matcher struct {
	root *trieNode
	size uint32
}

// NewMatcher create new matcher
func NewMatcher() *Matcher {
	return &Matcher{newTrieNode(0), 0}
}

// Build build
func (m *Matcher) Build(patterns []string) {
	for _, s := range patterns {
		m.addPattern(s)
	}
	m.build()
}

func (m *Matcher) build() {
	l := list.New()
	l.PushBack(m.root)
	for l.Len() > 0 {
		if tn, ok := l.Remove(l.Front()).(*trieNode); ok {
			var p *trieNode
			for i := range tn.child {
				if tn.child[i] == nil {
					continue
				}
				if tn == m.root {
					tn.child[i].fail = m.root
				} else {
					p = tn.fail
					for p != nil {
						if p.child[i] != nil {
							tn.child[i].fail = p.child[i]
							break
						}
						p = p.fail
					}
					if p == nil {
						tn.child[i].fail = m.root
					}
				}
				l.PushBack(tn.child[i])
			}
		} else {
			panic("build matcher error.")
		}
	}
}

func (m *Matcher) addPattern(pattern string) {
	curNode := m.root
	for i, v := range []byte(pattern) {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode(uint32(i + 1))
		}
		curNode = curNode.child[v]
	}
	curNode.isEnding = true
	m.size++
}

// Match match
func (m *Matcher) Match(text string) (string, error) {
	curNode := m.root
	var p *trieNode
	for i, v := range []byte(text) {
		for curNode.child[v] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = m.root
		}
		p = curNode
		for p != m.root {
			if p.isEnding {
				start := i - int(p.length) + 1
				end := start + int(p.length)
				return text[start:end], errors.New("containing prohibited words")
			}
			p = p.fail
		}
	}
	return "", nil
}
