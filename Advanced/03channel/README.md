###通道的操作
_箭头代表的是数据流向_
- 发送值  
`ch1 <- 1`
- 接受值  
`x := <- ch1 给x赋值`  
`<- ch1 直接丢掉了` 
- 关闭  
`close(ch1)`
_关闭一个未初始化的或已经关闭的channel会引发panic_    
![channel](./assets/channel-1645622455690.png)
