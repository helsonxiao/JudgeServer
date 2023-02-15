package judger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/utils"
)

const DEFAULT_MAX_OUTPUT_SIZE = 16 * 1024 * 1024

func Judge(params JudgeParams) ([]JudgeResult, *utils.ServerError) {
	testCaseDir := path.Join(configs.TestCaseDir, params.TestCaseId)
	testCaseInfoPath := path.Join(testCaseDir, "info")
	infoBytes, infoErr := ioutil.ReadFile(testCaseInfoPath)
	if infoErr != nil {
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: infoErr.Error()}
	}
	var testCaseInfo TestCaseInfo
	if err := json.Unmarshal(infoBytes, &testCaseInfo); err != nil {
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
	}
	// fmt.Println(testCaseInfo.TestCases)

	command := utils.FillWith(params.RunConfig.Command, map[string]string{
		"{exe_path}":   params.ExePath,
		"{exe_dir}":    params.SubmissionDir,
		"{max_memory}": strconv.Itoa(params.MaxMemory / 1024),
	})

	// 由于部分语言的编译产物无法直接运行，RunConfig.Command 指令模板会确保第一个 cut 可执行，其余 cut 作为参数传递
	commandCuts := strings.Split(command, " ")
	exePath := commandCuts[0]
	args := commandCuts[1:]

	env := []string{}
	env = append(env, params.RunConfig.Env...)
	env = append(env, "PATH="+os.Getenv("PATH"))

	// 收集并发评测的结果
	var results []JudgeResult
	var resultsWg sync.WaitGroup
	resultChan := make(chan JudgeResult, len(testCaseInfo.TestCases))
	resultsWg.Add(1)
	go func() {
		defer resultsWg.Done()
		for r := range resultChan {
			results = append(results, r)
		}
	}()

	// 并发评测
	var judgeWg sync.WaitGroup
	for testCaseName, testCaseDetail := range testCaseInfo.TestCases {
		// fmt.Println(testCaseDetail)
		judgeWg.Add(1)
		go judgeOne(&judgeWg, resultChan, JudgeOneParams{
			Args:           args,
			Env:            env,
			ExePath:        exePath,
			MaxCpuTime:     params.MaxCpuTime,
			MaxMemory:      params.MaxMemory,
			SeccompRule:    params.RunConfig.SeccompRule,
			SubmissionDir:  params.SubmissionDir,
			TestCaseDir:    testCaseDir,
			TestCaseName:   testCaseName,
			TestCaseDetail: testCaseDetail,
		})
	}
	judgeWg.Wait()
	close(resultChan)
	resultsWg.Wait()
	return results, nil
}

func judgeOne(wg *sync.WaitGroup, resultChan chan JudgeResult, params JudgeOneParams) {
	defer wg.Done()
	// fmt.Println("judgeOne", params.TestCaseName)
	inputPath := path.Join(params.TestCaseDir, params.TestCaseName+".in")
	// fmt.Println("inputPath", inputPath)

	// TODO: support file io

	userOutputPath := path.Join(params.SubmissionDir, params.TestCaseName+".out")
	// fmt.Println("userOutputPath", userOutputPath)

	maxOutputSize := DEFAULT_MAX_OUTPUT_SIZE
	if params.TestCaseDetail.OutputSize != nil {
		maxOutputSize = Max(maxOutputSize, *params.TestCaseDetail.OutputSize*2)
	}

	runResult, runErr := Run(RunConfig{
		MaxCpuTime:           params.MaxCpuTime,
		MaxRealTime:          params.MaxCpuTime * 3,
		MaxMemory:            params.MaxMemory,
		MaxStack:             128 * 1024 * 1024,
		MaxOutputSize:        maxOutputSize,
		MaxProcessNumber:     -1, // unlimited
		MemoryLimitCheckOnly: 0,  // strict mode
		Env:                  params.Env,
		ExePath:              params.ExePath,
		Args:                 params.Args,
		InputPath:            inputPath,
		OutputPath:           userOutputPath,
		ErrorPath:            userOutputPath,
		LogPath:              configs.JudgerLogPath,
		SecCompRuleName:      &params.SeccompRule,
		Uid:                  configs.RunUserUid,
		Gid:                  configs.RunGroupGid,
	})
	if runErr != nil {
		// TODO: log runErr.Error()
		resultChan <- JudgeResult{
			CpuTime:   -1,
			Error:     -1,
			ExitCode:  -1,
			OutputMd5: "",
			Output:    "",
			Memory:    -1,
			RealTime:  -1,
			Result:    ResultSystemError,
			Signal:    -1,
			TestCase:  params.TestCaseName,
		}
		return
	}

	result := JudgeResult{
		CpuTime:   runResult.CpuTime,
		Error:     runResult.Error,
		ExitCode:  runResult.ExitCode,
		OutputMd5: "", // TODO
		Output:    "", // TODO
		Memory:    runResult.Memory,
		RealTime:  runResult.RealTime,
		Result:    runResult.Result,
		Signal:    runResult.Signal,
		TestCase:  params.TestCaseName,
	}
	fmt.Println(result.TestCase)
	// TODO: check and update result
	resultChan <- result
}

func spj() {}

func compareOutput() {}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
