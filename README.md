# Go commander

This is a very small and easy ssh command runner.
You can configure it with simple yaml files to run multiple commands on multiple servers.

Config example:

```
Servers: [192.168.0.1]
User: root
Key: /.ssh/id_rsa
Commands: 
- ps aux
```

Just put the config files into the `configs` folder. The commands would be executed in alphabetical order.
If you want, you can add subfolders too.

## Depencie install

```
go get gopkg.in/hypersleep/easyssh.v0
go get gopkg.in/yaml.v2
```

Or simple with glide:

```
glide install
```