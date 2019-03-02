package command

// Query to run a command
// command to match the query and
// site to target and optional argument
type Query interface {
	target() string
	set([]string)
	get() map[string]string
}

// Site is a query
// for a command directed to sites
type Site struct {
	name string
	arg  string
	site string
}

func (s *Site) set(a []string) {
	s.name, s.arg, s.site = a[1], a[2], a[3]
}

func (s *Site) target() string {
	return s.name
}

func (s *Site) get() map[string]string {
	return map[string]string{
		"name": s.name,
		"arg":  s.arg,
		"site": s.site,
	}
}

// Help is a query
// for a help command
type Help struct {
	name string
	arg  string
}

func (s *Help) set(a []string) {
	s.name, s.arg = a[1], a[2]
}

func (s *Help) target() string {
	return s.name
}

func (s *Help) get() map[string]string {
	return map[string]string{
		"name": s.name,
		"arg":  s.arg,
	}
}
