package systemd

// ServiceConfig is a struct that holds all the systemd service configuration
type ServiceConfig struct {
	unit Unit

	service Service

	install Install
}

// Unit holds the [Unit] section options for a systemd service
type Unit struct {
	Description              string
	Documentation            string
	Wants                    string
	Requires                 string
	Requisite                string
	BindsTo                  string
	PartOf                   string
	Upholds                  string
	Conflicts                string
	Before                   string
	After                    string
	OnFailure                string
	OnSuccess                string
	PropagatesReloadTo       string
	ReloadPropagatedFrom     string
	PropagatesStopTo         string
	StopPropagatedFrom       string
	JoinsNamespaceOf         string
	RequiresMountsFor        string
	WantsMountsFor           string
	OnSuccessJobMode         string
	OnFailureJobMode         string
	IgnoreOnIsolate          bool
	StopWhenUnneeded         bool
	RefuseManualStart        bool
	RefuseManualStop         bool
	AllowIsolate             bool
	DefaultDependencies      bool
	SurviveFinalKillSignal   bool
	CollectMode              string
	FailureAction            string
	SuccessAction            string
	FailureActionExitStatus  int
	SuccessActionExitStatus  int
	JobTimeoutSec            string
	JobRunningTimeoutSec     string
	JobTimeoutAction         string
	JobTimeoutRebootArgument string
	StartLimitIntervalSec    string
	StartLimitBurst          int
	StartLimitAction         string
	RebootArgument           string
	SourcePath               string
}

// Service holds the [Service] section options for a systemd service
type Service struct {
	Type                        string
	BusName                     string
	ExecStart                   string
	ExecStartPre                string
	ExecStartPost               string
	ExecReload                  string
	ExecStop                    string
	ExecStopPost                string
	Restart                     string
	RestartSec                  string
	RemainAfterExit             bool
	TimeoutSec                  string
	Environment                 string
	EnvironmentFile             string
	WorkingDirectory            string
	User                        string
	Group                       string
	PIDFile                     string
	StandardOutput              string
	StandardError               string
	SyslogIdentifier            string
	SyslogFacility              string
	SyslogLevel                 string
	SyslogLevelPrefix           bool
	CapabilityBoundingSet       string
	AmbientCapabilities         string
	NoNewPrivileges             bool
	NotifyAccess                string
	PrivateTmp                  bool
	ProtectSystem               string
	ProtectHome                 string
	ReadOnlyPaths               string
	ReadWritePaths              string
	InaccessiblePaths           string
	RuntimeDirectory            string
	StateDirectory              string
	CacheDirectory              string
	LogsDirectory               string
	ConfigurationDirectory      string
	ExitType                    string
	GuessMainPID                bool
	RestartMode                 string
	SuccessExitStatus           string
	RestartPreventExitStatus    string
	RestartForceExitStatus      string
	RootDirectoryStartOnly      bool
	NonBlocking                 bool
	Sockets                     string
	FileDescriptorStoreMax      int
	FileDescriptorStorePreserve string
	USBFunctionDescriptors      string
	USBFunctionStrings          string
	OOMPolicy                   string
	OpenFile                    string
	ReloadSignal                string
}

// Install holds the [Install] section options for a systemd service
type Install struct {
	WantedBy        string
	RequiredBy      string
	UpheldBy        string
	Alias           string
	Also            string
	DefaultInstance string
}
