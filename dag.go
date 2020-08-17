package dag

type DAG struct {
	vertex *OrderedMap
}

func (d *DAG) GetVertex(id string) (Vertex, bool) {
	val, exist := d.vertex.Get(id)
	if !exist {
		return nil, false
	}
	return val.(Vertex), true
}

func (d *DAG) AddVertex(v Vertex) error {
	_, found := d.vertex.PutOrGet(v.ID(), v)
	if found {
		return vertexIDExistErr(v.ID())
	}
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

	return target.AddSource(source)
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

	return target.RemoveSource(source.ID())
}

func (d *DAG) Size() int {
	return d.vertex.Size()
}

func (d *DAG) checkVertex(v Vertex) error {
	_, exist := d.GetVertex(v.ID())
	if !exist {
		return vertexIDNotFoundErr(v.ID())
	}
	return nil
}
