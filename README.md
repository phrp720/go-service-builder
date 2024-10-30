# service-builder

service-builder is a simple , customizable   Go library that  makes the creation of Linux and Windows services easier.

## Installation

```bash
go get github.com/phrp720/service-builder
```
## Features
- Builds and run Linux services with systemd
- Builds and run Windows services with nssm

## Usage

### For Linux

#### Example : Creating a systemd service
Running the code below will create a systemd service file named `dummy.service` in the `/etc/systemd/system/` directory with the content defined in the `service` object. After creating the file, it enables and starts the service automatically.

```go
builder := systemd.NewServiceBuilder()
service := builder.
// Unit
Description("Dummy Service").
Before("network.target").
After("network.target").
BindsTo("dummy.target").
Conflicts("shutdown.target").
Documentation("https://example.com").
OnFailure("reboot.target").
PartOf("multi-user.target").
Requires("network.target").
Wants("network-online.target").
// Service
ExecStart("/usr/bin/dummy").
ExecStartPre("/usr/bin/dummy-pre").
ExecStartPost("/usr/bin/dummy-post").
ExecStop("/usr/bin/dummy-stop").
// Install
RequiredBy("multi-user.target").
WantedBy("multi-user.target").
Build()

err := systemd.CreateService(service, "AbsolutePathWithService",false)
if err != nil {
fmt.Print(err)
return
}
err := systemd.StartService("ServiceName",false)
if err != nil {
fmt.Print(err)
return
}
```

### For Windows

#### Example: Creating a Windows service
Running the code below will create a Windows service via nssm  with the content defined in the `service` object. After creating the service, we run the StartService function which starts the service automatically.

```go
nssm.InitNssm("The_Place_Where_Nssm_Will_Be_Installed")
builder := nssm.NewServiceBuilder()
	service := builder.ServiceName("serviceName").AppDirectory("C:\\Program Files\\Service_Folder").Application("appName").Build()
	err := nssm.CreateService(service)
	if err != nil {log.Print(err)}
	err = nssm.StartService(service.ServiceName)
	if err != nil {log.Print(err)}
```
> [!Warning]
>
> To create services, you must run your go project as an Administrator.

