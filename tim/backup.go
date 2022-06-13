package tim

import "encoding/json"

type Backup struct {
	Tasks  map[string]*Task  `json:"tasks"`
	Groups map[string]*Group `json:"groups"`
	Nodes  []*Node           `json:"nodes"`
}

func NewBackup() *Backup {
	return &Backup{
		Tasks:  make(map[string]*Task),
		Groups: make(map[string]*Group),
		Nodes:  make([]*Node, 0),
	}
}

func (b *Backup) AddTask(task *Task, group *Group) {
	b.Tasks[task.ID] = task
	b.Nodes = append(b.Nodes, NewNodeFromTask(task, NewNodeFromGroup(group)))
}

func (b *Backup) AddGroup(group *Group) {
	b.Groups[group.ID] = group
	b.Nodes = append(b.Nodes, NewNodeFromGroup(group))
}

func (b *Backup) ToJSON() (string, error) {
	r, err := json.MarshalIndent(b, "", "  ")

	if err != nil {
		return "", err
	}

	return string(r), nil
}
