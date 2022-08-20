package configs

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

var JudgerPath string
var JudgerWorkspace string
var TestCaseDir string
var SpjSrcDir string
var SpjExeDir string

// Logs
var LogDir string
var CompilerLogPath string
var JudgerLogPath string
var ServerLogPath string

// Uid / Gid
var CompilerUserUid int
var CompilerGroupGid int
var RunUserUid int
var RunGroupGid int
var SpjUserUid int
var SpjGroupGid int

func SetupEnv() {
	JudgerPath = getEnv("JUDGER_PATH", "/usr/lib/judger/libjudger.so")
	JudgerWorkspace = getEnv("JUDGER_WORKSPACE", "/judger/run")
	TestCaseDir = getEnv("TEST_CASE_DIR", "/test_case")
	SpjSrcDir = getEnv("SPJ_SRC_DIR", "/judger/spj")
	SpjExeDir = getEnv("SPJ_EXE_DIR", "/judger/spj")

	LogDir = getEnv("LOG_DIR", "/log")
	CompilerLogPath = filepath.Join(LogDir, "compile.log")
	JudgerLogPath = filepath.Join(LogDir, "judger.log")
	ServerLogPath = filepath.Join(LogDir, "server.log")

	CompilerUser, _ := user.Lookup("compiler")
	CompilerUserUid, _ = strconv.Atoi(CompilerUser.Uid)
	CompilerUserGroup, _ := user.LookupGroup("compiler")
	CompilerGroupGid, _ = strconv.Atoi(CompilerUserGroup.Gid)
	RunUser, _ := user.Lookup("code")
	RunUserUid, _ = strconv.Atoi(RunUser.Uid)
	RunUserGroup, _ := user.LookupGroup("code")
	RunGroupGid, _ = strconv.Atoi(RunUserGroup.Gid)
	SpjUser, _ := user.Lookup("spj")
	SpjUserUid, _ = strconv.Atoi(SpjUser.Uid)
	SpjUserGroup, _ := user.LookupGroup("spj")
	SpjGroupGid, _ = strconv.Atoi(SpjUserGroup.Gid)
}

func getEnv(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		return defaultValue
	} else {
		return envValue
	}
}
