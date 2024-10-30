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
	b.service.serviceName = serviceName
	return b
}

func (b *ServiceBuilder) AppAffinity(appAffinity string) *ServiceBuilder {
	b.service.parameters.AppAffinity = appAffinity
	return b
}

func (b *ServiceBuilder) AppDirectory(appDirectory string) *ServiceBuilder {
	b.service.parameters.AppDirectory = appDirectory
	return b
}

func (b *ServiceBuilder) AppEnvironment(appEnvironment string) *ServiceBuilder {
	b.service.parameters.AppEnvironment = appEnvironment
	return b
}

func (b *ServiceBuilder) AppEnvironmentExtra(appEnvironmentExtra string) *ServiceBuilder {
	b.service.parameters.AppEnvironmentExtra = appEnvironmentExtra
	return b
}

func (b *ServiceBuilder) AppExit(appExit string) *ServiceBuilder {
	b.service.parameters.AppExit = appExit
	return b
}

func (b *ServiceBuilder) Application(application string) *ServiceBuilder {
	b.service.parameters.Application = application
	return b
}

func (b *ServiceBuilder) AppNoConsole(appNoConsole string) *ServiceBuilder {
	b.service.parameters.AppNoConsole = appNoConsole
	return b
}

func (b *ServiceBuilder) AppParameters(appParameters string) *ServiceBuilder {
	b.service.parameters.AppParameters = appParameters
	return b
}

func (b *ServiceBuilder) AppPriority(appPriority string) *ServiceBuilder {
	b.service.parameters.AppPriority = appPriority
	return b
}

func (b *ServiceBuilder) AppRestartDelay(appRestartDelay string) *ServiceBuilder {
	b.service.parameters.AppRestartDelay = appRestartDelay
	return b
}

func (b *ServiceBuilder) AppRotateBytes(appRotateBytes string) *ServiceBuilder {
	b.service.parameters.AppRotateBytes = appRotateBytes
	return b
}

func (b *ServiceBuilder) AppRotateBytesHigh(appRotateBytesHigh string) *ServiceBuilder {
	b.service.parameters.AppRotateBytesHigh = appRotateBytesHigh
	return b
}

func (b *ServiceBuilder) AppRotateFiles(appRotateFiles string) *ServiceBuilder {
	b.service.parameters.AppRotateFiles = appRotateFiles
	return b
}

func (b *ServiceBuilder) AppRotateOnline(appRotateOnline string) *ServiceBuilder {
	b.service.parameters.AppRotateOnline = appRotateOnline
	return b
}

func (b *ServiceBuilder) AppRotateSeconds(appRotateSeconds string) *ServiceBuilder {
	b.service.parameters.AppRotateSeconds = appRotateSeconds
	return b
}

func (b *ServiceBuilder) AppStderr(appStderr string) *ServiceBuilder {
	b.service.parameters.AppStderr = appStderr
	return b
}

func (b *ServiceBuilder) AppStderrCreationDisposition(appStderrCreationDisposition string) *ServiceBuilder {
	b.service.parameters.AppStderrCreationDisposition = appStderrCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStderrFlagsAndAttributes(appStderrFlagsAndAttributes string) *ServiceBuilder {
	b.service.parameters.AppStderrFlagsAndAttributes = appStderrFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStderrShareMode(appStderrShareMode string) *ServiceBuilder {
	b.service.parameters.AppStderrShareMode = appStderrShareMode
	return b
}

func (b *ServiceBuilder) AppStdin(appStdin string) *ServiceBuilder {
	b.service.parameters.AppStdin = appStdin
	return b
}

func (b *ServiceBuilder) AppStdinCreationDisposition(appStdinCreationDisposition string) *ServiceBuilder {
	b.service.parameters.AppStdinCreationDisposition = appStdinCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStdinFlagsAndAttributes(appStdinFlagsAndAttributes string) *ServiceBuilder {
	b.service.parameters.AppStdinFlagsAndAttributes = appStdinFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStdinShareMode(appStdinShareMode string) *ServiceBuilder {
	b.service.parameters.AppStdinShareMode = appStdinShareMode
	return b
}

func (b *ServiceBuilder) AppStdout(appStdout string) *ServiceBuilder {
	b.service.parameters.AppStdout = appStdout
	return b
}

func (b *ServiceBuilder) AppStdoutCreationDisposition(appStdoutCreationDisposition string) *ServiceBuilder {
	b.service.parameters.AppStdoutCreationDisposition = appStdoutCreationDisposition
	return b
}

func (b *ServiceBuilder) AppStdoutFlagsAndAttributes(appStdoutFlagsAndAttributes string) *ServiceBuilder {
	b.service.parameters.AppStdoutFlagsAndAttributes = appStdoutFlagsAndAttributes
	return b
}

func (b *ServiceBuilder) AppStdoutShareMode(appStdoutShareMode string) *ServiceBuilder {
	b.service.parameters.AppStdoutShareMode = appStdoutShareMode
	return b
}

func (b *ServiceBuilder) AppStopMethodConsole(appStopMethodConsole string) *ServiceBuilder {
	b.service.parameters.AppStopMethodConsole = appStopMethodConsole
	return b
}

func (b *ServiceBuilder) AppStopMethodSkip(appStopMethodSkip string) *ServiceBuilder {
	b.service.parameters.AppStopMethodSkip = appStopMethodSkip
	return b
}

func (b *ServiceBuilder) AppStopMethodThreads(appStopMethodThreads string) *ServiceBuilder {
	b.service.parameters.AppStopMethodThreads = appStopMethodThreads
	return b
}

func (b *ServiceBuilder) AppStopMethodWindow(appStopMethodWindow string) *ServiceBuilder {
	b.service.parameters.AppStopMethodWindow = appStopMethodWindow
	return b
}

func (b *ServiceBuilder) AppThrottle(appThrottle string) *ServiceBuilder {
	b.service.parameters.AppThrottle = appThrottle
	return b
}

func (b *ServiceBuilder) DependOnGroup(dependOnGroup string) *ServiceBuilder {
	b.service.parameters.DependOnGroup = dependOnGroup
	return b
}

func (b *ServiceBuilder) DependOnService(dependOnService string) *ServiceBuilder {
	b.service.parameters.DependOnService = dependOnService
	return b
}

func (b *ServiceBuilder) Description(description string) *ServiceBuilder {
	b.service.parameters.Description = description
	return b
}

func (b *ServiceBuilder) DisplayName(displayName string) *ServiceBuilder {
	b.service.parameters.DisplayName = displayName
	return b
}

func (b *ServiceBuilder) ImagePath(imagePath string) *ServiceBuilder {
	b.service.parameters.ImagePath = imagePath
	return b
}

func (b *ServiceBuilder) Name(name string) *ServiceBuilder {
	b.service.parameters.Name = name
	return b
}

func (b *ServiceBuilder) ObjectName(objectName string) *ServiceBuilder {
	b.service.parameters.ObjectName = objectName
	return b
}

func (b *ServiceBuilder) Start(start string) *ServiceBuilder {
	b.service.parameters.Start = start
	return b
}

func (b *ServiceBuilder) Type(serviceType string) *ServiceBuilder {
	b.service.parameters.Type = serviceType
	return b
}

func (b *ServiceBuilder) Build() ServiceConfig {
	return b.service
}
