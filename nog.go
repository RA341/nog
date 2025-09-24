package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

const cleanArg = "--clean"

const BuildFile = "build.go"
const NogFile = "nog.go"

const NogExec = "nog"

var NogExecOld = NogExec + ".old"

func GoRebuildUrself() {
	slog.Info("args", "a", os.Args, "2", len(os.Args))

	if len(os.Args) > 1 && os.Args[1] == cleanArg {
		cleanupAfterUrSelf()
		os.Exit(0)
	}

	executablePath := GetExecutable()

	execStat, err := os.Stat(executablePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		Fatal("Failed to stat executablePath: %v", err)
	}

	nogStat, err := os.Stat(NogFile)
	if err != nil {
		Fatal("Failed to stat executablePath: %v", err)
	}

	buildStat, err := os.Stat(BuildFile)
	if err != nil {
		Fatal("Failed to stat executablePath: %v", err)
	}

	isBuildMod := buildStat.ModTime().After(execStat.ModTime())
	isNogMod := nogStat.ModTime().After(execStat.ModTime())

	//slog.Debug("Mod time result:", "rea", compare)
	if isBuildMod || isNogMod {
		slog.Info(
			"Executable is not up to date, rebuilding myself",
			BuildFile, isBuildMod,
			NogFile, isNogMod,
		)
		BuildExec(executablePath)
	} else {
		slog.Info("Executable up to date")
	}
}

func cleanupAfterUrSelf() {
	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	slog.Info("Cleaning up after MySelf")
	err := os.RemoveAll("./" + NogExecOld + exe)
	if err != nil {
		slog.Error("Failed to remove old executable: %v", err)
		return
	}
}

func GetExecutable() string {
	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	executablePath := "./" + NogExec + "" + exe
	return executablePath
}

func BuildExec(executablePath string) {
	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	rename := "./" + NogExecOld + exe

	slog.Info("Renaming to", "rename", rename)

	err := os.Rename(executablePath, rename)
	if err != nil {
		Fatal("Failed to remove executablePath: %v", err)
	}

	cmd := []string{
		"go",
		"build",
		"-o", executablePath,
		"nog.go",
		"build.go",
	}
	RunCmd(cmd...)

	slog.Info("Created new", "executable", executablePath)

	slog.Info("Calling new executable", "path", executablePath)
	newCmd := []string{executablePath}
	newCmd = append(newCmd, os.Args[1:]...)
	newExecCmd := exec.Command(newCmd[0], newCmd[1:]...)
	newExecCmd.Stdout = os.Stdout
	newExecCmd.Stderr = os.Stderr
	err = newExecCmd.Run()
	if err != nil {
		Fatal("Failed to start new executablePath: %v", err)
	}

	// remove old executable in the background
	newCmdS := exec.Command(executablePath, cleanArg)
	newCmdS.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(
		`"%s" %s`,
		executablePath,
		cleanArg,
	)}

	if err = newCmdS.Start(); err != nil {
		slog.Error("Failed to start new executablePath: %v", err)
	}

	slog.Info("Goodbye", "executable", executablePath)
	os.Exit(0)
}

func RunCmd(cmd ...string) {
	execCmd := exec.Command(cmd[0], cmd[1:]...)
	err := execCmd.Run()
	if err != nil {
		Fatal("Unable to run command: %v, %v", cmd, err)
	}
}
