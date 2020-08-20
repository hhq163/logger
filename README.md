# logger
基于uber zap封装的日志类

1.支持生成json格式日志和普通调试日志    
+ json格式：  
{"level":"debug","ts":"2019-12-03T10:39:07.552+0800","caller":"test/test_queue.go:69","msg":"get item queueID:test begin","svcName":""}

+ 普通调试日志：  
2019-08-21T17:35:01.736+0800	debug	impl/impl.go:19	gateway NewImpl begin

2.支持udp协议上传到elk  