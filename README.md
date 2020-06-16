# biu

#### 介绍
Biu是一个模型结构生成工具，利用模版生成代码，目前主要功能是利用Excel模板生成对应的实体文件支持CSharp与Golang生成。未来可以集成mysql中抽取生成，或者生成其他语言模板，还有生成markdown文档等等。

#### 软件架构
代码结构：

1. cli命令框架
2. Channel组成Pipline流水线

#### 使用教程

###### 创建模板文件

```shell
./biu -o c
```

###### 生成Go文件

1. -o operation 操作 gen生成 c创建模板 

2. -t源类型 excel 

3. -d 根路径 

4. -l 语言 go cs

```shell
./biu -o gen -t excel -d /bin/data -l go
```

###### 默认生成csharp文件，并且在同级目录下

```shell
./biu
```

#### 代码结构
1. 流水线 pipline 由各道工序组装而成，每道工序可以建立N个并行工序站，Fan-in来分流，Fan-out来合并管道
2. 取数器 pumper 产生数据，即生产者
3. 转化器 transformer 转化数据（拆解，分组，转化装配）
4. 生成器 gener 利用摸具（模板）生产最终结构数据
5. 输出器 exporter 输出成文件