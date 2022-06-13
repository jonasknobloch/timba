package tim

type Node struct {
	ID     string `json:"id"`
	Parent string `json:"parent,omitempty"`
}

func NewNodeFromTask(task *Task, parent *Node) *Node {
	if parent == nil {
		return &Node{
			ID: task.ID,
		}
	}

	return &Node{
		ID:     task.ID,
		Parent: parent.ID,
	}
}

func NewNodeFromGroup(group *Group) *Node {
	if group == nil {
		return nil
	}

	return &Node{
		ID: group.ID,
	}
}
