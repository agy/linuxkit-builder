package task

type Task struct {
	Bucket       string  `json:"bucket"`
	ImageId      string  `json:"image_id"`
	ImportTaskId *string `json:"task_id"`
	Key          string  `json:"key"`
	Name         string  `json:"name"`
	SnapshotId   string  `json:"snapshot_id,omitempty"`
	Status       string  `json:"status,omitempty"`
	WaitTime     int     `json:"wait_time"`
}
