package dag

type Vertex interface {
	ID() string
	GetSourcesLength() int
	GetTargetsLength() int
	GetSources() []Vertex
	GetTargets() []Vertex
	AddSource(p Vertex) error
	AddTarget(c Vertex) error
	RemoveSource(id string) error
	RemoveTarget(id string) error
}
