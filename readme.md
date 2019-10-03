# SCUT_Academic_Notifier
+ RT the notices in SCUT Academic Notice
+ [SCUT office of Academic Affairs](http://jw.scut.edu.cn/zhinan/cms/index.do)

### Environment
+ go 1.11.5 windows/amd64
+ Notifier: [ServerChan](http://sc.ftqq.com/3.version)
+ Requirement: None

### Setup
```json
// config.json
{
    "SCKEY": "xxxxxx"
}
```

### Run
```bash
# Setup config.json

go run app.go
```

### Tips
+ Will only send notices up to 15 days ago

```go
// 半个月
var SendTimeRange time.Duration = 15 * 24 * time.Hour
nt, err := time.ParseInLocation("2006-01-02 15:04:05", v.Date+" 00:00:00", time.Local)

if err == nil {
    if nt.Before(time.Now().Add(-SendTimeRange)) {
        continue
    }
}
```

> text：消息标题，最长为256，必填。
> 
> desp：消息内容，最长64Kb，可空，支持MarkDown。
>
> Refer from http://sc.ftqq.com/3.version

### Screenshot
![Screenshot](./assets/Screenshot.png)
![Screenshot 2](./assets/Screenshot_2.png)