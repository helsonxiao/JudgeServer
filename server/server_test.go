package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	assert.Equal(t, reflect.String, reflect.TypeOf(body.Data.Hostname).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.Cpu).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.CpuCore).Kind())
	assert.Equal(t, reflect.Float32, reflect.TypeOf(body.Data.Memory).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(body.Data.JudgerVersion).Kind())
	assert.Equal(t, "pong", body.Data.Action)
}
