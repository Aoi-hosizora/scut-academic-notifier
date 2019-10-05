# SCUT_Academic_Notifier
+ RT the notices in SCUT Academic Notice (教务通知) and SE College Focus (软院公务通知)
+ [SCUT Office of Academic Affairs](http://jw.scut.edu.cn/zhinan/cms/index.do)
+ [SCUT Software Engineering College Focus](http://www2.scut.edu.cn/sse/xyjd_17232/list.htm)

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
+ Will only send notices up to 1 month ago

```go
// A month
var SendTimeRange time.Duration = 30 * 24 * time.Hour

sendLists := make([]models.NoticeItem, 0)
for _, v := range diffs {
    nt, err := time.ParseInLocation("2006-01-02 15:04:05", v.Date+" 00:00:00", time.Local)
    if err == nil && nt.After(time.Now().Add(-SendTimeRange)) {
        sendLists = append(sendLists, v)
    }
}
```

+ Each RT will only contain 10 notices at most

```go
var SendMaxCnt int = 10

// Ceiling
ceil := int(math.Ceil(float64(len(sendLists)) / float64(SendMaxCnt)))
for i := 0; i < ceil; i++ {
    msg := ""
    for j := i * SendMaxCnt; j < i*SendMaxCnt+SendMaxCnt; j++ {
        if j < len(sendLists) {
            ni := sendLists[j]
            msg += fmt.Sprintf("+ %s\n", ni.String())
        } else {
            break
        }
    }
    // Send
}
```

+ Will only retry 5 times when panic continuely

> text：消息标题，最长为256，必填。
> 
> desp：消息内容，最长64Kb，可空，支持MarkDown。
>
> Refer from http://sc.ftqq.com/3.version

### Screenshot
![Screenshot](./assets/Screenshot.png)
![Screenshot 2](./assets/Screenshot_2.png)