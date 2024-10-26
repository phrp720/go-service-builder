package systemd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

// ToServiceFile generates the content of the .service file
func ToServiceFile(units map[string]interface{}, service map[string]interface{}, install map[string]interface{}) string {
	var sb strings.Builder
	if len(units) > 0 {
		sb.WriteString("[Units]\n")
		for key, value := range units {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
	}
	if len(service) > 0 {
		sb.WriteString("[Service]\n")
		for key, value := range service {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
	}
	if len(install) > 0 {
		sb.WriteString("[Install]\n")
		for key, value := range install {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
	}

	return sb.String()
}

func toMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		// Check if the field is not a zero value
		if !field.IsZero() {
			result[fieldName] = field.Interface()
		}
	}
	return result
}

// ToMaps converts the ServiceConfig struct to a map of unit, service and install
func (s ServiceConfig) ToMaps() (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	unitMap := toMap(s.unit)
	serviceMap := toMap(s.service)
	installMap := toMap(s.install)
	return unitMap, serviceMap, installMap
}

// Generate writes the .service file content to the specified path
func Generate(s ServiceConfig, path string) error {
	content := ToServiceFile(s.ToMaps())
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// GenerateDefault writes the .service file content to /etc/systemd/system/
func GenerateDefault(s ServiceConfig, file string) error {
	content := ToServiceFile(s.ToMaps())
	path := "/etc/systemd/system/" + file
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// GenerateAndStart writes the .service file content to /etc/systemd/system/ and starts the service
func GenerateAndStart(s ServiceConfig, file string, enable bool) error {
	if err := GenerateDefault(s, file); err != nil {
		return err
	}

	if enable {
		// Enable the service
		cmd := exec.Command("sudo", "systemctl", "enable", file)
		if err := cmd.Run(); err != nil {
			fmt.Print(err, "\n")
			return err
		}
	}

	// Start the service
	cmd := exec.Command("sudo", "systemctl", "start", file)
	if err := cmd.Run(); err != nil {
		fmt.Print(err, "\n")
		return err
	}
	return nil
}
