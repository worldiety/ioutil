package ioutil

import "path/filepath"

// A Path is a slice of names.
type Path []string

// Parent returns the parent path segments. The parent of the empty element is nil.
func (p Path) Parent() Path {
	if len(p) == 0 {
		return nil
	}

	return p[0 : len(p)-1]
}

// IsEmpty returns true, if no path segments are present
func (p Path) IsEmpty() bool {
	return len(p) == 0
}

// Name returns the name of the last segment or the empty string.
func (p Path) Name() string {
	if p.IsEmpty() {
		return ""
	}

	return p[len(p)-1]
}

// String returns this path as a joined platform specific path.
func (p Path) String() string {
	return filepath.Join(p...)
}
