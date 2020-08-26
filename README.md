# go-installer
Generate Installer for your application using GoLang

## Purpose
Most of the GoLang based applications are self sufficient single executable. Sometime one may need to ship some static
files or images or platform dependent DLLs or Shared library in case of CGO based applications.

This project helps to create that installer only by configuring the single `config.json` file.
This project is internally using [binclude](https://github.com/lu4p/binclude)

## Usage
* Clone this repository
* Create your own `config.json` See example below
* Create your own `file.list` See example below
* Build using command `make build`
* Make sure you have `PATH` contains `$GOPATH/bin`

## Typical `config.json`
```json
{
	"Components": [{
			"Name": "MyProduct",
			"TargetDir": "bin",
			"Files": [{
					"Name": "file1.txt",
					"Platform": "linux",
					"Action": {
						"command": ""
					},
					"SourceRelativePath": "./stage/file1.txt"
				},
				{
					"Name": "file2.txt",
					"Platform": "linux",
					"Action": {
						"command": ""
					},
					"SourceRelativePath": "./stage/file2.txt"
				}
			]
		},
		{
			"Name": "MyProductLib",
			"TargetDir": "lib",
			"Files": [{
					"Name": "file3.txt",
					"Platform": "linux",
					"Action": {
						"command": ""
					},
					"SourceRelativePath": "./stage/file3.txt"
				},
				{
					"Name": "file4.txt",
					"Platform": "linux",
					"Action": {
						"command": ""
					},
					"SourceRelativePath": "./stage/file4.txt"
				},
				{
					"Name": "file5.txt",
					"Platform": "linux",
					"Action": {
						"command": ""
					},
					"SourceRelativePath": "./stage/file5.txt"
				}
			]
		},
		{
			"Name": "MyProductExtra",
			"TargetDir": "extra",
			"Files": [{
				"Name": "allfiles.tar",
				"Platform": "linux",
				"Action": {
					"command": ""
				},
				"SourceRelativePath": "./stage/allfiles.tar"
			}]
		}
	]
}
```
## Typical `file.list`
```bash
./stage/file1.txt
./stage/file2.txt
./stage/file3.txt
./stage/file4.txt
./stage/file5.txt
./stage/allfiles.tar
```