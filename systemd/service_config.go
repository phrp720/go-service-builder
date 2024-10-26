package systemd

// ServiceConfig is a struct that holds all the systemd service configuration
type ServiceConfig struct {
	unit Unit

	service Service

	install Install
}

// [Unit]
type Unit struct {
	Description        string
	Documentation      string
	After              string
	Wants              string
	Requires           string
	BindsTo            string
	PartOf             string
	Conflicts          string
	Before             string
	OnFailure          string
	StartLimitInterval string
	StartLimitBurst    int
}

// [Service]
type Service struct {
	ServiceType            string
	ExecStart              string
	ExecStartPre           string
	ExecStartPost          string
	ExecReload             string
	ExecStop               string
	ExecStopPost           string
	Restart                string
	RestartSec             string
	TimeoutSec             string
	Environment            string
	EnvironmentFile        string
	WorkingDirectory       string
	User                   string
	Group                  string
	PIDFile                string
	StandardOutput         string
	StandardError          string
	SyslogIdentifier       string
	SyslogFacility         string
	SyslogLevel            string
	SyslogLevelPrefix      bool
	CapabilityBoundingSet  string
	AmbientCapabilities    string
	NoNewPrivileges        bool
	PrivateTmp             bool
	ProtectSystem          string
	ProtectHome            string
	ReadOnlyPaths          string
	ReadWritePaths         string
	InaccessiblePaths      string
	RuntimeDirectory       string
	StateDirectory         string
	CacheDirectory         string
	LogsDirectory          string
	ConfigurationDirectory string
}

// [Install]
type Install struct {
	InstallWantedBy   string
	InstallRequiredBy string
}
