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
Running the code below will create a systemd service file named `dummy.service` in the `/etc/systemd/system/` directory with the content defined in the `service` object. After creating the file.We start the service with StartService function(the example below is a root service so you need to run your code project as superUser).

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

err := systemd.CreateService(service, "/etc/systemd/system/dummy.service",true)
if err != nil {
	fmt.Print(err)
return
}
err := systemd.StartService("ServiceName",true)
if err != nil {
	fmt.Print(err)
return
}
```

### For Windows

#### Example: Creating a Windows service
Running the code below will create a Windows service via nssm  with the content defined in the `service` object. After creating the service, we run the StartService function which starts the service automatically.
Then nssm that we initiated will be responsible for the service.

```go
err:=nssm.InitNssm("Folder_Where_Nssm_Will_Be_Installed") // InitNssm downloads nssm and extracts it  to the specified folder
if err != nil {log.Print(err)}
builder := nssm.NewServiceBuilder()
service := builder.ServiceName("dummyService").
	AppDirectory("C:\\Program Files\\Service_Folder").
	Application("C:\\Program Files\\Service_Folder/dummyService.exe").
	Build()
err := nssm.CreateService(service)
if err != nil {log.Print(err)}
err = nssm.StartService(service.ServiceName)
if err != nil {log.Print(err)}
```
> [!Warning]
>
> To create services in Windows, it is mandatory to  run the go project as an Administrator.

