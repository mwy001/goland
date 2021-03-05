package cmdlog

import (
	"os/exec"
	"github.com/mwy001/goland/pkg/log"
)

func logCmdExecute(jobid string, cmd *exec.Cmd) {
	log.L().Infof("Execute cmd, %v, %v, jobid, %v", cmd.Path, cmd.Args, jobid)
}

func logCmdSuccess(jobid string, cmd *exec.Cmd) {
	log.L().Infof("Cmd success, %v, %v, jobid, %v", cmd.Path, cmd.Args, jobid)
}

func logCmdError(jobid string, cmd *exec.Cmd, err error) {
	log.L().Errorf("Error execute cmd, %v, %v, jobid, %v, err, %v", cmd.Path, cmd.Args, jobid, err)
}

func executeOneCmd(jobid string, cmd *exec.Cmd) bool {
	logCmdExecute(jobid, cmd)
	err := cmd.Run()

	if err != nil {
		logCmdError(jobid, cmd, err)
		return false
	}

	logCmdSuccess(jobid, cmd)

	return true
}

func executeOneCmdWithOutput(jobid string, cmd *exec.Cmd) (string, error) {
	logCmdExecute(jobid, cmd)
	output, err := cmd.CombinedOutput()

	if err != nil {
		logCmdError(jobid, cmd, err)
	} else {
		logCmdSuccess(jobid, cmd)
	}

	return string(output), err
}
