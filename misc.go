package dag

import "fmt"

var (
	vertexIDExistErr = func(id string) error {
		return fmt.Errorf("the vertex id: %s already exist in DAG", id)
	}
	vertexIDNotFoundErr = func(id string) error {
		return fmt.Errorf("the vertex id: %s not found in DAG", id)
	}
)
