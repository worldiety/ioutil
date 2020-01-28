package ioutil

/*
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

type OpenOptions struct {
}

func NewOpenOptions() *OpenOptions {
	return &OpenOptions{}
}

func (o *OpenOptions) Read(read bool) *OpenOptions {
	return o
}

func (o *OpenOptions) Write(write bool) *OpenOptions {
	return o
}

func (o *OpenOptions) Create(create bool) *OpenOptions {
	return o
}

func (o *OpenOptions) CreateNew(createNew bool) *OpenOptions {
	return o
}

func (o *OpenOptions) Truncate(truncate bool) *OpenOptions {
	return o
}

func (o *OpenOptions) Append(append bool) *OpenOptions {
	return o
}

func (o *OpenOptions) Open(name string) io.ReadWriteCloser {
	return nil
}

type FSQuery struct {
}

func NewQuery() *FSQuery {
	return nil
}

func (q *FSQuery) Select() *FSQuery {
	return q
}

func (q *FSQuery) Files() *FSQuery {
	return q
}

func (q *FSQuery) Where() *FSQuery {
	return q
}

func (q *FSQuery) File() *FSQuery {
	return q
}

func (q *FSQuery) Has() *FSQuery {
	return q
}

func (q *FSQuery) Name() *FSQuery {
	return q
}

func (q *FSQuery) Attribute() *FSQuery {
	return q
}

func (q *FSQuery) IsHidden() *FSQuery {
	return q
}

func (q *FSQuery) NotHidden() *FSQuery {
	return q
}

func (q *FSQuery) From(p string) *FSQuery {
	return q
}

func (q *FSQuery) Recursively() *FSQuery {
	return q
}

func (q *FSQuery) And() *FSQuery {
	return q
}

func (q *FSQuery) EndsWidth(s string) *FSQuery {
	return q
}

func (q *FSQuery) Execute() *FSQuery {
	return q
}*/
