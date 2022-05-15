package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

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

const cSpjSrc = `
#include <stdio.h>
int main(){
		return 1;
}
`

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
