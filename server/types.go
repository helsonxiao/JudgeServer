package server

import (
	"github.com/helsonxiao/JudgeServer/compiler"
	"github.com/helsonxiao/JudgeServer/judger"
)

type SpjCompileDto struct {
	Src              string                 `json:"src" binding:"required"`
	SpjVersion       string                 `json:"spj_version" binding:"required"`
	SpjCompileConfig compiler.CompileConfig `json:"spj_compile_config" binding:"required"`
}

type SpjConfig struct {
	ExeName     string `json:"exe_name"`
	Command     string `json:"command"`
	SeccompRule string `json:"seccomp_rule"`
}

type JudgeDto struct {
	LanguageConfig   LanguageConfig         `json:"language_config" binding:"required"`
	Src              string                 `json:"src" binding:"required"`
	MaxCpuTime       int                    `json:"max_cpu_time" binding:"required"`
	MaxMemory        int                    `json:"max_memory" binding:"required"`
	TestCaseId       string                 `json:"test_case_id"`
	TestCase         string                 `json:"test_case"`
	Output           bool                   `json:"output" binding:"required"`
	SpjVersion       string                 `json:"spj_version"`
	SpjConfig        SpjConfig              `json:"spj_config"`
	SpjCompileConfig compiler.CompileConfig `json:"spj_compile_config"`
	SpjSrc           string                 `json:"spj_src"`
}

type LanguageConfig struct {
	Compile *compiler.CompileConfig `json:"compile"`
	Run     struct {
		Command     string   `json:"command" binding:"required"`
		ExeName     string   `json:"exe_name"`
		Env         []string `json:"env" binding:"required"`
		SeccompRule string   `json:"seccomp_rule" binding:"required"`
	} `json:"run" binding:"required"`
}

type JudgeResponseDto []struct {
	CpuTime   int
	RealTime  int
	Memory    int
	Signal    int
	ExitCode  int
	Error     judger.ErrorCode
	Result    judger.ResultCode
	TestCase  string
	OutputMd5 string
	Output    string
}
