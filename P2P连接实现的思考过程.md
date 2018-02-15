1. 连接应该是一个自定义的Strut类似：
```Go
type PeerConn struct{
    IP string  //对应的连接 IP地址
    conn *net.TcpConn //对应的连接
    msgChannel  chan message //接收到的message 都维护到这个channel里面
}
```
2. 连接池
在进行GetHeadertongb