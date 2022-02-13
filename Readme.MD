# cduledemo

A demo project for [cdule](https://github.com/deepaksinghvi/cdule) library.

```
go get github.com/deepaksinghvi/cdule
or use the following for go1.18 with tag v1.0.2
go1.18beta2 get github.com/deepaksinghvi/cdule@v1.0.2

```
#### Initiate Cdule Background Watcher
```
cdule := cdule.Cdule{}
cdule.NewCdule()

```

#### Build Job
```
	testJob := job2.TestJob{}
	testJobModel, _ := watcher.NewJob(&testJob, nil).Build(utils.EveryMinute)
	printJobs(testJobModel)
```

#### Stop Cdule Watcher
```
	time.Sleep(5 * time.Minute)
	cdule.StopWatcher()
```


For more details check [cdule](https://github.com/deepaksinghvi/cdule) documentation.
