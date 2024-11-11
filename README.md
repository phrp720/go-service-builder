# go-service-builder

go-service-builder is a simple , customizable   Go library that  makes the creation of Linux and Windows services easier.

## Installation

```bash
go get github.com/phrp720/go-service-builder
```
## Features
- Builds and run Linux services with systemd
- Builds and run Windows services with nssm

## Usage

### For Linux

#### Example : Creating a systemd service 
Running the code below will generate and start a systemd service  named `dummy.service` in the `/etc/systemd/system/`.

```go
builder := systemd.NewServiceBuilder()
service := builder.
	// Unit
	Description("Dummy Service").
	Before("network.target").
	After("network.target").
	Documentation("https://example.com").
	OnFailure("reboot.target").
	PartOf("multi-user.target").
	Requires("network.target").
	// Service
	ExecStart("/usr/bin/dummy").
	ExecStop("/usr/bin/dummy-stop").
	// Install
	RequiredBy("multi-user.target").
	Build()

err := systemd.CreateService(service, "/etc/systemd/system/dummy.service",true)
if err != nil {
	fmt.Print(err)
return
}
err := systemd.StartService("dummy.service",true)
if err != nil {
	fmt.Print(err)
return
}
```

### For Windows

#### Example: Creating a Windows service
Running the code below will create and start a Windows service via [nssm](https://nssm.cc/).
Then nssm that we initiated will be responsible for the health of the service.

```go
err:=nssm.InitNssm("Folder_Where_Nssm_Will_Be") // InitNssm downloads nssm and extracts it  to the specified folder
if err != nil {log.Print(err)}
builder := nssm.NewServiceBuilder()
service := builder.ServiceName("dummyService").
	AppDirectory("C:\\Program Files\\Service_Folder").
	Application("C:\\Program Files\\Service_Folder\\dummyService.exe").
	Build()
err := nssm.CreateService(service)
if err != nil {log.Print(err)}
err = nssm.StartService(service.ServiceName)
if err != nil {log.Print(err)}
```
> [!Warning]
>
> To create services in Windows or in places where root privileges are needed, it is mandatory to  run your project as an Administrator.

