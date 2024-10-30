package nssm

// ServiceBuilder is a builder for creating ServiceConfig
type ServiceBuilder struct {
	service ServiceConfig
}

// NewServiceBuilder creates a new ServiceBuilder
func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		service: ServiceConfig{},
	}
}
func (b *ServiceBuilder) ServiceName(serviceName string) *ServiceBuilder {
	b.service.ServiceName = serviceName
	return b
}

func (b *ServiceBuilder) AppAffinity(appAffinity string) *ServiceBuilder {
	b.service.Parameters.AppAffinity = appAffinity
	return b
}

func (b *ServiceBuilder) AppDirectory(appDirectory string) *ServiceBuilder {
	b.service.Parameters.AppDirectory = appDirectory
	return b
}

func (b *ServiceBuilder) AppEnvironment(appEnvironment string) *ServiceBuilder {
	b.service.Parameters.AppEnvironment = appEnvironment
	return b
}

func (b *ServiceBuilder) AppEnvironmentExtra(appEnvironmentExtra string) *ServiceBuilder {
	b.service.Parameters.AppEnvironmentExtra = appEnvironmentExtra
	return b
}

func (b *ServiceBuilder) AppExit(appExit string) *ServiceBuilder {
	b.service.Parameters.AppExit = appExit
	return b
}

func (b *ServiceBuilder) Application(application string) *ServiceBuilder {
	b.service.Parameters.Application = application
	return b
}

func (b *ServiceBuilder) AppNoConsole(appNoConsole string) *ServiceBuilder {
	b.service.Parameters.AppNoConsole = appNoConsole
	return b
}

func (b *ServiceBuilder) AppParameters(appParameters string) *ServiceBuilder {
	b.service.Parameters.AppParameters = appParameters
	return b
}

func (b *ServiceBuilder) AppPriority(appPriority string) *ServiceBuilder {
	b.service.Parameters.AppPriority = appPriority
	return b
}

func (b *ServiceBuilder) AppRestartDelay(appRestartDelay string) *ServiceBuilder {
	b.service.Parameters.AppRestartDelay = appRestartDelay
	return b
}

func (b *ServiceBuilder) AppRotateBytes(appRotateBytes string) *ServiceBuilder {
	b.service.Parameters.AppRotateBytes = appRotateBytes
	return b
}

func (b *ServiceBuilder) AppRotateBytesHigh(appRotateBytesHigh string) *ServiceBuilder {
	b.service.Parameters.AppRotateBytesHigh = appRotateBytesHigh
	return b
}

func (b *ServiceBuilder) AppRotateFiles(appRotateFiles string) *ServiceBuilder {
	b.service.Parameters.AppRotateFiles = appRotateFiles
	return b
}

func (b *ServiceBuilder) AppRotateOnline(appRotateOnline string) *ServiceBuilder {
	b.service.Parameters.AppRotateOnline = appRotateOnline
	return b
}

func (b *ServiceBuilder) AppRotateSeconds(appRotateSeconds string) *ServiceBuilder {
	b.service.Parameters.AppRotateSeconds = appRotateSeconds
	return b
}

func (b *ServiceBuilder) AppStderr(appStderr string) *ServiceBuilder {
	b.service.Parameters.AppStderr = appStderr
	return b
}

func (b *ServiceBuilder) AppStderrCreationDisposition(appStderrCreationDisposition string) *ServiceBuilder {
	b.service.Parameters.AppStderrCreationDisposition = appStderrCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStderrFlagsAndAttributes(appStderrFlagsAndAttributes string) *ServiceBuilder {
	b.service.Parameters.AppStderrFlagsAndAttributes = appStderrFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStderrShareMode(appStderrShareMode string) *ServiceBuilder {
	b.service.Parameters.AppStderrShareMode = appStderrShareMode
	return b
}

func (b *ServiceBuilder) AppStdin(appStdin string) *ServiceBuilder {
	b.service.Parameters.AppStdin = appStdin
	return b
}

func (b *ServiceBuilder) AppStdinCreationDisposition(appStdinCreationDisposition string) *ServiceBuilder {
	b.service.Parameters.AppStdinCreationDisposition = appStdinCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStdinFlagsAndAttributes(appStdinFlagsAndAttributes string) *ServiceBuilder {
	b.service.Parameters.AppStdinFlagsAndAttributes = appStdinFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStdinShareMode(appStdinShareMode string) *ServiceBuilder {
	b.service.Parameters.AppStdinShareMode = appStdinShareMode
	return b
}

func (b *ServiceBuilder) AppStdout(appStdout string) *ServiceBuilder {
	b.service.Parameters.AppStdout = appStdout
	return b
}

func (b *ServiceBuilder) AppStdoutCreationDisposition(appStdoutCreationDisposition string) *ServiceBuilder {
	b.service.Parameters.AppStdoutCreationDisposition = appStdoutCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStdoutFlagsAndAttributes(appStdoutFlagsAndAttributes string) *ServiceBuilder {
	b.service.Parameters.AppStdoutFlagsAndAttributes = appStdoutFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStdoutShareMode(appStdoutShareMode string) *ServiceBuilder {
	b.service.Parameters.AppStdoutShareMode = appStdoutShareMode
	return b
}

func (b *ServiceBuilder) AppStopMethodConsole(appStopMethodConsole string) *ServiceBuilder {
	b.service.Parameters.AppStopMethodConsole = appStopMethodConsole
	return b
}

func (b *ServiceBuilder) AppStopMethodSkip(appStopMethodSkip string) *ServiceBuilder {
	b.service.Parameters.AppStopMethodSkip = appStopMethodSkip
	return b
}

func (b *ServiceBuilder) AppStopMethodThreads(appStopMethodThreads string) *ServiceBuilder {
	b.service.Parameters.AppStopMethodThreads = appStopMethodThreads
	return b
}

func (b *ServiceBuilder) AppStopMethodWindow(appStopMethodWindow string) *ServiceBuilder {
	b.service.Parameters.AppStopMethodWindow = appStopMethodWindow
	return b
}

func (b *ServiceBuilder) AppThrottle(appThrottle string) *ServiceBuilder {
	b.service.Parameters.AppThrottle = appThrottle
	return b
}

func (b *ServiceBuilder) DependOnGroup(dependOnGroup string) *ServiceBuilder {
	b.service.Parameters.DependOnGroup = dependOnGroup
	return b
}

func (b *ServiceBuilder) DependOnService(dependOnService string) *ServiceBuilder {
	b.service.Parameters.DependOnService = dependOnService
	return b
}

func (b *ServiceBuilder) Description(description string) *ServiceBuilder {
	b.service.Parameters.Description = description
	return b
}

func (b *ServiceBuilder) DisplayName(displayName string) *ServiceBuilder {
	b.service.Parameters.DisplayName = displayName
	return b
}

func (b *ServiceBuilder) ImagePath(imagePath string) *ServiceBuilder {
	b.service.Parameters.ImagePath = imagePath
	return b
}

func (b *ServiceBuilder) Name(name string) *ServiceBuilder {
	b.service.Parameters.Name = name
	return b
}

func (b *ServiceBuilder) ObjectName(objectName string) *ServiceBuilder {
	b.service.Parameters.ObjectName = objectName
	return b
}

func (b *ServiceBuilder) Start(start string) *ServiceBuilder {
	b.service.Parameters.Start = start
	return b
}

func (b *ServiceBuilder) Type(serviceType string) *ServiceBuilder {
	b.service.Parameters.Type = serviceType
	return b
}

func (b *ServiceBuilder) Build() ServiceConfig {
	return b.service
}
