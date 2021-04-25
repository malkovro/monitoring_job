# Monitoring Job

## Purpose

Build an executable to read the Temp & Humidity from a DHT22 and push it over to an API.

## Build it

```bash
$ docker build . -t mj_compiler
$ docker run -v $(pwd):/app mj_compiler
```

The executable will be: ./station

## To use it

Copy it on the pi:

```bash
$ scp ./station pi@the_pi_ip:/your/path/within/the/pi
```

Create a symbolic link of the station executable or add its location to the PATH:
```bash
$ sudo ln -s /home/leo/station /usr/local/bin/station
```

### Temperature and Humidity Reading:

On a pi, plug the DHT22 onthe GPIO 4 and launch the executable:

```
$ sudo station send climate -u https://your_monitoring_api.com
```

### Light Switch

Plug the relay board (HW 316) to the pi (use the 5V VCC of the pi even though you should be using an external power source with more mA...) and the relay input onto the GPIO 17.

```
$ sudo station switch light --on
$ sudo station switch light --on=false
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
    cmd: sudo station send climate -u https://monitoring-api.04plastic.com
    time: 0 */30 * * *
    onError: Continue
```

Add another set of jobs to switch the light on and off:
```.jobber
  TurnLightOn:
    cmd: sudo station switch light --on
    time: 0 0 6 * * *
    onError: Continue
  TurnLighOff:
    cmd: sudo station switch light --on=false
    time: 0 0 19 * * *
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
