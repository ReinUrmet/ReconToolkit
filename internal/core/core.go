package core

type Host struct {
	Name     string
	IPs      []string
	Status   string
	Title    string
	Tech     []string
	Tags     []string
	Sources  []string
	Ports    []Port
	Findings []Finding
}

// Finding is a security observation about a host.
type Finding struct {
	Title    string
	Severity string
	Evidence string
	Ref      string
}

// Port is one open port on a host.
type Port struct {
	Number  int
	State   string
	Service string
}
