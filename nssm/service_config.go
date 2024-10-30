package nssm

// ServiceConfig basic available  values where found here: https://gist.github.com/magnetikonline/2217fd95cf15a0324696
// ServiceConfig holds the configuration for an NSSM service and the service name
type ServiceConfig struct {
	serviceName string
	parameters  Parameters
}

// Parameters holds the configuration options for an NSSM service
type Parameters struct {
	AppAffinity                  string
	AppDirectory                 string
	AppEnvironment               string
	AppEnvironmentExtra          string
	AppExit                      string
	Application                  string
	AppNoConsole                 string
	AppParameters                string
	AppPriority                  string
	AppRestartDelay              string
	AppRotateBytes               string
	AppRotateBytesHigh           string
	AppRotateFiles               string
	AppRotateOnline              string
	AppRotateSeconds             string
	AppStderr                    string
	AppStderrCreationDisposition string
	AppStderrFlagsAndAttributes  string
	AppStderrShareMode           string
	AppStdin                     string
	AppStdinCreationDisposition  string
	AppStdinFlagsAndAttributes   string
	AppStdinShareMode            string
	AppStdout                    string
	AppStdoutCreationDisposition string
	AppStdoutFlagsAndAttributes  string
	AppStdoutShareMode           string
	AppStopMethodConsole         string
	AppStopMethodSkip            string
	AppStopMethodThreads         string
	AppStopMethodWindow          string
	AppThrottle                  string
	DependOnGroup                string
	DependOnService              string
	Description                  string
	DisplayName                  string
	ImagePath                    string
	Name                         string
	ObjectName                   string
	Start                        string
	Type                         string
}
