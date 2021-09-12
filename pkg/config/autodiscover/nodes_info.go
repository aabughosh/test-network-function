package autodiscover

import "time"

const (
	resourceTypeNodes = "node"
)

// NodeList holds the data from an `oc get nodes -o json` command
type NodeList struct {
	Items []NodeResource `json:"items"`
}

// NodeResource defines deployment resources
type NodeResource struct {
	Metadata struct {
		Name        string            `json:"name"`
		Namespace   string            `json:"namespace"`
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
	} `json:"metadata"`

	Spec struct {
		Replicas int `json:"replicas"`
	}
}

// GetName returns node's metadata section's name field.
func (node *NodeResource) GetName() string {
	return node.Metadata.Name
}

// GetNamespace returns node's metadata section's namespace field.
func (node *NodeResource) GetNamespace() string {
	return node.Metadata.Namespace
}

// GetReplicas returns node's spec section's replicas field.
func (node *NodeResource) GetReplicas() int {
	return node.Spec.Replicas
}
func NodeList(timeout time.Duration, labels map[string]*string) *NodeNames {
	args := []string{"oc", "get", "nodes", "-o", "custom-columns=NAME:.metadata.name"}
	var labelsString string
	for labelName, labelValue := range labels {
		labelsString += labelName
		if labelValue != nil {
			labelsString += "=" + *labelValue
		}
		labelsString += ","
	}
	if labelsString != "" {
		labelsString = labelsString[:len(labelsString)-1]
		args = append(args, "-l", labelsString)
	}
	return &NodeNames{
		timeout: timeout,
		result:  tnf.ERROR,
		args:    args,
	}
}
