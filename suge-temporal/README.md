# 搭建环境
首先安装 docker-compose
```bigquery
$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
## 搭建kafka环境
编写docker-compose.yml文件

```bigquery
[root@cloud ]# cat docker-compose.yml 
version: '3.3'
services:
  zookeeper:
    image: wurstmeister/zookeeper
    container_name: zookeeper
    ports:
      - 2181:2181
    volumes:
      - /data/zookeeper/data:/data
      - /data/zookeeper/datalog:/datalog
      - /data/zookeeper/logs:/logs
    restart: always
  kafka1:
    image: wurstmeister/kafka
    depends_on:
      - zookeeper
    container_name: kafka1
    ports:
      - 9092:9092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: 192.168.56.4:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.56.4:9092
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092
      KAFKA_LOG_DIRS: /data/kafka-data
      KAFKA_LOG_RETENTION_HOURS: 24
    volumes:
      - /data/kafka1/data:/data/kafka-data
    restart: unless-stopped  
  kafka2:
    image: wurstmeister/kafka
    depends_on:
      - zookeeper
    container_name: kafka2
    ports:
      - 9093:9093
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ZOOKEEPER_CONNECT: 192.168.56.4:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.56.4:9093
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9093
      KAFKA_LOG_DIRS: /data/kafka-data
      KAFKA_LOG_RETENTION_HOURS: 24
    volumes:
      - /data/kafka2/data:/data/kafka-data
    restart: unless-stopped
  kafka3:
    image: wurstmeister/kafka
    depends_on:
      - zookeeper
    container_name: kafka3
    ports:
      - 9094:9094
    environment:
      KAFKA_BROKER_ID: 3
      KAFKA_ZOOKEEPER_CONNECT: 192.168.56.4:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.56.4:9094
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9094
      KAFKA_LOG_DIRS: /data/kafka-data
      KAFKA_LOG_RETENTION_HOURS: 24
    volumes:
      - /data/kafka3/data:/data/kafka-data
    restart: unless-stopped
```

192.168.56.4 地址要替换成kafka主机地址

启动kafka环境
```bigquery
cd docker-compose
(base) daviddeMacBook-Pro:docker-compose david$ docker-compose up -d
```
创建名字为suge-task topic
```bigquery
docker-compose exec kafka1 bash
cd /opt/kafka/bin 
./kafka-topics.sh --create --topic suge-task --zookeeper 127.0.0.1:2181 --partitions 3 --replication-factor 2
```

## 搭建temporal环境
克隆代码
```bigquery
git clone git@github.com:daijinwei/DDD-solidity.git
```
启动temppral环境
```bigquery
cd docker-compose
(base) daviddeMacBook-Pro:docker-compose david$ docker-compose up -d

```
temppral环境是否运行正常
```bigquery
docker-compose david$ docker-compose ps
         Name                       Command               State                               Ports                             
--------------------------------------------------------------------------------------------------------------------------------
temporal                 /etc/temporal/entrypoint.s ...   Up      6933/tcp, 6934/tcp, 6935/tcp, 6939/tcp,                       
                                                                  0.0.0.0:7233->7233/tcp, 7234/tcp, 7235/tcp, 7239/tcp          
temporal-admin-tools     tini -- sleep infinity           Up                                                                    
temporal-elasticsearch   /bin/tini -- /usr/local/bi ...   Up      9200/tcp, 9300/tcp                                            
temporal-postgresql      docker-entrypoint.sh postgres    Up      5432/tcp                                                      
temporal-ui              ./start-ui-server.sh             Up      0.0.0.0:8080->8080/tcp 
```

UI访问路径
```bigquery
http://127.0.0.1:8080/namespaces/default/workflows
```
后端访问地址
```bigquery
127.0.0.1:7233
```

# 编译代码
```bigquery
make create-work
make start-produce-work
make start-consume-work
```
# 运行程序
## 提交work
```bigquery
./bin/create-work_linux_amd64
```

## 开始生产者
```bigquery
./bin/start-produce-work_linux_amd64
```
## 开始消费者
```bigquery
./bin/start-consume-work_linux_amd64
```
# 测试
本地测试
单元测试暂时补充