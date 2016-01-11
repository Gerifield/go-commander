# Go commander

This is a very small and easy ssh command runner.
You can configure it with simple yaml file to run multiple commands on a server.

Config example:

```
Server: 192.168.0.1
User: root
Key: /.ssh/id_rsa
Commands: 
- ps aux
```

