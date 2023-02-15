package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/helsonxiao/JudgeServer/compiler"
	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/judger"
	"github.com/helsonxiao/JudgeServer/utils"
)

func Judge(judgeDto JudgeDto) (JudgeResponseDto, *utils.ServerError) {
	// TODO: implement file io mode
	// const ioMode = judger.IOModeStandard
	testCase := judgeDto.TestCase
	testCaseId := judgeDto.TestCaseId
	if (testCase != "" && testCaseId != "") || (testCase == "" && testCaseId == "") {
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: "invalid parameter"}
	}

	// init
	compileConfig := judgeDto.LanguageConfig.Compile
	runConfig := judgeDto.LanguageConfig.Run
	submissionId, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: uuidErr.Error()}
	}

	// TODO: support spj

	// TODO: use custom testCase

	// init submission dir
	submissionDir, submissionDirErr := initSubmissionEnv(submissionId.String())
	// fmt.Println(submissionDir)
	if submissionDirErr != nil {
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: submissionDirErr.Error()}
	}
	// TODO: remove submissionDir in production
	// defer os.RemoveAll(submissionDir)

	// testCaseDir := filepath.Join(configs.TestCaseDir, testCaseId)
	// fmt.Println(testCaseDir)

	var exePath string
	if compileConfig != nil {
		srcPath := filepath.Join(submissionDir, compileConfig.SrcName)
		if err := ioutil.WriteFile(srcPath, []byte(judgeDto.Src), 0400); err != nil {
			return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
		}
		if err := os.Chown(srcPath, configs.CompilerUserUid, 0); err != nil {
			return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
		}

		var compileErr *utils.ServerError
		exePath, compileErr = compiler.Compile(*compileConfig, srcPath, submissionDir)
		// fmt.Println(exePath)
		if compileErr != nil {
			return nil, compileErr
		}
	} else {
		exePath = filepath.Join(submissionDir, runConfig.ExeName)
		if err := ioutil.WriteFile(exePath, []byte(judgeDto.Src), 0400); err != nil {
			return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
		}
	}
	if err := os.Chown(exePath, configs.RunUserUid, 0); err != nil {
		fmt.Println("chown exe path failed", exePath)
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
	}
	if err := os.Chmod(exePath, 0500); err != nil {
		fmt.Println("chmod exe path failed", exePath)
		return nil, &utils.ServerError{Name: "JudgeClientError", Message: err.Error()}
	}

	// TODO: init test case dir

	// fmt.Println(runConfig)

	results, judgeErr := judger.Judge(judger.JudgeParams{
		RunConfig: judger.JudgeRunConfig{
			Command:     runConfig.Command,
			SeccompRule: runConfig.SeccompRule,
			Env:         runConfig.Env,
		},
		ExePath:       exePath,
		MaxCpuTime:    judgeDto.MaxCpuTime,
		MaxMemory:     judgeDto.MaxMemory,
		SubmissionDir: submissionDir,
		TestCaseId:    testCaseId,
	})

	fmt.Println(judgeErr)
	fmt.Println(results[0].TestCase)

	return make(JudgeResponseDto, 1), nil
}

func initSubmissionEnv(submissionID string) (string, error) {
	submissionDir := filepath.Join(configs.JudgerWorkspace, submissionID)

	if err := os.MkdirAll(submissionDir, 0711); err != nil {
		return "", err
	}

	if err := os.Chown(submissionDir, configs.CompilerUserUid, configs.RunGroupGid); err != nil {
		return "", err
	}

	return submissionDir, nil
}
