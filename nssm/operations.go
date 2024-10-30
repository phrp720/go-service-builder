package nssm

import (
	"fmt"
	"github.com/phrp720/service-builder/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getNssm() string {
	// Determine the appropriate nssm.exe based on the system architecture
	var nssmPath string
	if runtime.GOARCH == "amd64" {
		nssmPath = "nssm/bin/win64/nssm.exe"
	} else {
		nssmPath = "nssm/bin/win32/nssm.exe"
	}
	absPath, err := filepath.Abs(nssmPath)
	if err != nil {
		panic(err)
	}
	return absPath
}

// CreateService sets the service configuration using nssm.exe
func CreateService(s ServiceConfig) error {
	params := util.ToMap(s.parameters)
	nssm := getNssm()
	cmd := exec.Command(nssm, "install", s.serviceName, s.parameters.AppDirectory)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run:  %w", err)
	}
	// Set each variable for the service using nssm.exe
	for key, value := range params {
		cmd := exec.Command(nssm, "set", s.serviceName, key, fmt.Sprintf("%v", value))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set %s: %w", key, err)
		}
	}

	return nil
}

// StartService starts the service using nssm.exe
func StartService(service string) error {
	cmd := exec.Command(getNssm(), "start", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start the service: %w", err)
	}
	return nil
}

// StopService stops the service using nssm.exe
func StopService(service string) error {
	cmd := exec.Command(getNssm(), "stop", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop the  service: %w", err)
	}
	return nil
}

// RemoveService removes the service using nssm.exe
func RemoveService(service string) error {
	cmd := exec.Command(getNssm(), "remove", service, "confirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove the service: %w", err)
	}
	return nil
}

// RestartService restarts the service
func RestartService(service string) error {
	cmd := exec.Command(getNssm(), "stop", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to restart the service: %w", err)
	}
	return nil
}
func PauseService(service string) error {
	cmd := exec.Command(getNssm(), "pause", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to pause the service: %w", err)
	}
	return nil
}

func ContinueService(service string) error {
	cmd := exec.Command(getNssm(), "continue", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to continue the service: %w", err)
	}
	return nil
}

func RotateService(service string) error {
	cmd := exec.Command(getNssm(), "rotate", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to rotate the service: %w", err)
	}
	return nil
}
