package utils

type H[T any] struct {
	Err  any `json:"err"` // err should be string or nil
	Data T   `json:"data"`
}

type SpjCompileJson struct {
	Src              string           `json:"src" binding:"required"`
	SpjVersion       string           `json:"spj_version" binding:"required"`
	SpjCompileConfig SpjCompileConfig `json:"spj_compile_config" binding:"required"`
}

type SpjCompileConfig struct {
	SrcName        string `json:"src_name" binding:"required"`
	ExeName        string `json:"exe_name" binding:"required"`
	MaxCpuTime     int    `json:"max_cpu_time" binding:"required"`
	MaxRealTime    int    `json:"max_real_time" binding:"required"`
	MaxMemory      int    `json:"max_memory" binding:"required"`
	CompileCommand string `json:"compile_command" binding:"required"`
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
