package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/judger"
	"github.com/helsonxiao/JudgeServer/utils"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var body utils.H[utils.ServerInfo]
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, body.Err)
	assert.Equal(t, reflect.String, reflect.TypeOf(body.Data.Hostname).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.Cpu).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.CpuCore).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.Memory).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(body.Data.JudgerVersion).Kind())
	assert.Equal(t, "pong", body.Data.Action)
}

var defaultEnv = []string{"LANG=en_US.UTF-8", "LANGUAGE=en_US:en", "LC_ALL=en_US.UTF-8"}

const py3Src = `
s = input()
s1 = s.split(" ")
print(int(s1[0]) + int(s1[1]))
`

const cSpjSrc = `
#include <stdio.h>
int main(){
	return 1;
}
`

func TestJudgeRoute(t *testing.T) {
	configs.SetupEnv()
	router := SetupRouter()
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(map[string]any{
		"language_config": map[string]any{
			"compile": map[string]any{
				"src_name":        "solution.py",
				"exe_name":        "__pycache__/solution.cpython-36.pyc",
				"max_cpu_time":    3000,
				"max_real_time":   5000,
				"max_memory":      128 * 1024 * 1024,
				"compile_command": "/usr/bin/python3 -m py_compile {src_path}",
			},
			"run": map[string]any{
				"command":      "/usr/bin/python3 {exe_path}",
				"seccomp_rule": "general",
				"env":          append([]string{"PYTHONIOENCODING=UTF-8"}, defaultEnv...),
			},
		},
		"src":          py3Src,
		"max_cpu_time": 1000,
		"max_memory":   128 * 1024 * 1024,
		"test_case_id": "normal",
		"output":       true,
	})
	req, _ := http.NewRequest("POST", "/judge", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resBody utils.H[JudgeResponseDto]
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, resBody.Err)
	assert.NotEqual(t, 0, len(resBody.Data))
	assert.Equal(t, reflect.Int, reflect.TypeOf(resBody.Data[0].CpuTime).Kind())
	assert.Equal(t, reflect.Int, reflect.TypeOf(resBody.Data[0].RealTime).Kind())
	assert.Equal(t, reflect.Int, reflect.TypeOf(resBody.Data[0].Memory).Kind())
	assert.Equal(t, 0, resBody.Data[0].Signal)
	assert.Equal(t, 0, resBody.Data[0].ExitCode)
	assert.Equal(t, judger.ErrorSuccess, resBody.Data[0].Error)
	assert.Equal(t, judger.ResultSuccess, resBody.Data[0].Result)
	assert.Equal(t, reflect.String, reflect.TypeOf(resBody.Data[0].TestCase).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(resBody.Data[0].OutputMd5).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(resBody.Data[0].Output).Kind())
}

func TestCompileSpjRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(map[string]any{
		"src":         cSpjSrc,
		"spj_version": "2",
		"spj_compile_config": map[string]any{
			"src_name":        "spj-{spj_version}.c",
			"exe_name":        "spj-{spj_version}",
			"max_cpu_time":    3000,
			"max_real_time":   5000,
			"max_memory":      1024 * 1024 * 1024,
			"compile_command": "/usr/bin/gcc -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c99 {src_path} -lm -o {exe_path}",
		},
	})
	req, _ := http.NewRequest("POST", "/compile_spj", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resBody utils.H[string]
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, resBody.Err)
	assert.Equal(t, "success", resBody.Data)
}

func TestCompileSpjErrRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(map[string]any{
		"src": cSpjSrc,
	})
	req, _ := http.NewRequest("POST", "/compile_spj", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resBody utils.H[any]
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotNil(t, resBody.Err)
	assert.Nil(t, resBody.Data)
}
