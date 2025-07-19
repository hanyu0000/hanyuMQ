消息队列系统中的五个关键组件及其相互关系：**Broker**、**Topic**、**Partition**、**客户端（Client/消费者）**、**消费者组（Consumer Group）**

### Broker

* 定义：Broker 是消息中间件的服务器节点，多个 Broker 可组成集群。
* 职责：接收生产者发送的消息、存储消息、响应消费者拉取请求。
* 存储结构：使用 `map[string]*Topic` 对不同 Topic 进行管理。

---

### Topic

* 定义：Topic 是消息的逻辑分类，相当于队列名称。
* 水平扩展：为了提高系统吞吐量和可用性，一个 Topic 可被拆分成多个 Partition，每个 Partition 可部署在不同的 Broker 上。
* 顺序性：Topic 級别不保证全局顺序，但单个 Partition 内消息保持严格顺序。

---

### Partition

* 定义：Partition 是 Topic 的物理分片，内部维护一个消息队列。
* 存储与持久化：

  * 内存队列：实时将消息追加到内存。
  * 持久化：将消息按块写入磁盘文件，以 `topic_partition` 命名；消费时由消费者维护 `offset`，通过块索引转换读取数据。
* 消费保证：支持至少一次消费或仅一次消费语义，可通过消费者实例并行消费同一 Partition 的不同消息。

---

### Subscription 模式

1. **点对点（P2P）**

- 每个 Partition 上同一时刻只有一个消费者实例在消费消息，保证消息仅被消费一次。
- 如无对应消费者组则动态创建组并加入消费者。
2. **发布/订阅（Pub/Sub）**

- 每个 Topic 的 Partition 可被多个消费者组同时订阅，实现消息广播。
- 不同消费者组之间互不影响，同组内消息在成员间负载分配。

---

### 消费者组（Consumer Group）

* 定义：一组消费者实例对一个或多个 Topic 进行联合消费。
* 标识：`group.id`（字符串）唯一标识一个消费者组。
* 分区分配：

  * 同组：一个组内的 Partition 在组内各消费者间平均分配，理想情况下消费者数与分区数保持一定比例。

    * 例：100 个消费者、20 个 Topic（每 Topic 5 个 Partition），每个 Topic 分配 5 个消费者。
  * 不同组：不同组订阅同一 Topic，实现广播消费；组间相互独立，组内可手动或自动分配各 Partition。

---

### 组件关系总览

```
Producer ──▶ Broker Cluster ──▶ Topic ──+──▶ Partition 1 ──▶ Consumer Group A ──▶ Consumer A1
                                         |                                  └─▶ Consumer A2
                                         +──▶ Partition 2 ──▶ Consumer Group B ──▶ Consumer B1
                                         ...
```

* Producer：向 Topic 发布消息。
* Broker Cluster：维护多个 Topic 与 Partition，提供存储与路由。
* Partition：消息物理存储单元，单 Partition 内顺序有保障。
* Consumer Group：协调组内消费者实例消费分区，支持点对点和发布/订阅两种模式。
