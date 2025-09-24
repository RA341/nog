package main

import (
	"errors"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	Exe()
}

const (
	cleanArg = "--clean"

	BuildFile = "build.go"
	NogFile   = "nog.go"

	NogExec = "nog"
)

var (
	exe        = ""
	NogExecOld = NogExec + ".old"
)

func Exe() {
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
}

func MustStat(path string) os.FileInfo {
	info, err := os.Stat(path)
	if err != nil {
		Fatal(
			"Failed to stat",
			"err", err,
			"path", path,
		)
	}
	return info
}

func GoRebuildUrself() {
	if len(os.Args) > 1 && os.Args[1] == cleanArg {
		cleanupAfterUrSelf()
		os.Exit(0)
	}

	executablePath := GetExecutable()
	rebuild := shouldRebuild(executablePath)
	if rebuild {
		slog.Info("rebuilding myself")
		BuildExec(executablePath)
	} else {
		slog.Info("up to date")
	}
}

func shouldRebuild(executablePath string) bool {
	execStat, err := os.Stat(executablePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		Fatal("Failed to stat executablePath: %v", err)
	}

	nogStat := MustStat(NogFile)
	buildStat := MustStat(BuildFile)

	isBuildMod := buildStat.ModTime().After(execStat.ModTime())
	isNogMod := nogStat.ModTime().After(execStat.ModTime())
	return isBuildMod || isNogMod
}

func cleanupAfterUrSelf() {
	slog.Info("Cleaning up after MySelf")
	err := os.RemoveAll("./" + NogExecOld + exe)
	if err != nil {
		slog.Error("Failed to remove old executable: %v", err)
		return
	}
}

func GetExecutable() string {
	executablePath := "./" + NogExec + "" + exe
	return executablePath
}

func BuildExec(executablePath string) {
	rename := "./" + NogExecOld + exe

	slog.Info("Renaming old build exec",
		"from", executablePath,
		"to", rename,
	)
	err := os.Rename(executablePath, rename)
	if err != nil {
		Fatal("Failed to rename path",
			"err", err,
			"from", executablePath,
			"to", rename,
		)
	}

	cmd := []string{
		"go",
		"build",
		"-o", executablePath,
		NogFile,
		BuildFile,
	}
	RunCmd(cmd[0], cmd[1:]...)

	slog.Info("Calling new executable", "path", executablePath)
	RunCmd(executablePath, os.Args[1:]...)

	removeOldExec(executablePath)

	slog.Info("Goodbye...")
	os.Exit(0)
}

func removeOldExec(executablePath string) {
	// remove old executable in the background
	newCmdS := exec.Command(executablePath, cleanArg)
	if err := newCmdS.Start(); err != nil {
		slog.Error("Failed to start new executablePath: %v", err)
	}
}

func RunCmd(executable string, args ...string) {
	execCmd := exec.Command(executable, args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin

	err := execCmd.Run()
	if err != nil {
		Fatal("Unable to run command: %v, %v", args, err)
	}
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
