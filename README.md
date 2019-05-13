# 说明

自动创建日志，日志的如下

127.0.0.1 - - Mon May 13 14:49:04 +0800 CST 2019 "OPTIONS /dig?refer=http%3A%2F%2Flocalhost%3A8888%2Fmovie%2F9233&time=Mon+May+13+14%3A49%3A04+%2B0800+CST+2019&ua=Mozilla%2F5.0+%28Symbian%2F3%3B+Series60%2F5.2+NokiaN8-00%2F012.002%3B+Profile%2FMIDP-2.1+Configuration%2FCLDC-1.1+%29+AppleWebKit%2F533.4+%28KHTML%2C+like+Gecko%29+NokiaBrowser%2F7.3.0+Mobile+Safari%2F533.4+3gpp-gba&uid=27e932fc8373b439f232755dc9135380&url=http%3A%2F%2Flocalhost%3A8888%2Fmovie%2F9233 HTTP/1.1" 200 43 "-" "ua" "-"


包括ip、日期、refer、url、ua、uid。目前版本ip是固定的，日期为当前时间，refer和url是根据ruleResouce()方法返回的resouce来创建url集合并随机选择。可以通过修改该方法来生成自己想要的url。ua则是在全局变量userAgents中随机选取。

# 使用

```
go run main.go -total 100 -filePath ./
```

total:生成的日志行数，默认是100

filePath:日志的存储路径，默认是当前路径(./)
