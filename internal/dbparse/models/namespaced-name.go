package models

const DefaultNamespace = "public"

type NamespacedName struct {
	Namespace string
	BaseName  string
}

func NewNamespacedNameSafe(namespace string, name string) *NamespacedName {
	return &NamespacedName{
		Namespace: namespace,
		BaseName:  name,
	}
}

func NewNamespacedName(namespace *string, name string) *NamespacedName {
	var createdNS string
	if namespace != nil {
		createdNS = *namespace
	} else {
		createdNS = DefaultNamespace
	}
	return &NamespacedName{
		Namespace: createdNS,
		BaseName:  name,
	}
}

func (n *NamespacedName) FullName() string {
	return n.Namespace + "." + n.BaseName
}
