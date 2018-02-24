
## 节点的维护
1. P2P 连接一旦建立以后，双方可以通过`getaddr` 和 `addr`来交换本地有效的链接
2. 对于维护的链接当节点的连接状态遇到下面的情况之一时，连接认为节点已经是无效的了
    + It claims to be from the future
    + It hasn't been seen in over a month
    + It has failed at least three times and never succeeded
    + It has failed ten times in the last week

     所以结构体应该是这样的：
      ```
        type Node struct {
            Addr        string
            Src         string
            Attempts    int
            TimeStamp   int64
            LastAttempt int64
            LastSuccess int64
            // no refcount or tried, that is available from context.
        }
      ```


 3. 对于每个节点来讲，在选择节点进行通信时应该选择最近活跃的那个节点(最近连接过的那个节点)
 4. 在单独的一个文件中存储这些节点(序列化为JSON),就像btcd那样。
## Sync Block实现
> 根据[官方教程](https://bitcoin.org/en/developer-guide#initial-block-download)里面的说反，SyncBlock一般是在第一次接入
Bitcoin网络或者是从Bitcoin网络中东Offline很久后才会调用(**一般本地最高Block落后24小时**)，同步方式有两种，一种是[blocks-first](https://bitcoin.org/en/developer-guide#blocks-first)
方式，另一种是[headers-first](https://bitcoin.org/en/developer-guide#headers-first)方式，这里我只考虑Header-Frist方式的实现

###  Header-Frist实现

![image.png](http://upload-images.jianshu.io/upload_images/5247090-d5273481168ded4a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

1. 获取区块链的Headeer
2. 发送GetHeaders Message获取所欲的未同步的区块头,最大为2000个Header队列，获取到以后
3. 收到Headers队列的响应后，开始初步验证
4. 校验规则： 确保所有字段遵循共识规则，并且根据nBits字段，Header的Hash值的长度低于目标阈值。
5. 在进行局部校验的同时，还可以同步的对已经验证好的Header请求对应的Block数据
---
1. 获取区块头放到一个单独的线程中去执行
2. 第一步获取到的区块头信息，通过channel发送到主线程上进行合法性验证
3. 这个时候我们需要一个连接池通过区块头来同步区块
4. 为了完成第三步，我们需要一个连接池来维护所有的连接，并且，需要见有效的节点的IP维护起来

