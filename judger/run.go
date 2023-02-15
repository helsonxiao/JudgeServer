package judger

import (
	"encoding/json"
	"os/exec"
	"strconv"

	"github.com/helsonxiao/JudgeServer/configs"
)

// Judger version 2.1.1
// https://github.com/QingdaoU/Judger/blob/016653cedb0765d96ba999d86e14a033cb2fa875/src/main.c#L13
func Run(config RunConfig) (*RunResult, error) {
	args := []string{}

	// parsing int args
	args = append(args, "--max_cpu_time="+strconv.Itoa(config.MaxCpuTime))
	args = append(args, "--max_real_time="+strconv.Itoa(config.MaxRealTime))
	args = append(args, "--max_memory="+strconv.Itoa(config.MaxMemory))
	args = append(args, "--max_stack="+strconv.Itoa(config.MaxStack))
	args = append(args, "--max_process_number="+strconv.Itoa(config.MaxProcessNumber))
	args = append(args, "--max_output_size="+strconv.Itoa(config.MaxOutputSize))
	args = append(args, "--memory_limit_check_only="+strconv.Itoa(config.MemoryLimitCheckOnly))
	args = append(args, "--uid="+strconv.Itoa(config.Uid))
	args = append(args, "--gid="+strconv.Itoa(config.Gid))

	// parsing string args
	args = append(args, "--exe_path="+config.ExePath)
	args = append(args, "--input_path="+config.InputPath)
	args = append(args, "--output_path="+config.OutputPath)
	args = append(args, "--error_path="+config.ErrorPath)
	args = append(args, "--log_path="+config.LogPath)
	if config.SecCompRuleName != nil {
		args = append(args, "--seccomp_rule_name="+*config.SecCompRuleName)
	}

	// parsing list args
	for _, arg := range config.Args {
		args = append(args, "--args="+arg)
	}
	for _, env := range config.Env {
		args = append(args, "--env="+env)
	}

	// fmt.Println(args)
	cmd := exec.Command(configs.JudgerPath, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result RunResult
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
