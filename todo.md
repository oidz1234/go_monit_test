* Create initila project structure

* Try and run functions from another folder

* odd bug with compliing, have to do somethign like this to work around

```
CGO_ENABLED=0 go build -o agent main.go
```

Keep an eye on this thread to fix things:

<https://github.com/golang/go/issues/26492>


* Add in the following:
    * Hostname
    * Uptime
    * Temp
    * who is logged in (on hold)
    * services running
    * load
    * Network/interface traffic
    * Mem used
    * disk used
    * i/o
    * Heartbeat (last checked in datetime)
        * if > X alert (for frontend todo)
    
   


* Create struct with all the data we will pass to the api
* Pass that data to the api




* one easy step install command
