# gop2p

golang p2p

### UDP 协议

client -> server : 上线

a

client -> server : 请求访问其他节点

c,{mac}

server -> client : 转发 其他节点想访问

cc,{ip},{port}

client -> client : 发送 NAT 表试探

ccc

client -> client : 发送 正式连接

cccc

### 交互流程

1. client 监听 UDP 11191 端口 ，并发送 a
2. client1 请求访问其他节点，发送 c ，3 秒后发送 cccc
3. server 转发 client1 的 cc 请求 到 client2
4. client2 接到 cc 后，发送 ccc
