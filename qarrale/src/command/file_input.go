package command

import (
	"compmath4/qarrale/src/input_output"
	"compmath4/qarrale/src/util"
	"os"
)

type fileInput struct {
	argc uint8
	name string
}

func (t *fileInput) execute(argc ...string) {

	file, err := os.Open(argc[0])
	util.HandleError(&err)
	defer file.Close()

	_, _ = input_output.ReadData(file)

}

func (t *fileInput) getArgs() uint8 {
	return t.argc
}

func (t *fileInput) getName() string {
	return t.name
}
