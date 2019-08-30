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
+ One RT will only contain 10 notices.
```go
// 一次发送的最大量
var SendMaxCnt int = 10
for i := 0; i < int(math.Ceil(float64(len(diffs))/float64(SendMaxCnt))); i++ {
    // ...
}
```
> text：消息标题，最长为256，必填。
> 
> desp：消息内容，最长64Kb，可空，支持MarkDown。

### Screenshot
![Screenshot](./assets/Screenshot.png)
![Screenshot 2](./assets/Screenshot_2.png)