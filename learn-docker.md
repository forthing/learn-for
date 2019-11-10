### docker笔记
> 本笔记记录docker学习历程，基础知识熟悉过，在第一篇做简要回顾。本学习笔记主要侧重于docker容器技术和容器云相关技术知识。
> 
#### 1、docker知识回顾
- 打包应⽤用及依赖包到⼀一个可移植的容器，发布到任何流⾏的Linux机器上，便可以实现虚拟化。
- Docker 镜像中包含了运行环境和配置，所以 Docker 可以简化部署多种应用实例工作。
- Docker 客户端通过命令行或者其他工具使用 [Docker API ](https://docs.docker.com/reference/api/docker_remote_api) 与 Docker 的守护进程通信。
- 通过docker的两个参数 `-i` `-t`，让docker运⾏的容器器现”对话”的能⼒。`-i`表 示允许你对容器的STDIN进⾏交互，`-t`在新容器器内指定⼀个伪终端或终端。 可以通过运⾏`exit`命令或者使⽤用`CTRL+D`来退出容器器。
- docker可以通过命令来进行相应操作，**docker run -d -p webapp python app.py**, `-d`让容器器在后台运⾏，-P将容器器内部使⽤的⽹络端口映射到我们使用的主机上。 使用 docker ps 来查看我们正在运⾏的容器。
- 创建镜像可以通过已有镜像改进或者用dockerfile制作镜像,镜像制作首先从FROM开始，除了选择现有镜像为基础镜像外，Docker 还存在一个特殊的镜像，名为 `scratch`。这个镜像是虚拟的概念，并不实际存在，它表示⼀个空白的镜像。
#### 2、docker镜像和容器理解
* 应⽤容器应该就是装载一些Application(应用)的⼀个容器，Dockers就是⼀个轻量级的虚拟机⽽已。可将 Docker 镜像看做只读模板，通过它可以创建 Docker 容器。Docker 容器就是 Docker 镜像的运行实例。
* 对于应⽤软件，镜像是软件生命 周期的构建和打包阶段，⽽容器则是启动和运行阶段。容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于⾃⼰的独立的命名空间。说到镜像与容器之间的关系，可以类⽐⾯向对象程序设计中的类和实例。
* 仓库理解：镜像构建完成后，可以很容易的在当前宿主机上运行，但是，如果需要在其它服务器上使⽤这个镜像，我们就需要⼀个集中的存储、分发镜像的服务，Docker Registry 就是这样的服务。
* image文件：Docker 把应⽤用程序及其依赖，打包在 image ⽂件⾥面。只有通过这个image⽂件，才能生成 Docker 容器。同一个image文件可以生成多个同时运行的容器实例。image 是⼆进制文件。实际开发中，一个 image ⽂件往往通过继承另⼀个image文件，加上一些个性化设置⽽⽣成。
* image文件生成的容器实例本身也是一个文件，称为容器文件，也就是说一旦容器生成，就会同时存在两个文件：image文件和容器文件，而且关闭容器并不会删除容器文件，只是容器停止运行而已。
#### 3、容器技术生态系统
|     |技术|说明|
|:---:|:---:|:---:|
|**容器核心技术**|容器规范、容器runtime、管理工具、定义工具、registries、容器OS|runtime是容器真正运行的地方、容器管理工具docker engine、dockerfile、容器OS|
|**容器平台技术**|容器编排引擎、容器管理平台、基于容器的PAAS|容器编排包括容器管理、调度、集群定义和服务发现，Rancher和ContainerShip是容器管理平台的代表，Deis\Flynn\Dokku是开源PAAS的代表|
|**容器支持技术**|容器网络、服务发现、监控、数据管理、日志管理、安全性|flannel\weave\calico都是开源解决方案，statusAPI、sysdig、cAdvisor开源容器监控方案，logspout和docker logs是开源日志管理方案|
#### 4、docker 运行方式
> - Docker daemon运行在host上，负责创建、运行、监控容器，构建、存储镜像，默认相应本地请求，若要远程，可以在配置文件中打开TCP监听。
> - 镜像具有分层结构，构建也是按层进行。容器启动后，一个新的可写层被加载到镜像的顶部。这一层叫做“容器层”。无论添加、删除、修改都发生在容器层，容器层是可写的，其他层都是只读的。容器层，用户看到的是一个叠加之后的文件系统。容器层修改文件时，具有copy-on-write特点。容器层保存的是镜像变化的部分。
> - dockerfile 构建镜像过程是从上层镜像运行一个容器，执行一条指令进行修改，执行类似docker commit的操作，生成一个新的镜像层。docker再基于刚刚提交的镜像运行一个新容器。这样循环构建下去。可以利用这个特点对容器进行调试，直到成功构建出该层镜像。
#### 5、dockerfile 常见指令和注意事项
- FROM、MAINTAINER、COPY、ADD、ENV、EXPOSE、VOLUME、WORKDIR、RUN、CMD、ENTRYPOINT
- 指令执行有shell模式和Exec模式，CMD和ENTRYPOINT推荐使用Exec格式，RUN两种格式都可以。
- ENTRYPOINT执行时不会被忽略，一定执行。
```shell
ENTRYPOINT ["/bin/echo", "hello"] CMD ["world"]
```
如果通过 `docker run -it [image] Cloudman`启动，则输出
```shell
hello Cloudman
```
> registry的由repository和tag两部分组成，完整格式为[registry-host]:[port]/[username]/xxx
> 

#### 6.容器相关技术
> 进入容器可以采用attach或者exec，attach进入的可以通过ctrl+p,然后ctrl+q退出。docker exec可进入相同的容器，执行exec的方式主要是 
> ```
> docker exec -it <container> bash|sh
> ```
- 容器一般分为服务类容器和工具类容器，服务类以daemon方式运行，
- docker rm是删除容器，docker rmi 是删除镜像
- docker -m或者 --memory-swap设置内存+swap的限额
- docker 可以通过-c限制核使用数，通过设置--blkio-weight来限制容器IO的优先级。
> - cgroup实现资源限额。前面做的相应的限额都是通过cgroup对进程配置的，我们可以在/sys/fs/cgroup/(cpu、memory、blkio)/docker/containerId查询相应配置信息。
> - namespace实现资源隔离。namespace实现了容器间资源的隔离，Linux使用了6种namespace，Mount、UTS、IPC、PID、Network、User。
> 
#### 7.容器网络
》主要讨论的是单个主机内的容器网络
> 容器网络有三种模式，none模式没有任何网络加载，对安全性要求高并且不需要联网的应用。
> host 模式共享docker host的网络栈，容器对网络性能要求较高的应用可以，不便之处是牺牲一些灵活性，要考虑端口冲突问题。
> bridge网络：docker安装会创建一个名为docker0的Linux bridge，通过veth pair来成对化容器中网络与网桥。
> user-defined网络：docker提供bridge\overlay\macvlan，后两种用于创建跨主机的网络。我们可以创建相应网络，在容器启动时指定使用，并且可以配置--subnet和--gateway，启动时通过--network指定，同时通过--subnet创建的网络能够指定静态IP。
- 容器间通信：可通过IP、DockerDNS Server 或joined容器三种方式通信。IP通信是指定网络，DNS方式只能在user-defined网络中使用。joined容器可以使2个容器或多个容器共享一个网络栈，非常适合不同容器程序通过loopback高效通信，监控其他容器的网络流量。
- 容器与外部世界连接：
--容器访问外部世界：将包的地址替换为host的地址发送出去，做了一次网络地址转换NAT,具体过程就是容器发送到docker0,docker0交给NAT处理，NAT将源地址换成host的IP，再访问外网。
--外部世界访问容器:主要依靠端口映射。通过`<host ip>:<32773>`这样的方式访问。
#### 8、容器存储
- storage driver:容器由最上面的可写容器层以及若干只只读的镜像层组成，容器数据就存放在这些层中。分层结构使镜像和容器的创建、共享以及分发变得非常高效，docker支持多种storage driver,由AUFS\Device Mapper\Btrfs\OverlayFS\VFS\ZFS,优先使用Linux发行版默认的storage driver。主要涉及无状态的应用，不需要持久化的数据。
- Data volume:volume数据可以永久保存，是目录或文件，具体使用上，docker使用两种类型。
-- bind mount: `-v <host path>:<container path>`即使容器销毁，数据依然存在，还可以指定读写权限，添加存在的单个文件。缺点是限制了容器的可移植性，如果要迁移到其他host，操作会失败。
-- docker managed volume:指定mount point就行，不需要指定源。可通过docker inspect查看volume。
- 数据共享：docker cp可以直接拷贝数据。
- volume container:专门为其他容器提供volume的容器。数据在host。
- data-packed volume container:数据打包到镜像中，容器可以自包含，不依赖host提供数据，适合只使用静态数据的场景。
#### 9.跨主机的容器管理与配置
- 利用docker machine远程安装多机器的docker。实现安装、创建、管理多台machine。
- 跨主机网络包括docker原生的overlay\macvlan,第三方方案有flannel\weave\calico.linux network namespace是Sandbox的标准实现，一个Endpoint只能属于一个网络，也只能属于一个Sandbox.
- 跨主机网络方案overlay：先部署Consul,创建overlay网络，在overlay中运行容器，验证overlay网络连通性。 docker会为每个overlay网络创建一个独立的nerwork namespace,其中会有一个linux bridge br0,endpoint还是由veth pair实现，一端连接到容器中，另一端连接到namespace的br0上。
- macvlan:利用linux kernel模块，允许同一个物理网卡配置多个mac地址，本质上是一种网卡虚拟化技术。macvlan会独占主机的网卡，一个网卡只能创建一个macvlan网络。macvlan网络的连通和隔离完全依赖VLAN\Ip subnet和路由，容器本身不做限制，用户可以像管理传统VLAN那样管理macvlan。
- flannel：为每个host分配一个subnet，容器从subnet中分配IP，可以在host间路由，实现跨主机通信。同一主机使用docker0连接，跨主机通过flannel转发。host-gw把每个主机都配置成网关，主机知道其他主机的subnet和转发地址，性能强于vxlan.
- weave：容器被接入类似于巨大的以太网交换机，所有容器都接入，无须NAT和端口映射。weave网络包含两个虚拟交换机，Linux bridge weave和Open vSwitch datapath, weave负责将容器接入网络，datapath负责在主机间VxLAN隧道中收发数据。
- calico：纯3层虚拟网路方案，为每个容器分配一个IP，每个host都是router，把不同host的容器连接起来。默认为容器只能与同一个calico网络中的容器通信，通过定义灵活的policy规则，可以进行控制。
#### 10、容器监控
- docker自带几个监控子命令，ps\top\stats,开源监控工具sysdig\Weave Scope\cAdvisor\Prometheus.
- prometheus:监控方案，提供监控数据收集、存储、处理、可视化、告警等功能。
#### 11.日志管理
> docker本身有Docker logs命令，该命令能够打印出自容器启动以来的完整日志，并且-f参数能继续打印新产生日志。docker默认日志启动是logging driver，，容器日志路径为/var/lib/docker/containers/`<container ID>`/ `<container ID>`-json.log
> - 开源日志方案ELK，Logstash负责从过年各个容器中提取日志，Logstash将日志转发到Elasticsearch进行索引和保存，Kibana分析和可视化数据。Filebeat可以将数据导入ELK。
#### 12.数据管理
volume driver；实现跨主机管理的数据维护，可以考虑基于Rex-Ray的driver，其生命周期不依赖Docker host和容器，是有状态容器理想的数据存储方式。