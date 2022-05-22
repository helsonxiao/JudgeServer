package server

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/helsonxiao/JudgeServer/configs"
)

func Judge(judgeDto JudgeDto) (JudgeResponseDto, error) {
	// const ioMode = judger.IOModeStandard
	testCase := judgeDto.TestCase
	testCaseId := judgeDto.TestCaseId
	if (testCase != "" && testCaseId != "") || (testCase == "" && testCaseId == "") {
		println("testCaseId", testCaseId)
		println("testCase", testCase)
		return nil, errors.New("invalid parameter")
	}

	// init
	// compileConfig := judgeDto.LanguageConfig.Compile
	// runConfig := judgeDto.LanguageConfig.Run
	submissionId, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return nil, uuidErr
	}

	// TODO: support spj

	// init submission dir
	submissionDir, submissionDirErr := initSubmissionEnv(submissionId.String())
	println(submissionDir)
	if submissionDirErr != nil {
		return nil, submissionDirErr
	}
	// TODO: remove submissionDir in production
	// defer os.RemoveAll(submissionDir)

	testCaseDir := filepath.Join(configs.TestCaseDir, testCaseId)
	print(testCaseDir)

	// TODO: compile codes
	// if compileConfig != nil {
	// }

	// TODO: init test case dir

	return make(JudgeResponseDto, 1), nil
}

func initSubmissionEnv(submissionID string) (string, error) {
	submissionDirPath := filepath.Join(configs.JudgerWorkspace, submissionID)

	// TODO: os.Chown
	if err := os.MkdirAll(submissionDirPath, 0777); err != nil {
		return "", err
	}

	return submissionDirPath, nil
}
