* Create initila project structure

* Try and run functions from another folder

* odd bug with compliing, have to do somethign like this to work around

```
CGO_ENABLED=0 go build -o agent main.go
```

Keep an eye on this thread to fix things:

<https://github.com/golang/go/issues/26492>


* Add in the following:
    * Hostname ✅
    * Uptime ✅
    * Temp ✅
    * who is logged in (on hold)
    * services running ✅
    * load ✅
    * Network/interface traffic ✅
    * Mem used  ✅
    * disk used ✅ 
        * Sensible defualts
        * Could get the X biggest disks and use those as defalts for now
    * i/o ✅
    * Heartbeat (last checked in datetime) ✅
        * if > X alert (for frontend todo)
    * disk read
    * disk write
    
* Default disks
    * Get default disks (top 2(?) by useage)
        * DOn't do for now, take to long (use / and home for linux and C for
          windows I guess)
    * Can configure these on frontend eventually
   


* Create struct with all the data we will pass to the api ✅
* Convert into json ✅
* Pass that data to the api






* one easy step install command

* Name ideas
    * Inspector Puffin
    * Monke Monitor/ Monitor Monke(y)

* morpheus api 
    * For agent error repoerting/monitoring


* Ability to configure the following on web ui
    * What disks to monitor 
    * What services to monitor

* all this configration stored sepratly and checked on first run and after
  every run

# b4 launch

* get the net stuff working lol (much easier to do on agent) ✅
* Wrap the core loop in a try catch ting incase backend is broken 
    * ORRRRR catch error and handle them instead of panic
* Logging proper (a bit, just connection stuff)
