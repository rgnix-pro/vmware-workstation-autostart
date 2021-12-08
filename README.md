# VMware Workstation AutoStart
[![Build Status](https://api.travis-ci.com/rgnix-pro/vmware-workstation-autostart.svg?branch=main&status=passed)](https://travis-ci.com/github/rgnix-pro/vmware-workstation-autostart)
### This is an auto start app for VMware Workstation to auto start VM's on windows reboot with VMware Workstation installed.
## Setup

#### 1. Edit config.json to match your paths

```javascript
{
    "vm-start": "D:\\Virtual Machines\\AutoStart\\start-vmware.json",
    "vmrun": "C:\\Program Files (x86)\\VMware\\VMware Workstation\\vmrun.exe"
}
```

#### 2. Edit or Create start-vmware.json 
```javascript
{
    "_comment": "Make sure the name is the same folder name as the .vmx file",
    "vs": [

        {
            "location": "D:\\Virtual Machines\\",
            "name": "vcenter.example.local",
            "start": true,
            "gui": false
        },
        {
            "location": "D:\\Virtual Machines\\",
            "name": "vsan-witness.example.local",
            "start": true,
            "gui": false
        }
    ]
}
```
#### 3. Add vm-start.exe & config.json to your startup folder on windows or create a short cut for vm-start.exe to its current folder 

`windows+r shell:startup `
