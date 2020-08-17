package dag

type Vertex interface {
	ID() string
	GetSource(id string) Vertex
	GetTarget(id string) Vertex
	Sources() []Vertex
	Targets() []Vertex
	AddSource(p Vertex) error
	AddTarget(c Vertex) error
	RemoveSource(id string) error
	RemoveTarget(id string) error
}
