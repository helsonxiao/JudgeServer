package utils

type H struct {
	Err  interface{} `json:"err"` // err must be string or nil
	Data interface{} `json:"data"`
}

type ServerInfo struct {
	Hostname      string  `json:"hostanme"`
	Cpu           float32 `json:"cpu"`
	CpuCore       float32 `json:"cpu_core"`
	Memory        float32 `json:"memory"`
	JudgerVersion string  `json:"judger_version"`
	Action        string  `json:"action"`
}

// TODO: implement this dummy func
func GetServerInfo() ServerInfo {
	info := ServerInfo{
		Hostname:      "1",
		Cpu:           0.3,
		CpuCore:       8,
		Memory:        14.4,
		JudgerVersion: "2.1.1",
		Action:        "pong",
	}
	return info
}