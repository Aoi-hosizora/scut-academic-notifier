# SCUT_Academic_Notifier

+ RT the notices in SCUT Academic Notice (教务通知) and SE College Focus (软院公务通知)
+ [SCUT Office of Academic Affairs](http://jw.scut.edu.cn/zhinan/cms/index.do)
+ [SCUT Software Engineering College Focus](http://www2.scut.edu.cn/sse/xyjd_17232/list.htm)

### Requirement

+ Notifier: [ServerChan](http://sc.ftqq.com/3.version)

### Run

```bash
# Copy config.example.yaml to ./src/config/static.go
# Setup wechat key

go run main.go
```

+ To change long polling frequent, please modify in `./src/config/config.yaml`

```yaml
server:
  polling-duration: 120 # second
```

+ For api and send content config, please modify in `./src/config/static.go`

```go
static := &Static{
    SendMaxCount: 10,                  // send 10 notices at most
    SendRange:    60 * 24 * time.Hour, // only send notice last 2 months
}
```

### Screenshots

+ Wechat

![Screenshots](./assets/Screenshot.png)
![Screenshots 2](./assets/Screenshot_2.png)

+ Server

![Screenshots_3](./assets/Screenshot_3.png)