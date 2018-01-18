package main

func join(sep string, a ...string) string {
	s := ""
	for i, as := range a {
		s += as
		if i < len(a)-1 {
			s += sep
		}
	}
	return s
}
