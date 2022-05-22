package configs

import "os"

var JudgerWorkspace string
var TestCaseDir string

func SetupEnv() {
	JudgerWorkspace = getEnv("JUDGER_WORKSPACE", "/judger/run")
	TestCaseDir = getEnv("TEST_CASE_DIR", "/test_case")
}

func getEnv(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		return defaultValue
	} else {
		return envValue
	}
}
