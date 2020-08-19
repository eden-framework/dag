package dag

type relation struct {
	sources []string
	targets []string
}

type DAG struct {
	vertex    *OrderedMap
	relations map[string]*relation
}

func NewDAG() *DAG {
	return &DAG{
		vertex:    NewOrderedMap(),
		relations: make(map[string]*relation),
	}
}

func (d *DAG) GetVertex(id string) (Vertex, bool) {
	val, exist := d.vertex.Get(id)
	if !exist {
		return nil, false
	}
	return val.(Vertex), true
}

func (d *DAG) AddVertex(vertices ...Vertex) error {
	for _, v := range vertices {
		_, found := d.vertex.PutOrGet(v.ID(), v)
		if found {
			return vertexIDExistErr(v.ID())
		}
		d.relations[v.ID()] = &relation{}
	}
	return nil
}

func (d *DAG) addRelation(sourceID, targetID string) bool {
	if r, ok := d.relations[sourceID]; ok {
		if r == nil {
			r = &relation{}
			d.relations[sourceID] = r
		}
		r.targets = append(r.targets, targetID)
	} else {
		return false
	}

	if r, ok := d.relations[targetID]; ok {
		if r == nil {
			r = &relation{}
			d.relations[targetID] = r
		}
		r.sources = append(r.sources, sourceID)
	} else {
		return false
	}

	return true
}

func (d *DAG) removeRelation(sourceID, targetID string) bool {
	if r, ok := d.relations[sourceID]; ok {
		for i, s := range r.targets {
			if targetID == s {
				r.targets = append(r.targets[:i], r.targets[i+1:]...)
			}
		}
	} else {
		return false
	}

	if r, ok := d.relations[targetID]; ok {
		for i, s := range r.sources {
			if sourceID == s {
				r.sources = append(r.sources[:i], r.sources[i+1:]...)
			}
		}
	} else {
		return false
	}

	return true
}

func (d *DAG) removeVertexRelation(id string) bool {
	if _, ok := d.relations[id]; !ok {
		return false
	}
	delete(d.relations, id)

	for _, r := range d.relations {
		for i, s := range r.sources {
			if s == id {
				r.sources = append(r.sources[:i], r.sources[i+1:]...)
			}
		}
		for i, s := range r.targets {
			if s == id {
				r.targets = append(r.targets[:i], r.targets[i+1:]...)
			}
		}
	}
	return true
}

func (d *DAG) RemoveVertex(v Vertex) error {
	v, found := d.GetVertex(v.ID())
	if !found {
		return vertexIDNotFoundErr(v.ID())
	}

	var err error
	sources := v.GetSources()
	for _, s := range sources {
		err = s.RemoveTarget(v.ID())
		if err != nil {
			return err
		}
	}

	targets := v.GetTargets()
	for _, t := range targets {
		err = t.RemoveSource(v.ID())
		if err != nil {
			return err
		}
	}

	d.vertex.Remove(v.ID())
	d.removeVertexRelation(v.ID())

	return nil
}

func (d *DAG) AddEdge(source Vertex, target Vertex) error {
	err := d.checkVertex(source)
	if err != nil {
		return err
	}

	err = d.checkVertex(target)
	if err != nil {
		return err
	}

	err = source.AddTarget(target)
	if err != nil {
		return err
	}

	err = target.AddSource(source)
	if err != nil {
		return err
	}

	if !d.addRelation(source.ID(), target.ID()) {
		return vertexRelationAddFailErr(source.ID(), target.ID())
	}
	return nil
}

func (d *DAG) RemoveEdge(source Vertex, target Vertex) error {
	err := d.checkVertex(source)
	if err != nil {
		return err
	}

	err = d.checkVertex(target)
	if err != nil {
		return err
	}

	err = source.RemoveTarget(target.ID())
	if err != nil {
		return err
	}

	err = target.RemoveSource(source.ID())
	if err != nil {
		return err
	}

	if !d.removeRelation(source.ID(), target.ID()) {
		return vertexRelationRemoveFailErr(source.ID(), target.ID())
	}
	return nil
}

func (d *DAG) Size() int {
	return d.vertex.Size()
}

// 拓扑排序
func (d *DAG) TopologicalSort() (list []Vertex, err error) {
	// searching root vertex
	var root Vertex
	d.vertex.Walk(func(key, val interface{}) bool {
		if val.(Vertex).GetSourcesLength() == 0 {
			if root != nil {
				err = rootVertexExistErr(root.ID(), val.(Vertex).ID())
				return false
			}
			root = val.(Vertex)
			list = append(list, root)
		}
		return true
	})
	if err != nil {
		return
	}

	if root == nil {
		err = rootVertexNotFoundErr
		return
	}

	if !d.removeVertexRelation(root.ID()) {
		err = vertexIDNotFoundErr(root.ID())
		return
	}

	for {
		var rootVertices = make([]string, 0)
		var found bool
		for id, r := range d.relations {
			if len(r.sources) == 0 {
				found = true
				rootVertices = append(rootVertices, id)
			}
		}

		if !found {
			break
		}

		for _, id := range rootVertices {
			if !d.removeVertexRelation(id) {
				err = vertexIDNotFoundErr(root.ID())
				return
			}

			v, _ := d.GetVertex(id)
			list = append(list, v)
		}
	}

	d.regenerateRelation()

	return
}

func (d *DAG) regenerateRelation() {
	d.relations = map[string]*relation{}
	d.vertex.Walk(func(key, val interface{}) bool {
		r := &relation{}
		v := val.(Vertex)
		for _, s := range v.GetSources() {
			r.sources = append(r.sources, s.ID())
		}
		for _, s := range v.GetTargets() {
			r.targets = append(r.targets, s.ID())
		}

		d.relations[key.(string)] = r
		return true
	})
}

func (d *DAG) checkVertex(v Vertex) error {
	_, exist := d.GetVertex(v.ID())
	if !exist {
		return vertexIDNotFoundErr(v.ID())
	}
	return nil
}
