package dag

import (
	"errors"
	"fmt"
)

var (
	vertexIDExistErr = func(id string) error {
		return fmt.Errorf("the vertex id: %s already exist in DAG", id)
	}
	vertexIDNotFoundErr = func(id string) error {
		return fmt.Errorf("the vertex id: %s not found in DAG", id)
	}
	vertexRelationAddFailErr = func(source, target string) error {
		return fmt.Errorf("the source id: %s add target id: %s fail", source, target)
	}
	vertexRelationRemoveFailErr = func(source, target string) error {
		return fmt.Errorf("the source id: %s remove target id: %s fail", source, target)
	}
	rootVertexExistErr = func(lastID, currentID string) error {
		return fmt.Errorf("there have tow (or more) root vertex exist in DAG, last one is: %s, this one is: %s", lastID, currentID)
	}
	rootVertexNotFoundErr = errors.New("can't find the root vertex, maybe there's a circle in DAG")
)
