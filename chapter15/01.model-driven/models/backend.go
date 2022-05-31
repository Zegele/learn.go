package models

type RankServiceBackend struct {
	Name     string          `json:"name"`
	Expected ExpectedBackend `json:"expected"`
	Status   StatusBackend   `json:"status"`
}

type ExpectedBackend struct { //预期
	Image   string   `json:"image"`
	Command []string `json:"command"`
	Count   int      `json:"count"`
}

type StatusBackend struct { //当前状态
	RunningCount int      `json:"runningCount"`
	InstanceIPs  []string `json:"instanceIPs"`
}
