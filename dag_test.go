package dag

import (
	"testing"
)

type mockVertex struct {
	id      string
	sources []Vertex
	targets []Vertex
	val     interface{}
}

func newMockVertex(id string, val interface{}) *mockVertex {
	return &mockVertex{
		id:  id,
		val: val,
	}
}

func (m mockVertex) ID() string {
	return m.id
}

func (m mockVertex) GetSourcesLength() int {
	return len(m.sources)
}

func (m mockVertex) GetTargetsLength() int {
	return len(m.targets)
}

func (m mockVertex) GetSources() []Vertex {
	return m.sources
}

func (m mockVertex) GetTargets() []Vertex {
	return m.targets
}

func (m *mockVertex) AddSource(p Vertex) error {
	if _, i := m.getVertex(p.ID(), m.targets); i >= 0 {
		return vertexIDExistErr(p.ID())
	}
	if _, i := m.getVertex(p.ID(), m.sources); i >= 0 {
		return vertexIDExistErr(p.ID())
	}

	m.sources = append(m.sources, p)
	return nil
}

func (m *mockVertex) AddTarget(c Vertex) error {
	if _, i := m.getVertex(c.ID(), m.targets); i >= 0 {
		return vertexIDExistErr(c.ID())
	}
	if _, i := m.getVertex(c.ID(), m.sources); i >= 0 {
		return vertexIDExistErr(c.ID())
	}

	m.targets = append(m.targets, c)
	return nil
}

func (m *mockVertex) RemoveSource(id string) error {
	_, i := m.getVertex(id, m.sources)
	if i < 0 {
		return vertexIDNotFoundErr(id)
	}
	m.sources = append(m.sources[:i], m.sources[i+1:]...)
	return nil
}

func (m *mockVertex) RemoveTarget(id string) error {
	_, i := m.getVertex(id, m.targets)
	if i < 0 {
		return vertexIDNotFoundErr(id)
	}
	m.targets = append(m.targets[:i], m.targets[i+1:]...)
	return nil
}

func (m *mockVertex) removeVertex(i int, vertices []Vertex) ([]Vertex, error) {
	vertices = append(vertices[i:], vertices[i+1:]...)
	return vertices, nil
}

func (m *mockVertex) getVertex(id string, vertices []Vertex) (Vertex, int) {
	for i, s := range vertices {
		if s.ID() == id {
			return s, i
		}
	}
	return nil, -1
}

func TestName(t *testing.T) {
	dag := NewDAG()

	vet1 := newMockVertex("1", "abc")
	vet2 := newMockVertex("2", "def")
	vet3 := newMockVertex("3", "ghi")
	vet4 := newMockVertex("4", "cvbxvb")

	err := dag.AddVertex(vet1, vet2, vet3, vet4)
	if err != nil {
		t.Error(err)
	}

	err = dag.AddEdge(vet1, vet2)
	if err != nil {
		t.Error(err)
	}

	err = dag.AddEdge(vet1, vet3)
	if err != nil {
		t.Error(err)
	}

	err = dag.AddEdge(vet2, vet4)
	if err != nil {
		t.Error(err)
	}

	err = dag.AddEdge(vet3, vet4)
	if err != nil {
		t.Error(err)
	}

	list, err := dag.TopologicalSort()
	if err != nil {
		t.Error(err)
	}

	t.Log(list)

}
