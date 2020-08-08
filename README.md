# scut-academic-notifier

+ ReSend the notices in SCUT Academic Notice (华工教务通知) and SE College Focus (软院公务通知)
+ Related websites: [SCUT Office of Academic Affairs](http://jw.scut.edu.cn/zhinan/cms/index.do), [SCUT Software Engineering College Focus](http://www2.scut.edu.cn/sse/xyjd_17232/list.htm)
+ Telebot API see [Telegram Bot API](https://core.telegram.org/bots/api)
+ Serverchan API see [Server酱](http://sc.ftqq.com/3.version)

### Endpoints

```text
*Common*
/start - show start message
/help - show this help message
/cancel - cancel the last action

*Account*
/bind - bind with sckey
/unbind - unbind this chat

*Notifier*
/send - send the current notices
```

### References

+ [tucnak/telebot.v2](https://github.com/tucnak/telebot/tree/v2)
+ [Aoi-hosizora/go-serverchan](https://github.com/Aoi-hosizora/go-serverchan)
+ [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
+ [Aoi-hosizora/telebot-wechat-scaffold](https://github.com/Aoi-hosizora/telebot-wechat-scaffold)
