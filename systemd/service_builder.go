package systemd

type ServiceBuilder struct {
	systemd ServiceConfig
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{systemd: ServiceConfig{}}
}

func (c *ServiceBuilder) Description(description string) *ServiceBuilder {
	c.systemd.unit.Description = description
	return c
}

func (c *ServiceBuilder) Documentation(doc string) *ServiceBuilder {
	c.systemd.unit.Documentation = doc
	return c
}

func (c *ServiceBuilder) After(after string) *ServiceBuilder {
	c.systemd.unit.After = after
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

func (c *ServiceBuilder) BindsTo(bindsTo string) *ServiceBuilder {
	c.systemd.unit.BindsTo = bindsTo
	return c
}

func (c *ServiceBuilder) PartOf(partOf string) *ServiceBuilder {
	c.systemd.unit.PartOf = partOf
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

func (c *ServiceBuilder) OnFailure(onFailure string) *ServiceBuilder {
	c.systemd.unit.OnFailure = onFailure
	return c
}

func (c *ServiceBuilder) StartLimitInterval(interval string) *ServiceBuilder {
	c.systemd.unit.StartLimitInterval = interval
	return c
}

func (c *ServiceBuilder) StartLimitBurst(burst int) *ServiceBuilder {
	c.systemd.unit.StartLimitBurst = burst
	return c
}

func (c *ServiceBuilder) ServiceType(serviceType string) *ServiceBuilder {
	c.systemd.service.ServiceType = serviceType
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

func (c *ServiceBuilder) InstallWantedBy(installWantedBy string) *ServiceBuilder {
	c.systemd.install.InstallWantedBy = installWantedBy
	return c
}

func (c *ServiceBuilder) InstallRequiredBy(installRequiredBy string) *ServiceBuilder {
	c.systemd.install.InstallRequiredBy = installRequiredBy
	return c
}

func (c *ServiceBuilder) Build() ServiceConfig {
	return c.systemd
}
