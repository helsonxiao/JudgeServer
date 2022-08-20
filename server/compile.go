package server

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/helsonxiao/JudgeServer/compiler"
	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/utils"
)

func CompileSpj(dto SpjCompileDto) (bool, *utils.ServerError) {
	src := dto.Src
	spjVersion := dto.SpjVersion
	compileConfig := dto.SpjCompileConfig
	compileConfig.SrcName = strings.Replace(compileConfig.SrcName, "{spj_version}", spjVersion, -1)
	compileConfig.ExeName = strings.Replace(compileConfig.ExeName, "{spj_version}", spjVersion, -1)
	spjSrcPath := path.Join(configs.SpjSrcDir, compileConfig.SrcName)
	ioutil.WriteFile(spjSrcPath, []byte(src), 0400)

	exePath, err := compiler.Compile(compileConfig, spjSrcPath, configs.SpjExeDir)
	if err != nil {
		if err.Name == "CompileError" {
			err.Name = "SPJCompileError"
		}
		return false, err
	}

	os.Chown(exePath, configs.SpjUserUid, 0)
	os.Chmod(exePath, 0500)
	return true, nil
}
