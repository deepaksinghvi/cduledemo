# cduledemo

A demo project for [cdule](https://github.com/deepaksinghvi/cdule) library.

## Usage Instructions or Steps

### 1. cdule library

Add cdule library as the dependency
```
go get github.com/deepaksinghvi/cdule
```
or use the following for go1.18 with tag v1.0.3

```
go1.18beta2 get github.com/deepaksinghvi/cdule@v1.0.3
```

### 2. Initiate Cdule Background Watcher
```
cdule := cdule.Cdule{}
cdule.NewCdule()

```

### 3. Build Job
```
testJob := job2.TestJob{}
jobData := make(map[string]string)
jobData["one"] = "1"
jobData["two"] = "2"
jobData["three"] = "3"
testJobModel, _ := watcher.NewJob(&testJob, jobData).Build(utils.EveryMinute)
```
On submission with the help of the Build() method, job will get scheduled and it will keep running it based on the cron scheduled provided.

### 4. Stop Cdule Watcher
```
time.Sleep(5 * time.Minute)
cdule.StopWatcher()
```

### 5. Run From command Line

Worker name to be picked based on the hostname
```
./cdule
```

or worker name can be supplied in the format of 
cdule <workername> 
 
```
cdule worker1
```

For more details check documentation.
* [cdule_repo](https://github.com/deepaksinghvi/cdule) 
* [cdule_wiki](https://github.com/deepaksinghvi/cdule/wiki)