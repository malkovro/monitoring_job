# Monitoring Job

## Purpose

Build an executable to read the Temp & Humidity from a DHT22 and push it over to an API.

## Build it
$ docker build . -t mj_compiler
$ docker run -v $(pwd):/app mj_compiler

The executable will be: ./app

## To use it

Copy it on the pi:

```bash
$ scp app pi@the_pi_ip:
```

On a pi, plug the DHT22 onthe GPIO 4 and launch the executable:

```
$ sudo ./app -monitoringBaseUrl https://your_monitoring_api.com
```

## To schedule it

Staying  in the go environmnent, we can use [jobber](https://github.com/dshearer/jobber).

It will require goland to be installed:
```bash
$ apt-get update && apt-get install golang # Install golang
$ wget https://github.com/dshearer/jobber/archive/v1.4.4.tar.gz # Download one of the release https://github.com/dshearer/jobber/releases
$ tar xcf v1.4.4.tar.gz # Extract the tar
$ cd jobber-1.4.4
$ make
$ sudo make install
$ sudo mkdir /usr/local/var/jobber/ # Looks like jobber will have trouble to start if this directory is not created
```

We need to daemonize jobbermaster so it can run then all our individual jobs, we can use systemd to do so:
```bash
$ cd /etc/systemd/system
$ sudo  vi  jobber.service
```

Edit the file to something like this:
```systemd.service
[Unit]
Description=Jobber daemon service
After=network-online.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=5
ExecStart=/usr/bin/env /usr/local/libexec/jobbermaster

[Install]
WantedBy=multi-user.target
```

Eventually enable jobber to work on  startup:
```bash
$ systemctl enable jobber
```

Now  we can  create our job, init a job file:
```bash
$ jobber init
```

Edit the jobber conf adding a job of the like:
```.jobber
  MonitoringReading:
    cmd: sudo /home/leo/app -monitoringBaseUrl=https://monitoring-api.04plastic.com
    time: 0 */30
    onError: Continue
```

In order for this to work, your user (myuser) must have password disabled for sudo, you can do this by adding a file to the  sudoers:
```bash
$ sudo visudo /etc/sudoers.d/010_myuser-nopasswd
```

The content must be something like that:
```sudoer
myuser ALL=(ALL) NOPASSWD: ALL
```

**🎉 Now all we can reload jobber with `jobber reload` and the job is scheduled, every 30 minutes a reading will be taken and uploaded to the monitoring API.**