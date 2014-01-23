package fsm

import (
	"testing"
)

func TestLiteral(t *testing.T) {
	nfa, err := ParseRegex("hello")
	if err != nil {
		t.Fatalf("Got error %s", err.Error())
	}
	resultNFA := NewNFA()
	a, b := resultNFA.NewState(), resultNFA.NewState()
	c, d := resultNFA.NewState(), resultNFA.NewState()
	e, f := resultNFA.NewState(), resultNFA.NewState()
	a.NewEdge(104, b)
	b.NewEdge(101, c)
	c.NewEdge(108, d)
	d.NewEdge(108, e)
	e.NewEdge(111, f)
	f.IsAccepting = true
	if !nfa.Equals(resultNFA) {
		t.Fatalf("NFA was not equal to constructed nfa")
	}
}

func TestStar(t *testing.T) {
	nfa, err := ParseRegex("b*")
	if err != nil {
		t.Fatalf("Got error %s", err.Error())
	}
	resultNFA := NewNFA()
	a, b := resultNFA.NewState(), resultNFA.NewState()
	c, d := resultNFA.NewState(), resultNFA.NewState()
	a.NewEdge(Epsilon, b)
	a.NewEdge(Epsilon, d)
	b.NewEdge(98, c)
	c.NewEdge(Epsilon, b)
	c.NewEdge(Epsilon, d)
	d.IsAccepting = true
	if !nfa.Equals(resultNFA) {
		t.Fatalf("NFA was not equal to constructed nfa")
	}
}

func TestUnion(t *testing.T) {
	nfa, err := ParseRegex("a|b")
	if err != nil {
		t.Fatalf("Got error %s", err.Error())
	}
	resultNFA := NewNFA()
	a, b := resultNFA.NewState(), resultNFA.NewState()
	c, d := resultNFA.NewState(), resultNFA.NewState()
	e, f := resultNFA.NewState(), resultNFA.NewState()
	a.NewEdge(Epsilon, b)
	b.NewEdge(Epsilon, d)
	b.NewEdge(97, c)
	d.NewEdge(98, e)
	c.NewEdge(Epsilon, f)
	e.NewEdge(Epsilon, f)
	f.IsAccepting = true
	if !nfa.Equals(resultNFA) {
		t.Fatalf("NFA was not equal to constructed nfa")
	}
}

func TestRange(t *testing.T) {
	nfa, err := ParseRegex("[0-9]+")
	if err != nil {
		t.Fatalf("Got error %s", err.Error())
	}
	resultNFA := NewNFA()
	a, b := resultNFA.NewState(), resultNFA.NewState()
	c, d := resultNFA.NewState(), resultNFA.NewState()
	a.NewEdge(Epsilon, b)
	for i := 48; i < 58; i++ {
		b.NewEdge(i, c)
	}
	c.NewEdge(Epsilon, d)
	c.NewEdge(Epsilon, b)
	d.IsAccepting = true
	if !nfa.Equals(resultNFA) {
		t.Fatalf("NFA was not equal to constructed nfa")
	}
}
