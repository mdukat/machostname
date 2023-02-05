# machostname

## Set hostname based on MAC address

### Why?

I needed some mechanism to automatically set hostname of my Fujitsu Futro S900 cluster booted from single PXE root hosted over NFS in read-only. Also I'm learning Go. Code itself is in single function, as it has to do one job, and do it well.

### How?
```
make
make install
systemctl enable machostname
```

### Example systemctl log
```
root@eric:~# systemctl status machostname
● machostname.service - Set hostname based on MAC address
     Loaded: loaded (/etc/systemd/system/machostname.service; enabled; vendor preset: enabled)
     Active: active (exited) since Sun 2023-02-05 19:10:58 UTC; 15min ago
    Process: 286 ExecStart=/usr/bin/machostname (code=exited, status=0/SUCCESS)
   Main PID: 286 (code=exited, status=0/SUCCESS)
        CPU: 21ms

Feb 05 19:10:58 optiplex systemd[1]: Starting Set hostname based on MAC address...
Feb 05 19:10:58 optiplex machostname[286]: Configuration not found in current directory open machostname.conf: no such file or directory
Feb 05 19:10:58 optiplex machostname[286]: Using /etc/machostname.conf
Feb 05 19:10:58 optiplex machostname[286]: Found interface 90:1b:0e:29:18:97
Feb 05 19:10:58 optiplex machostname[286]: Found the hostname eric 90:1b:0e:29:18:97
Feb 05 19:10:58 optiplex machostname[286]: Hostname changed to eric
Feb 05 19:10:58 eric systemd[1]: Finished Set hostname based on MAC address.
```
