package parser

import (
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"io/ioutil"
	"os"
	"strings"
)

const TF_ENV_VAR = "TF_VAR_"

func LoadEnvVars(inputVars map[string]cty.Value) (map[string]cty.Value, error){
	file, err := ioutil.TempFile("./","tmp")
	if err != nil {
		return inputVars, err
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		key, val := pair[0], pair[1]
		if strings.Contains(key, TF_ENV_VAR) {
			key := strings.TrimPrefix(key,TF_ENV_VAR)
			// Since env variables take lower precedence, only update if tfvars has not included a value yet
			if _, ok := inputVars[key]; !ok {
				file.WriteString(fmt.Sprintf("%s=%s\n", key, val))
			}
		}
	}
	envVars, err := LoadTFVars(file.Name())
	if err != nil {
		return inputVars, nil
	}

	for key, val := range envVars {
		inputVars[key] = val
	}

	return inputVars, nil
}
