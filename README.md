# go-service-builder

`go-service-builder` is a lightweight and customizable Go library designed to simplify the creation and management of Linux and Windows services.

## Installation

```bash
go get github.com/phrp720/go-service-builder
```
## Features
- **Linux Support:** Create and manage services with `systemd`.
- **Windows Support:** Create and manage services using `nssm`.

## Usage

### For Linux: Creating a systemd Service 

The following example demonstrates how to create and start a systemd service named `dummy.service` in `/etc/systemd/system/`.

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

### For Windows: Creating a Service with NSSM

The example below illustrates how to create and start a Windows service using [nssm](https://nssm.cc/). NSSM will ensure the service's health.

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
> To create services on Windows or to write services in restricted directories on Linux, root or Administrator privileges are required.

## Contributing
Contributions, issues, and feature requests are welcome! Feel free to check out the repository and submit a pull request or report any issues.
