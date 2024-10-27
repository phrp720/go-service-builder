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
	c.systemd.unit.Description = description
	return c
}

func (c *ServiceBuilder) Documentation(documentation string) *ServiceBuilder {
	c.systemd.unit.Documentation = documentation
	return c
}

func (c *ServiceBuilder) Wants(wants string) *ServiceBuilder {
	c.systemd.unit.Wants = wants
	return c
}

func (c *ServiceBuilder) Requires(requires string) *ServiceBuilder {
	c.systemd.unit.Requires = requires
	return c
}

func (c *ServiceBuilder) Requisite(requisite string) *ServiceBuilder {
	c.systemd.unit.Requisite = requisite
	return c
}

func (c *ServiceBuilder) BindsTo(bindsTo string) *ServiceBuilder {
	c.systemd.unit.BindsTo = bindsTo
	return c
}

func (c *ServiceBuilder) PartOf(partOf string) *ServiceBuilder {
	c.systemd.unit.PartOf = partOf
	return c
}

func (c *ServiceBuilder) Upholds(upholds string) *ServiceBuilder {
	c.systemd.unit.Upholds = upholds
	return c
}

func (c *ServiceBuilder) Conflicts(conflicts string) *ServiceBuilder {
	c.systemd.unit.Conflicts = conflicts
	return c
}

func (c *ServiceBuilder) Before(before string) *ServiceBuilder {
	c.systemd.unit.Before = before
	return c
}

func (c *ServiceBuilder) After(after string) *ServiceBuilder {
	c.systemd.unit.After = after
	return c
}

func (c *ServiceBuilder) OnFailure(onFailure string) *ServiceBuilder {
	c.systemd.unit.OnFailure = onFailure
	return c
}

func (c *ServiceBuilder) OnSuccess(onSuccess string) *ServiceBuilder {
	c.systemd.unit.OnSuccess = onSuccess
	return c
}

func (c *ServiceBuilder) PropagatesReloadTo(propagatesReloadTo string) *ServiceBuilder {
	c.systemd.unit.PropagatesReloadTo = propagatesReloadTo
	return c
}

func (c *ServiceBuilder) ReloadPropagatedFrom(reloadPropagatedFrom string) *ServiceBuilder {
	c.systemd.unit.ReloadPropagatedFrom = reloadPropagatedFrom
	return c
}

func (c *ServiceBuilder) PropagatesStopTo(propagatesStopTo string) *ServiceBuilder {
	c.systemd.unit.PropagatesStopTo = propagatesStopTo
	return c
}

func (c *ServiceBuilder) StopPropagatedFrom(stopPropagatedFrom string) *ServiceBuilder {
	c.systemd.unit.StopPropagatedFrom = stopPropagatedFrom
	return c
}

func (c *ServiceBuilder) JoinsNamespaceOf(joinsNamespaceOf string) *ServiceBuilder {
	c.systemd.unit.JoinsNamespaceOf = joinsNamespaceOf
	return c
}

func (c *ServiceBuilder) RequiresMountsFor(requiresMountsFor string) *ServiceBuilder {
	c.systemd.unit.RequiresMountsFor = requiresMountsFor
	return c
}

func (c *ServiceBuilder) WantsMountsFor(wantsMountsFor string) *ServiceBuilder {
	c.systemd.unit.WantsMountsFor = wantsMountsFor
	return c
}

func (c *ServiceBuilder) OnSuccessJobMode(onSuccessJobMode string) *ServiceBuilder {
	c.systemd.unit.OnSuccessJobMode = onSuccessJobMode
	return c
}

func (c *ServiceBuilder) OnFailureJobMode(onFailureJobMode string) *ServiceBuilder {
	c.systemd.unit.OnFailureJobMode = onFailureJobMode
	return c
}

func (c *ServiceBuilder) IgnoreOnIsolate(ignoreOnIsolate bool) *ServiceBuilder {
	c.systemd.unit.IgnoreOnIsolate = ignoreOnIsolate
	return c
}

func (c *ServiceBuilder) StopWhenUnneeded(stopWhenUnneeded bool) *ServiceBuilder {
	c.systemd.unit.StopWhenUnneeded = stopWhenUnneeded
	return c
}

func (c *ServiceBuilder) RefuseManualStart(refuseManualStart bool) *ServiceBuilder {
	c.systemd.unit.RefuseManualStart = refuseManualStart
	return c
}

func (c *ServiceBuilder) RefuseManualStop(refuseManualStop bool) *ServiceBuilder {
	c.systemd.unit.RefuseManualStop = refuseManualStop
	return c
}

func (c *ServiceBuilder) AllowIsolate(allowIsolate bool) *ServiceBuilder {
	c.systemd.unit.AllowIsolate = allowIsolate
	return c
}

func (c *ServiceBuilder) DefaultDependencies(defaultDependencies bool) *ServiceBuilder {
	c.systemd.unit.DefaultDependencies = defaultDependencies
	return c
}

func (c *ServiceBuilder) SurviveFinalKillSignal(surviveFinalKillSignal bool) *ServiceBuilder {
	c.systemd.unit.SurviveFinalKillSignal = surviveFinalKillSignal
	return c
}

func (c *ServiceBuilder) CollectMode(collectMode string) *ServiceBuilder {
	c.systemd.unit.CollectMode = collectMode
	return c
}

func (c *ServiceBuilder) FailureAction(failureAction string) *ServiceBuilder {
	c.systemd.unit.FailureAction = failureAction
	return c
}

func (c *ServiceBuilder) SuccessAction(successAction string) *ServiceBuilder {
	c.systemd.unit.SuccessAction = successAction
	return c
}

func (c *ServiceBuilder) FailureActionExitStatus(failureActionExitStatus int) *ServiceBuilder {
	c.systemd.unit.FailureActionExitStatus = failureActionExitStatus
	return c
}

func (c *ServiceBuilder) SuccessActionExitStatus(successActionExitStatus int) *ServiceBuilder {
	c.systemd.unit.SuccessActionExitStatus = successActionExitStatus
	return c
}

func (c *ServiceBuilder) JobTimeoutSec(jobTimeoutSec string) *ServiceBuilder {
	c.systemd.unit.JobTimeoutSec = jobTimeoutSec
	return c
}

func (c *ServiceBuilder) JobRunningTimeoutSec(jobRunningTimeoutSec string) *ServiceBuilder {
	c.systemd.unit.JobRunningTimeoutSec = jobRunningTimeoutSec
	return c
}

func (c *ServiceBuilder) JobTimeoutAction(jobTimeoutAction string) *ServiceBuilder {
	c.systemd.unit.JobTimeoutAction = jobTimeoutAction
	return c
}

func (c *ServiceBuilder) JobTimeoutRebootArgument(jobTimeoutRebootArgument string) *ServiceBuilder {
	c.systemd.unit.JobTimeoutRebootArgument = jobTimeoutRebootArgument
	return c
}

func (c *ServiceBuilder) StartLimitIntervalSec(startLimitIntervalSec string) *ServiceBuilder {
	c.systemd.unit.StartLimitIntervalSec = startLimitIntervalSec
	return c
}

func (c *ServiceBuilder) StartLimitBurst(startLimitBurst int) *ServiceBuilder {
	c.systemd.unit.StartLimitBurst = startLimitBurst
	return c
}

func (c *ServiceBuilder) StartLimitAction(startLimitAction string) *ServiceBuilder {
	c.systemd.unit.StartLimitAction = startLimitAction
	return c
}

func (c *ServiceBuilder) RebootArgument(rebootArgument string) *ServiceBuilder {
	c.systemd.unit.RebootArgument = rebootArgument
	return c
}

func (c *ServiceBuilder) SourcePath(sourcePath string) *ServiceBuilder {
	c.systemd.unit.SourcePath = sourcePath
	return c
}

// Service

func (c *ServiceBuilder) ServiceType(serviceType string) *ServiceBuilder {
	c.systemd.service.Type = serviceType
	return c
}

func (c *ServiceBuilder) BusName(busName string) *ServiceBuilder {
	c.systemd.service.BusName = busName
	return c
}

func (c *ServiceBuilder) ExecStart(execStart string) *ServiceBuilder {
	c.systemd.service.ExecStart = execStart
	return c
}

func (c *ServiceBuilder) ExecStartPre(execStartPre string) *ServiceBuilder {
	c.systemd.service.ExecStartPre = execStartPre
	return c
}

func (c *ServiceBuilder) ExecStartPost(execStartPost string) *ServiceBuilder {
	c.systemd.service.ExecStartPost = execStartPost
	return c
}

func (c *ServiceBuilder) ExecReload(execReload string) *ServiceBuilder {
	c.systemd.service.ExecReload = execReload
	return c
}

func (c *ServiceBuilder) ExecStop(execStop string) *ServiceBuilder {
	c.systemd.service.ExecStop = execStop
	return c
}

func (c *ServiceBuilder) ExecStopPost(execStopPost string) *ServiceBuilder {
	c.systemd.service.ExecStopPost = execStopPost
	return c
}

func (c *ServiceBuilder) Restart(restart string) *ServiceBuilder {
	c.systemd.service.Restart = restart
	return c
}

func (c *ServiceBuilder) RestartSec(restartSec string) *ServiceBuilder {
	c.systemd.service.RestartSec = restartSec
	return c
}

func (c *ServiceBuilder) RemainAfterExit(remainAfterExit bool) *ServiceBuilder {
	c.systemd.service.RemainAfterExit = remainAfterExit
	return c
}

func (c *ServiceBuilder) TimeoutSec(timeoutSec string) *ServiceBuilder {
	c.systemd.service.TimeoutSec = timeoutSec
	return c
}

func (c *ServiceBuilder) Environment(environment string) *ServiceBuilder {
	c.systemd.service.Environment = environment
	return c
}

func (c *ServiceBuilder) EnvironmentFile(environmentFile string) *ServiceBuilder {
	c.systemd.service.EnvironmentFile = environmentFile
	return c
}

func (c *ServiceBuilder) WorkingDirectory(workingDirectory string) *ServiceBuilder {
	c.systemd.service.WorkingDirectory = workingDirectory
	return c
}

func (c *ServiceBuilder) User(user string) *ServiceBuilder {
	c.systemd.service.User = user
	return c
}

func (c *ServiceBuilder) Group(group string) *ServiceBuilder {
	c.systemd.service.Group = group
	return c
}

func (c *ServiceBuilder) PIDFile(pidFile string) *ServiceBuilder {
	c.systemd.service.PIDFile = pidFile
	return c
}

func (c *ServiceBuilder) StandardOutput(standardOutput string) *ServiceBuilder {
	c.systemd.service.StandardOutput = standardOutput
	return c
}

func (c *ServiceBuilder) StandardError(standardError string) *ServiceBuilder {
	c.systemd.service.StandardError = standardError
	return c
}

func (c *ServiceBuilder) SyslogIdentifier(syslogIdentifier string) *ServiceBuilder {
	c.systemd.service.SyslogIdentifier = syslogIdentifier
	return c
}

func (c *ServiceBuilder) SyslogFacility(syslogFacility string) *ServiceBuilder {
	c.systemd.service.SyslogFacility = syslogFacility
	return c
}

func (c *ServiceBuilder) SyslogLevel(syslogLevel string) *ServiceBuilder {
	c.systemd.service.SyslogLevel = syslogLevel
	return c
}

func (c *ServiceBuilder) SyslogLevelPrefix(syslogLevelPrefix bool) *ServiceBuilder {
	c.systemd.service.SyslogLevelPrefix = syslogLevelPrefix
	return c
}

func (c *ServiceBuilder) CapabilityBoundingSet(capabilityBoundingSet string) *ServiceBuilder {
	c.systemd.service.CapabilityBoundingSet = capabilityBoundingSet
	return c
}

func (c *ServiceBuilder) AmbientCapabilities(ambientCapabilities string) *ServiceBuilder {
	c.systemd.service.AmbientCapabilities = ambientCapabilities
	return c
}

func (c *ServiceBuilder) NoNewPrivileges(noNewPrivileges bool) *ServiceBuilder {
	c.systemd.service.NoNewPrivileges = noNewPrivileges
	return c
}

func (c *ServiceBuilder) NotifyAccess(notifyAccess string) *ServiceBuilder {
	c.systemd.service.NotifyAccess = notifyAccess
	return c
}

func (c *ServiceBuilder) PrivateTmp(privateTmp bool) *ServiceBuilder {
	c.systemd.service.PrivateTmp = privateTmp
	return c
}

func (c *ServiceBuilder) ProtectSystem(protectSystem string) *ServiceBuilder {
	c.systemd.service.ProtectSystem = protectSystem
	return c
}

func (c *ServiceBuilder) ProtectHome(protectHome string) *ServiceBuilder {
	c.systemd.service.ProtectHome = protectHome
	return c
}

func (c *ServiceBuilder) ReadOnlyPaths(readOnlyPaths string) *ServiceBuilder {
	c.systemd.service.ReadOnlyPaths = readOnlyPaths
	return c
}

func (c *ServiceBuilder) ReadWritePaths(readWritePaths string) *ServiceBuilder {
	c.systemd.service.ReadWritePaths = readWritePaths
	return c
}

func (c *ServiceBuilder) InaccessiblePaths(inaccessiblePaths string) *ServiceBuilder {
	c.systemd.service.InaccessiblePaths = inaccessiblePaths
	return c
}

func (c *ServiceBuilder) RuntimeDirectory(runtimeDirectory string) *ServiceBuilder {
	c.systemd.service.RuntimeDirectory = runtimeDirectory
	return c
}

func (c *ServiceBuilder) StateDirectory(stateDirectory string) *ServiceBuilder {
	c.systemd.service.StateDirectory = stateDirectory
	return c
}

func (c *ServiceBuilder) CacheDirectory(cacheDirectory string) *ServiceBuilder {
	c.systemd.service.CacheDirectory = cacheDirectory
	return c
}

func (c *ServiceBuilder) LogsDirectory(logsDirectory string) *ServiceBuilder {
	c.systemd.service.LogsDirectory = logsDirectory
	return c
}

func (c *ServiceBuilder) ConfigurationDirectory(configurationDirectory string) *ServiceBuilder {
	c.systemd.service.ConfigurationDirectory = configurationDirectory
	return c
}

func (c *ServiceBuilder) ExitType(exitType string) *ServiceBuilder {
	c.systemd.service.ExitType = exitType
	return c
}

func (c *ServiceBuilder) GuessMainPID(guessMainPID bool) *ServiceBuilder {
	c.systemd.service.GuessMainPID = guessMainPID
	return c
}

func (c *ServiceBuilder) RestartMode(restartMode string) *ServiceBuilder {
	c.systemd.service.RestartMode = restartMode
	return c
}

func (c *ServiceBuilder) SuccessExitStatus(successExitStatus string) *ServiceBuilder {
	c.systemd.service.SuccessExitStatus = successExitStatus
	return c
}

func (c *ServiceBuilder) RestartPreventExitStatus(restartPreventExitStatus string) *ServiceBuilder {
	c.systemd.service.RestartPreventExitStatus = restartPreventExitStatus
	return c
}

func (c *ServiceBuilder) RestartForceExitStatus(restartForceExitStatus string) *ServiceBuilder {
	c.systemd.service.RestartForceExitStatus = restartForceExitStatus
	return c
}

func (c *ServiceBuilder) RootDirectoryStartOnly(rootDirectoryStartOnly bool) *ServiceBuilder {
	c.systemd.service.RootDirectoryStartOnly = rootDirectoryStartOnly
	return c
}

func (c *ServiceBuilder) NonBlocking(nonBlocking bool) *ServiceBuilder {
	c.systemd.service.NonBlocking = nonBlocking
	return c
}

func (c *ServiceBuilder) Sockets(sockets string) *ServiceBuilder {
	c.systemd.service.Sockets = sockets
	return c
}

func (c *ServiceBuilder) FileDescriptorStoreMax(fileDescriptorStoreMax int) *ServiceBuilder {
	c.systemd.service.FileDescriptorStoreMax = fileDescriptorStoreMax
	return c
}

func (c *ServiceBuilder) FileDescriptorStorePreserve(fileDescriptorStorePreserve string) *ServiceBuilder {
	c.systemd.service.FileDescriptorStorePreserve = fileDescriptorStorePreserve
	return c
}

func (c *ServiceBuilder) USBFunctionDescriptors(usbFunctionDescriptors string) *ServiceBuilder {
	c.systemd.service.USBFunctionDescriptors = usbFunctionDescriptors
	return c
}

func (c *ServiceBuilder) USBFunctionStrings(usbFunctionStrings string) *ServiceBuilder {
	c.systemd.service.USBFunctionStrings = usbFunctionStrings
	return c
}

func (c *ServiceBuilder) OOMPolicy(oomPolicy string) *ServiceBuilder {
	c.systemd.service.OOMPolicy = oomPolicy
	return c
}

func (c *ServiceBuilder) OpenFile(openFile string) *ServiceBuilder {
	c.systemd.service.OpenFile = openFile
	return c
}

func (c *ServiceBuilder) ReloadSignal(reloadSignal string) *ServiceBuilder {
	c.systemd.service.ReloadSignal = reloadSignal
	return c
}

// Install

func (c *ServiceBuilder) WantedBy(wantedBy string) *ServiceBuilder {
	c.systemd.install.WantedBy = wantedBy
	return c
}

func (c *ServiceBuilder) RequiredBy(requiredBy string) *ServiceBuilder {
	c.systemd.install.RequiredBy = requiredBy
	return c
}

func (c *ServiceBuilder) UpheldBy(upheldBy string) *ServiceBuilder {
	c.systemd.install.UpheldBy = upheldBy
	return c
}

func (c *ServiceBuilder) Alias(alias string) *ServiceBuilder {
	c.systemd.install.Alias = alias
	return c
}

func (c *ServiceBuilder) Also(also string) *ServiceBuilder {
	c.systemd.install.Also = also
	return c
}

func (c *ServiceBuilder) DefaultInstance(defaultInstance string) *ServiceBuilder {
	c.systemd.install.DefaultInstance = defaultInstance
	return c
}

func (c *ServiceBuilder) Build() ServiceConfig {
	return c.systemd
}
