package task

type Task struct {
	ImageId      string  `json:"image_id"`
	ImportTaskId *string `json:"task_id"`
	SnapshotId   string  `json:"snapshot_id,omitempty"`
	Status       string  `json:"status,omitempty"`
	WaitTime     int     `json:"wait_time"`
}
