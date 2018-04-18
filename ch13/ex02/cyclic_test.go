package cyclic

import "testing"

func TestIsCyclicLinkedList(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	d, e, f := &link{value: "d"}, &link{value: "e"}, &link{value: "f"}
	a.tail, b.tail, c.tail = b, a, c
	d.tail, e.tail, f.tail = e, nil, nil
	for _, test := range []struct {
		name string
		link *link
		want bool
	}{
		{"normal cyclic", a, true},
		{"self cyclic", c, true},
		{"d->e", d, false},
		{"f", f, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := IsCyclic(test.link)
			if got != test.want {
				t.Errorf("error link want %t got %t\n", test.want, got)
			}
		})
	}
}
func TestIsCyclicSlice(t *testing.T) {
	type pointer struct {
		tail []*pointer
	}
	a := &pointer{}
	b := &pointer{}
	c := &pointer{}
	d := &pointer{}
	e := &pointer{}
	f := &pointer{}

	a.tail, b.tail, c.tail = append(a.tail, b), append(b.tail, a), append(c.tail, c)
	d.tail, e.tail, f.tail = append(d.tail, e), nil, nil
	for _, test := range []struct {
		name    string
		pointer *pointer
		want    bool
	}{
		{"normal cyclic", a, true},
		{"self cyclic", c, true},
		{"d->e", d, false},
		{"f", f, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := IsCyclic(test.pointer)
			if got != test.want {
				t.Errorf("error link want %t got %t\n", test.want, got)
			}
		})
	}
}
func TestIsCyclicArray(t *testing.T) {
	type pointer struct {
		tail [1]*pointer
	}
	a := &pointer{}
	b := &pointer{}
	c := &pointer{}
	d := &pointer{}
	e := &pointer{}
	f := &pointer{}

	a.tail[0], b.tail[0], c.tail[0] = b, a, c
	d.tail[0], e.tail[0], f.tail[0] = e, nil, nil
	for _, test := range []struct {
		name    string
		pointer *pointer
		want    bool
	}{
		{"normal cyclic", a, true},
		{"self cyclic", c, true},
		{"d->e", d, false},
		{"f", f, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := IsCyclic(test.pointer)
			if got != test.want {
				t.Errorf("error link want %t got %t\n", test.want, got)
			}
		})
	}
}
func TestIsCyclicMap(t *testing.T) {
	type pointer struct {
		tail map[string]*pointer
	}
	a := &pointer{make(map[string]*pointer)}
	b := &pointer{make(map[string]*pointer)}
	c := &pointer{make(map[string]*pointer)}
	d := &pointer{make(map[string]*pointer)}
	e := &pointer{make(map[string]*pointer)}
	f := &pointer{make(map[string]*pointer)}

	a.tail["a"], b.tail["b"], c.tail["c"] = b, a, c
	d.tail["d"], e.tail["e"], f.tail["f"] = e, nil, nil
	for _, test := range []struct {
		name    string
		pointer *pointer
		want    bool
	}{
		{"normal cyclic", a, true},
		{"self cyclic", c, true},
		{"d->e", d, false},
		{"f", f, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := IsCyclic(test.pointer)
			if got != test.want {
				t.Errorf("error link want %t got %t\n", test.want, got)
			}
		})
	}
}
