package notifier

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"io/ioutil"
	"os"
	"os/exec"
)

type CommandNotifyRunner struct {
	Command string
}

func NewCommandNotifyRunner(command string) *CommandNotifyRunner {
	return &CommandNotifyRunner{Command: command}
}

func (cnr *CommandNotifyRunner) Trigger(mr *apis.MergeRequest) error {
	log.Infof("Start Trigger Command %s", cnr.Command)
	binary, err := exec.LookPath(cnr.Command)
	if err != nil {
		return fmt.Errorf("Commmand runner process fail.Command %#v. Error message: %s.", cnr, err)
	}
	cmd := exec.Command(binary)
	env := os.Environ()
	env = append(env, fmt.Sprintf("MR_TITLE=%s", mr.Title))
	cmd.Env = env
	cmdOut, _ := cmd.StdoutPipe()

	startErr := cmd.Start()
	if startErr != nil {
		return fmt.Errorf("Commmand runner process fail.Command %#v. Error message: %s.", cnr, err)
	}

	// read stdout and stderr
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	log.Debug(string(stdOutput[:]))

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("Commmand runner process fail.Command %#v. Error message: %s.", cnr, err)
	}
	log.Infof("Command %s Done.", cnr.Command)
	return nil
}
