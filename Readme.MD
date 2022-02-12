# cduledemo

A demo project for [cdule](https://github.com/deepaksinghvi/cdule) library.

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
