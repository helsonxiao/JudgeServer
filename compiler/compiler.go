package compiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/judger"
	"github.com/helsonxiao/JudgeServer/utils"
)

type CompileConfig struct {
	SrcName        string    `json:"src_name"`
	ExeName        string    `json:"exe_name"`
	MaxCpuTime     int       `json:"max_cpu_time"`
	MaxRealTime    int       `json:"max_real_time"`
	MaxMemory      int       `json:"max_memory"`
	CompileCommand string    `json:"compile_command"`
	Env            *[]string `json:"env"`
}

func Compile(config CompileConfig, srcPath string, outputDir string) (string, *utils.ServerError) {
	exePath := filepath.Join(outputDir, config.ExeName)
	replacements := map[string]string{
		"{src_path}": srcPath,
		"{exe_dir}":  outputDir,
		"{exe_path}": exePath,
	}
	command := utils.FillWith(config.CompileCommand, replacements)
	// fmt.Println("compile command: ", command)
	compilerOut := filepath.Join(outputDir, "compiler.out")

	args := strings.Split(command, " ")
	// fmt.Println(args)

	os.Chdir(outputDir)
	env := []string{}
	if config.Env != nil {
		env = *config.Env
	}
	env = append(env, "PATH="+os.Getenv("PATH"))

	result, runErr := judger.Run(judger.Config{
		MaxCpuTime:           config.MaxCpuTime,
		MaxRealTime:          config.MaxRealTime,
		MaxMemory:            config.MaxMemory,
		MaxStack:             128 * 1024 * 1024,
		MaxOutPutSize:        20 * 1024 * 1024,
		MaxProcessNumber:     -1,
		MemoryLimitCheckOnly: 0,
		ExePath:              args[0],
		Args:                 args[1:],
		InputPath:            srcPath,
		OutputPath:           compilerOut,
		ErrorPath:            compilerOut,
		LogPath:              configs.CompilerLogPath,
		SecCompRuleName:      nil,
		Uid:                  configs.CompilerUserUid,
		Gid:                  configs.CompilerGroupGid,
		Env:                  env,
	})
	if runErr != nil {
		return "", &utils.ServerError{Name: "JudgerError", Message: runErr.Error()}
	}

	if result.Result == judger.ResultSuccess {
		os.Remove(compilerOut)
		return exePath, nil
	}

	// fmt.Println("Compile Result: ", result)
	_, err := os.Stat(compilerOut)
	var errOut string
	if err == nil {
		errBytes, _ := ioutil.ReadFile(compilerOut)
		errOut = string(errBytes)
		os.Remove(compilerOut)
	} else {
		errOut = fmt.Sprintf("Compiler runtime error, info: %#v", result)
	}
	return exePath, &utils.ServerError{Name: "CompileError", Message: errOut}
}
