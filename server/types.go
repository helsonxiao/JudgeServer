package server

import (
	"github.com/helsonxiao/JudgeServer/judger"
)

type SpjCompileJson struct {
	Src              string           `json:"src" binding:"required"`
	SpjVersion       string           `json:"spj_version" binding:"required"`
	SpjCompileConfig SpjCompileConfig `json:"spj_compile_config" binding:"required"`
}

type SpjCompileConfig struct {
	SrcName        string `json:"src_name"`
	ExeName        string `json:"exe_name"`
	MaxCpuTime     int    `json:"max_cpu_time"`
	MaxRealTime    int    `json:"max_real_time"`
	MaxMemory      int    `json:"max_memory"`
	CompileCommand string `json:"compile_command"`
}

type SpjConfig struct {
	ExeName     string `json:"exe_name"`
	Command     string `json:"command"`
	SeccompRule string `json:"seccomp_rule"`
}

type JudgeJson struct {
	LanguageConfig   LanguageConfig   `json:"language_config" binding:"required"`
	Src              string           `json:"src" binding:"required"`
	MaxCpuTime       int              `json:"max_cpu_time" binding:"required"`
	MaxMemory        int              `json:"max_memory" binding:"required"`
	TestCaseId       string           `json:"test_case_id"`
	TestCase         string           `json:"test_case"`
	Output           bool             `json:"output" binding:"required"`
	SpjVersion       string           `json:"spj_version"`
	SpjConfig        SpjConfig        `json:"spj_config"`
	SpjCompileConfig SpjCompileConfig `json:"spj_compile_config"`
	SpjSrc           string           `json:"spj_src"`
}

type LanguageConfig struct {
	Compile struct {
		SrcName        string `json:"src_name" binding:"required"`
		ExeName        string `json:"exe_name" binding:"required"`
		MaxCpuTime     int    `json:"max_cpu_time" binding:"required"`
		MaxRealTime    int    `json:"max_real_time" binding:"required"`
		MaxMemory      int    `json:"max_memory" binding:"required"`
		CompileCommand string `json:"compile_command" binding:"required"`
	} `json:"compile" binding:"required"`
	Run struct {
		Command     string   `json:"command" binding:"required"`
		SeccompRule string   `json:"seccomp_rule" binding:"required"`
		Env         []string `json:"env" binding:"required"`
	} `json:"run" binding:"required"`
}

type JudgeResponse []struct {
	CpuTime   int
	RealTime  int
	Memory    int
	Signal    int
	ExitCode  int
	Error     judger.Error
	Result    judger.Result
	TestCase  string
	OutputMd5 string
	Output    string
}
