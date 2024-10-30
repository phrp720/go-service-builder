package nssm

import (
	"embed"
	"fmt"
	"github.com/phrp720/service-builder/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//go:embed bin/win64/nssm.exe
//go:embed bin/win32/nssm.exe
var embeddedFiles embed.FS
var nssmPath string

func InitNssm(path string) {
	var err error
	nssmPath, err = ExtractNssm(path)
	if err != nil {
		log.Fatalf("Failed to extract nssm.exe: %v", err)
	}
}

// extractFile extracts a file from the embedded files
func extractFile(embeddedPath, outputPath string) error {
	data, err := embeddedFiles.ReadFile(embeddedPath)
	if err != nil {
		return err
	}
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0755)
}

// ExtractNssm extracts the nssm.exe file from the embedded files
func ExtractNssm(outputDir string) (string, error) {
	var embeddedPath string
	if runtime.GOARCH == "amd64" {
		embeddedPath = "bin/win64/nssm.exe"
	} else {
		embeddedPath = "bin/win32/nssm.exe"
	}
	outputPath := filepath.Join(outputDir, "nssm.exe")
	if err := extractFile(embeddedPath, outputPath); err != nil {
		return "", err
	}
	return outputPath, nil
}

// CreateService sets the service configuration using nssm.exe
func CreateService(s ServiceConfig) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	params := util.ToMap(s.Parameters)
	cmd := exec.Command(nssmPath, "install", s.ServiceName, s.Parameters.AppDirectory)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run:  %w", err)
	}
	// Set each variable for the service using nssm.exe
	for key, value := range params {
		cmd := exec.Command(nssmPath, "set", s.ServiceName, key, fmt.Sprintf("%v", value))
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
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "start", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start the service: %w", err)
	}
	return nil
}

// StopService stops the service using nssm.exe
func StopService(service string) error {
	cmd := exec.Command(nssmPath, "stop", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop the  service: %w", err)
	}
	return nil
}

// RemoveService removes the service using nssm.exe
func RemoveService(service string) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "remove", service, "confirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove the service: %w", err)
	}
	return nil
}

// RestartService restarts the service
func RestartService(service string) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "stop", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to restart the service: %w", err)
	}
	return nil
}
func PauseService(service string) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "pause", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to pause the service: %w", err)
	}
	return nil
}

func ContinueService(service string) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "continue", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to continue the service: %w", err)
	}
	return nil
}

func RotateService(service string) error {
	if nssmPath == "" {
		log.Fatalf("nssm isn't initialized")
	}
	cmd := exec.Command(nssmPath, "rotate", service)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to rotate the service: %w", err)
	}
	return nil
}
