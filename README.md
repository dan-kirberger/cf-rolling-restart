# cf-rolling-restart
Rolling restart plugin for Cloud Foundry CLI. Serving as my introduction to Go as well so slow and sloppy progress
to come initially!

## Purpose
Perform a rolling restart of an application, n instances at a time (currently 1). Going this route may be easier for some
organizations than doing a full new deploy, which may involve a comparatively long build or stage process.

## Usage
```
 $> cf rolling-restart my-app
    Performing rolling restart of my-app
    Instance 0 restarting. Waiting for minimum healthy instances before proceeding (Currently 2/2)
    Instance 1 restarting. Waiting for minimum healthy instances before proceeding (Currently 2/2)
    Successfully restarted my-app
```

## Caveats (and future work)
* Only supports one app at a time (for now)
* Only supports one instance at a time (for now)
* Your app will be running at reduced capacity until the restart is complete (for now)
* Needs tests
* Needs dependency management/build process