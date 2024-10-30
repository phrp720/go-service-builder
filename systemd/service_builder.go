package systemd

// ServiceBuilder is a builder for creating ServiceConfig
type ServiceBuilder struct {
	systemd ServiceConfig
}

// NewServiceBuilder creates a new ServiceBuilder
func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		systemd: ServiceConfig{},
	}
}

// Unit

func (c *ServiceBuilder) Description(description string) *ServiceBuilder {
	c.systemd.Unit.Description = description
	return c
}

func (c *ServiceBuilder) Documentation(documentation string) *ServiceBuilder {
	c.systemd.Unit.Documentation = documentation
	return c
}

func (c *ServiceBuilder) Wants(wants string) *ServiceBuilder {
	c.systemd.Unit.Wants = wants
	return c
}

func (c *ServiceBuilder) Requires(requires string) *ServiceBuilder {
	c.systemd.Unit.Requires = requires
	return c
}

func (c *ServiceBuilder) Requisite(requisite string) *ServiceBuilder {
	c.systemd.Unit.Requisite = requisite
	return c
}

func (c *ServiceBuilder) BindsTo(bindsTo string) *ServiceBuilder {
	c.systemd.Unit.BindsTo = bindsTo
	return c
}

func (c *ServiceBuilder) PartOf(partOf string) *ServiceBuilder {
	c.systemd.Unit.PartOf = partOf
	return c
}

func (c *ServiceBuilder) Upholds(upholds string) *ServiceBuilder {
	c.systemd.Unit.Upholds = upholds
	return c
}

func (c *ServiceBuilder) Conflicts(conflicts string) *ServiceBuilder {
	c.systemd.Unit.Conflicts = conflicts
	return c
}

func (c *ServiceBuilder) Before(before string) *ServiceBuilder {
	c.systemd.Unit.Before = before
	return c
}

func (c *ServiceBuilder) After(after string) *ServiceBuilder {
	c.systemd.Unit.After = after
	return c
}

func (c *ServiceBuilder) OnFailure(onFailure string) *ServiceBuilder {
	c.systemd.Unit.OnFailure = onFailure
	return c
}

func (c *ServiceBuilder) OnSuccess(onSuccess string) *ServiceBuilder {
	c.systemd.Unit.OnSuccess = onSuccess
	return c
}

func (c *ServiceBuilder) PropagatesReloadTo(propagatesReloadTo string) *ServiceBuilder {
	c.systemd.Unit.PropagatesReloadTo = propagatesReloadTo
	return c
}

func (c *ServiceBuilder) ReloadPropagatedFrom(reloadPropagatedFrom string) *ServiceBuilder {
	c.systemd.Unit.ReloadPropagatedFrom = reloadPropagatedFrom
	return c
}

func (c *ServiceBuilder) PropagatesStopTo(propagatesStopTo string) *ServiceBuilder {
	c.systemd.Unit.PropagatesStopTo = propagatesStopTo
	return c
}

func (c *ServiceBuilder) StopPropagatedFrom(stopPropagatedFrom string) *ServiceBuilder {
	c.systemd.Unit.StopPropagatedFrom = stopPropagatedFrom
	return c
}

func (c *ServiceBuilder) JoinsNamespaceOf(joinsNamespaceOf string) *ServiceBuilder {
	c.systemd.Unit.JoinsNamespaceOf = joinsNamespaceOf
	return c
}

func (c *ServiceBuilder) RequiresMountsFor(requiresMountsFor string) *ServiceBuilder {
	c.systemd.Unit.RequiresMountsFor = requiresMountsFor
	return c
}

func (c *ServiceBuilder) WantsMountsFor(wantsMountsFor string) *ServiceBuilder {
	c.systemd.Unit.WantsMountsFor = wantsMountsFor
	return c
}

func (c *ServiceBuilder) OnSuccessJobMode(onSuccessJobMode string) *ServiceBuilder {
	c.systemd.Unit.OnSuccessJobMode = onSuccessJobMode
	return c
}

func (c *ServiceBuilder) OnFailureJobMode(onFailureJobMode string) *ServiceBuilder {
	c.systemd.Unit.OnFailureJobMode = onFailureJobMode
	return c
}

func (c *ServiceBuilder) IgnoreOnIsolate(ignoreOnIsolate bool) *ServiceBuilder {
	c.systemd.Unit.IgnoreOnIsolate = ignoreOnIsolate
	return c
}

func (c *ServiceBuilder) StopWhenUnneeded(stopWhenUnneeded bool) *ServiceBuilder {
	c.systemd.Unit.StopWhenUnneeded = stopWhenUnneeded
	return c
}

func (c *ServiceBuilder) RefuseManualStart(refuseManualStart bool) *ServiceBuilder {
	c.systemd.Unit.RefuseManualStart = refuseManualStart
	return c
}

func (c *ServiceBuilder) RefuseManualStop(refuseManualStop bool) *ServiceBuilder {
	c.systemd.Unit.RefuseManualStop = refuseManualStop
	return c
}

func (c *ServiceBuilder) AllowIsolate(allowIsolate bool) *ServiceBuilder {
	c.systemd.Unit.AllowIsolate = allowIsolate
	return c
}

func (c *ServiceBuilder) DefaultDependencies(defaultDependencies bool) *ServiceBuilder {
	c.systemd.Unit.DefaultDependencies = defaultDependencies
	return c
}

func (c *ServiceBuilder) SurviveFinalKillSignal(surviveFinalKillSignal bool) *ServiceBuilder {
	c.systemd.Unit.SurviveFinalKillSignal = surviveFinalKillSignal
	return c
}

func (c *ServiceBuilder) CollectMode(collectMode string) *ServiceBuilder {
	c.systemd.Unit.CollectMode = collectMode
	return c
}

func (c *ServiceBuilder) FailureAction(failureAction string) *ServiceBuilder {
	c.systemd.Unit.FailureAction = failureAction
	return c
}

func (c *ServiceBuilder) SuccessAction(successAction string) *ServiceBuilder {
	c.systemd.Unit.SuccessAction = successAction
	return c
}

func (c *ServiceBuilder) FailureActionExitStatus(failureActionExitStatus int) *ServiceBuilder {
	c.systemd.Unit.FailureActionExitStatus = failureActionExitStatus
	return c
}

func (c *ServiceBuilder) SuccessActionExitStatus(successActionExitStatus int) *ServiceBuilder {
	c.systemd.Unit.SuccessActionExitStatus = successActionExitStatus
	return c
}

func (c *ServiceBuilder) JobTimeoutSec(jobTimeoutSec string) *ServiceBuilder {
	c.systemd.Unit.JobTimeoutSec = jobTimeoutSec
	return c
}

func (c *ServiceBuilder) JobRunningTimeoutSec(jobRunningTimeoutSec string) *ServiceBuilder {
	c.systemd.Unit.JobRunningTimeoutSec = jobRunningTimeoutSec
	return c
}

func (c *ServiceBuilder) JobTimeoutAction(jobTimeoutAction string) *ServiceBuilder {
	c.systemd.Unit.JobTimeoutAction = jobTimeoutAction
	return c
}

func (c *ServiceBuilder) JobTimeoutRebootArgument(jobTimeoutRebootArgument string) *ServiceBuilder {
	c.systemd.Unit.JobTimeoutRebootArgument = jobTimeoutRebootArgument
	return c
}

func (c *ServiceBuilder) StartLimitIntervalSec(startLimitIntervalSec string) *ServiceBuilder {
	c.systemd.Unit.StartLimitIntervalSec = startLimitIntervalSec
	return c
}

func (c *ServiceBuilder) StartLimitBurst(startLimitBurst int) *ServiceBuilder {
	c.systemd.Unit.StartLimitBurst = startLimitBurst
	return c
}

func (c *ServiceBuilder) StartLimitAction(startLimitAction string) *ServiceBuilder {
	c.systemd.Unit.StartLimitAction = startLimitAction
	return c
}

func (c *ServiceBuilder) RebootArgument(rebootArgument string) *ServiceBuilder {
	c.systemd.Unit.RebootArgument = rebootArgument
	return c
}

func (c *ServiceBuilder) SourcePath(sourcePath string) *ServiceBuilder {
	c.systemd.Unit.SourcePath = sourcePath
	return c
}

// Service

func (c *ServiceBuilder) ServiceType(ServiceType string) *ServiceBuilder {
	c.systemd.Service.Type = ServiceType
	return c
}

func (c *ServiceBuilder) BusName(busName string) *ServiceBuilder {
	c.systemd.Service.BusName = busName
	return c
}

func (c *ServiceBuilder) ExecStart(execStart string) *ServiceBuilder {
	c.systemd.Service.ExecStart = execStart
	return c
}

func (c *ServiceBuilder) ExecStartPre(execStartPre string) *ServiceBuilder {
	c.systemd.Service.ExecStartPre = execStartPre
	return c
}

func (c *ServiceBuilder) ExecStartPost(execStartPost string) *ServiceBuilder {
	c.systemd.Service.ExecStartPost = execStartPost
	return c
}

func (c *ServiceBuilder) ExecReload(execReload string) *ServiceBuilder {
	c.systemd.Service.ExecReload = execReload
	return c
}

func (c *ServiceBuilder) ExecStop(execStop string) *ServiceBuilder {
	c.systemd.Service.ExecStop = execStop
	return c
}

func (c *ServiceBuilder) ExecStopPost(execStopPost string) *ServiceBuilder {
	c.systemd.Service.ExecStopPost = execStopPost
	return c
}

func (c *ServiceBuilder) Restart(restart string) *ServiceBuilder {
	c.systemd.Service.Restart = restart
	return c
}

func (c *ServiceBuilder) RestartSec(restartSec string) *ServiceBuilder {
	c.systemd.Service.RestartSec = restartSec
	return c
}

func (c *ServiceBuilder) RemainAfterExit(remainAfterExit bool) *ServiceBuilder {
	c.systemd.Service.RemainAfterExit = remainAfterExit
	return c
}

func (c *ServiceBuilder) TimeoutSec(timeoutSec string) *ServiceBuilder {
	c.systemd.Service.TimeoutSec = timeoutSec
	return c
}

func (c *ServiceBuilder) Environment(environment string) *ServiceBuilder {
	c.systemd.Service.Environment = environment
	return c
}

func (c *ServiceBuilder) EnvironmentFile(environmentFile string) *ServiceBuilder {
	c.systemd.Service.EnvironmentFile = environmentFile
	return c
}

func (c *ServiceBuilder) WorkingDirectory(workingDirectory string) *ServiceBuilder {
	c.systemd.Service.WorkingDirectory = workingDirectory
	return c
}

func (c *ServiceBuilder) User(user string) *ServiceBuilder {
	c.systemd.Service.User = user
	return c
}

func (c *ServiceBuilder) Group(group string) *ServiceBuilder {
	c.systemd.Service.Group = group
	return c
}

func (c *ServiceBuilder) PIDFile(pidFile string) *ServiceBuilder {
	c.systemd.Service.PIDFile = pidFile
	return c
}

func (c *ServiceBuilder) StandardOutput(standardOutput string) *ServiceBuilder {
	c.systemd.Service.StandardOutput = standardOutput
	return c
}

func (c *ServiceBuilder) StandardError(standardError string) *ServiceBuilder {
	c.systemd.Service.StandardError = standardError
	return c
}

func (c *ServiceBuilder) SyslogIdentifier(syslogIdentifier string) *ServiceBuilder {
	c.systemd.Service.SyslogIdentifier = syslogIdentifier
	return c
}

func (c *ServiceBuilder) SyslogFacility(syslogFacility string) *ServiceBuilder {
	c.systemd.Service.SyslogFacility = syslogFacility
	return c
}

func (c *ServiceBuilder) SyslogLevel(syslogLevel string) *ServiceBuilder {
	c.systemd.Service.SyslogLevel = syslogLevel
	return c
}

func (c *ServiceBuilder) SyslogLevelPrefix(syslogLevelPrefix bool) *ServiceBuilder {
	c.systemd.Service.SyslogLevelPrefix = syslogLevelPrefix
	return c
}

func (c *ServiceBuilder) CapabilityBoundingSet(capabilityBoundingSet string) *ServiceBuilder {
	c.systemd.Service.CapabilityBoundingSet = capabilityBoundingSet
	return c
}

func (c *ServiceBuilder) AmbientCapabilities(ambientCapabilities string) *ServiceBuilder {
	c.systemd.Service.AmbientCapabilities = ambientCapabilities
	return c
}

func (c *ServiceBuilder) NoNewPrivileges(noNewPrivileges bool) *ServiceBuilder {
	c.systemd.Service.NoNewPrivileges = noNewPrivileges
	return c
}

func (c *ServiceBuilder) NotifyAccess(notifyAccess string) *ServiceBuilder {
	c.systemd.Service.NotifyAccess = notifyAccess
	return c
}

func (c *ServiceBuilder) PrivateTmp(privateTmp bool) *ServiceBuilder {
	c.systemd.Service.PrivateTmp = privateTmp
	return c
}

func (c *ServiceBuilder) ProtectSystem(protectSystem string) *ServiceBuilder {
	c.systemd.Service.ProtectSystem = protectSystem
	return c
}

func (c *ServiceBuilder) ProtectHome(protectHome string) *ServiceBuilder {
	c.systemd.Service.ProtectHome = protectHome
	return c
}

func (c *ServiceBuilder) ReadOnlyPaths(readOnlyPaths string) *ServiceBuilder {
	c.systemd.Service.ReadOnlyPaths = readOnlyPaths
	return c
}

func (c *ServiceBuilder) ReadWritePaths(readWritePaths string) *ServiceBuilder {
	c.systemd.Service.ReadWritePaths = readWritePaths
	return c
}

func (c *ServiceBuilder) InaccessiblePaths(inaccessiblePaths string) *ServiceBuilder {
	c.systemd.Service.InaccessiblePaths = inaccessiblePaths
	return c
}

func (c *ServiceBuilder) RuntimeDirectory(runtimeDirectory string) *ServiceBuilder {
	c.systemd.Service.RuntimeDirectory = runtimeDirectory
	return c
}

func (c *ServiceBuilder) StateDirectory(stateDirectory string) *ServiceBuilder {
	c.systemd.Service.StateDirectory = stateDirectory
	return c
}

func (c *ServiceBuilder) CacheDirectory(cacheDirectory string) *ServiceBuilder {
	c.systemd.Service.CacheDirectory = cacheDirectory
	return c
}

func (c *ServiceBuilder) LogsDirectory(logsDirectory string) *ServiceBuilder {
	c.systemd.Service.LogsDirectory = logsDirectory
	return c
}

func (c *ServiceBuilder) ConfigurationDirectory(configurationDirectory string) *ServiceBuilder {
	c.systemd.Service.ConfigurationDirectory = configurationDirectory
	return c
}

func (c *ServiceBuilder) ExitType(exitType string) *ServiceBuilder {
	c.systemd.Service.ExitType = exitType
	return c
}

func (c *ServiceBuilder) GuessMainPID(guessMainPID bool) *ServiceBuilder {
	c.systemd.Service.GuessMainPID = guessMainPID
	return c
}

func (c *ServiceBuilder) RestartMode(restartMode string) *ServiceBuilder {
	c.systemd.Service.RestartMode = restartMode
	return c
}

func (c *ServiceBuilder) SuccessExitStatus(successExitStatus string) *ServiceBuilder {
	c.systemd.Service.SuccessExitStatus = successExitStatus
	return c
}

func (c *ServiceBuilder) RestartPreventExitStatus(restartPreventExitStatus string) *ServiceBuilder {
	c.systemd.Service.RestartPreventExitStatus = restartPreventExitStatus
	return c
}

func (c *ServiceBuilder) RestartForceExitStatus(restartForceExitStatus string) *ServiceBuilder {
	c.systemd.Service.RestartForceExitStatus = restartForceExitStatus
	return c
}

func (c *ServiceBuilder) RootDirectoryStartOnly(rootDirectoryStartOnly bool) *ServiceBuilder {
	c.systemd.Service.RootDirectoryStartOnly = rootDirectoryStartOnly
	return c
}

func (c *ServiceBuilder) NonBlocking(nonBlocking bool) *ServiceBuilder {
	c.systemd.Service.NonBlocking = nonBlocking
	return c
}

func (c *ServiceBuilder) Sockets(sockets string) *ServiceBuilder {
	c.systemd.Service.Sockets = sockets
	return c
}

func (c *ServiceBuilder) FileDescriptorStoreMax(fileDescriptorStoreMax int) *ServiceBuilder {
	c.systemd.Service.FileDescriptorStoreMax = fileDescriptorStoreMax
	return c
}

func (c *ServiceBuilder) FileDescriptorStorePreserve(fileDescriptorStorePreserve string) *ServiceBuilder {
	c.systemd.Service.FileDescriptorStorePreserve = fileDescriptorStorePreserve
	return c
}

func (c *ServiceBuilder) USBFunctionDescriptors(usbFunctionDescriptors string) *ServiceBuilder {
	c.systemd.Service.USBFunctionDescriptors = usbFunctionDescriptors
	return c
}

func (c *ServiceBuilder) USBFunctionStrings(usbFunctionStrings string) *ServiceBuilder {
	c.systemd.Service.USBFunctionStrings = usbFunctionStrings
	return c
}

func (c *ServiceBuilder) OOMPolicy(oomPolicy string) *ServiceBuilder {
	c.systemd.Service.OOMPolicy = oomPolicy
	return c
}

func (c *ServiceBuilder) OpenFile(openFile string) *ServiceBuilder {
	c.systemd.Service.OpenFile = openFile
	return c
}

func (c *ServiceBuilder) ReloadSignal(reloadSignal string) *ServiceBuilder {
	c.systemd.Service.ReloadSignal = reloadSignal
	return c
}

// Install

func (c *ServiceBuilder) WantedBy(wantedBy string) *ServiceBuilder {
	c.systemd.Install.WantedBy = wantedBy
	return c
}

func (c *ServiceBuilder) RequiredBy(requiredBy string) *ServiceBuilder {
	c.systemd.Install.RequiredBy = requiredBy
	return c
}

func (c *ServiceBuilder) UpheldBy(upheldBy string) *ServiceBuilder {
	c.systemd.Install.UpheldBy = upheldBy
	return c
}

func (c *ServiceBuilder) Alias(alias string) *ServiceBuilder {
	c.systemd.Install.Alias = alias
	return c
}

func (c *ServiceBuilder) Also(also string) *ServiceBuilder {
	c.systemd.Install.Also = also
	return c
}

func (c *ServiceBuilder) DefaultInstance(defaultInstance string) *ServiceBuilder {
	c.systemd.Install.DefaultInstance = defaultInstance
	return c
}

func (c *ServiceBuilder) Build() ServiceConfig {
	return c.systemd
}
