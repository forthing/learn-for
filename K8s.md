## K8S概念与相应技术
#### 1、K8S核心概念及相应功能介绍

> kubernetes是开源容器集群管理系统，主要功能有：
> - 容器的应用部署、维护和滚动升级
> - 负载均衡和服务发现
> - 跨机器和跨地区的集群调度
> - 自动伸缩
> - 无状态服务和有状态服务
> - 广泛的volume支持
> - 插件机制保证扩展性

*核心概念*有：
+ **Pod**：一组紧密关联的容器集合，支持多个容器在一个Pod中共享网络和文件系统，可以通过进程间通信和文件共享这种方式完成服务。Pod有如下特征：
  + 包含过个共享IPC、Network和UTC namespace，直接通过localhost通信，每个pod有唯一的IP
  + pod内容器可以访问共享的Volume，Pod优雅终止
  + 具有改变系统配置的权限，网络插件中有应用
  + 具有重启策略、拉取策略、资源限制、健康检查
  + 可以通过容器生命周期钩子函数用于监听容器生命周期特定事件
+ **Namespace**：对一组资源和对象的抽象集合，比如可以用来将系统内部的对象划分为不同的项目组或者用户组。常见的pod、service、replicaSet和deployment等都是属于某一个namespace的(默认是default)，而node, persistentVolumes等则不属于任何namespace。
  + namespace可以进行查询、创建、删除，default 和 kube-system 命名空间不可删除。Events是否属于namespace取决于产生events的对象。
+ **Node**：可以是物理机也可以是虚拟机。Node本质上不是Kubernetes来创建的， Kubernetes只是管理Node上的资源。为了管理Pod，每个Node节点上至少需要运行container runtime（Docker）、kubelet和kube-proxy服务。
  + 常见操作有查询、cordon设置调度，taint设置
+ **service**：Service是对一组提供相同功能的Pods的抽象，并为他们提供一个统一的入口，借助 Service 应用可以方便的实现服务发现与负载均衡，并实现应用的零宕机升级。Service通过标签(label)来选取后端Pod，一般配合ReplicaSet或者Deployment来保证后端容器的正常运行。
  + ClusterIP: 默认类型，自动分配一个仅集群内部可以访问的虚拟IP
  + NodePort: 在ClusterIP基础上为Service在每台机器上绑定一个端口，这样就可以通过 NodeIP:NodePort 来访问该服务
  + LoadBalancer: 在NodePort的基础上，借助cloud provider创建一个外部的负载均衡器，并将请求转发到 NodeIP:NodePort
  + ExternalName: 将服务通过DNS CNAME记录方式转发到指定的域名
+ **Volume存储卷**：Kubernetes提供了更强大的Volume机制和插件，解决了容器数据持久化以及容器间共享数据的问题。Kubernetes存储卷的生命周期与Pod绑定。
  + 支持emptyDir\hostPath\NFS\secret\persistentVolumeClaim等
  + PersistentVolume(PV)是集群之中的一块网络存储。跟 Node 一样，也是集群的资源。PersistentVolume (PV)和PersistentVolumeClaim (PVC)提供了方便的持久化卷，
+ **Deployment 无状态应用**：针对无状态类型的应用，Kubernetes使用Deloyment的Controller对象与之对应。其典型的应用场景包括：定义Deployment来创建Pod和ReplicaSet、滚动升级和回滚应用、扩容和缩容、暂停和继续Deployment
+ **StatefulSet 有状态应用**:StatefulSet则是为了有状态服务而设计，其应用场景包括：稳定的持久化存储，稳定的网络标志，有序部署，有序扩展，有序收缩，有序删除。
+ **DaemonSet 守护进程集**：DaemonSet保证在特定或所有Node节点上都运行一个Pod实例，常用来部署一些集群的日志采集、监控或者其他系统管理应用。如日志收集、系统监控、系统程序，目前支持OnDelete和RollingUpdate。
+ **Ingress**：Service：使用Service提供集群内部的负载均衡，Kube-proxy负责将service请求负载均衡到后端的Pod。Ingress Controller：使用Ingress提供集群外部的负载均衡。
  + Service和Pod的IP仅可在集群内部访问。集群外部的请求需要通过负载均衡转发到service所在节点暴露的端口上，然后再由kube-proxy通过边缘路由器将其转发到相关的Pod，Ingress可以给service提供集群外部访问的URL、负载均衡、HTTP路由等，为了配置这些Ingress规则，集群管理员需要部署一个Ingress Controller，它监听Ingress和service的变化，并根据规则配置负载均衡并提供访问入口。
+ **HPA（Horizontal Pod Autoscaling）水平伸缩**：Horizontal Pod Autoscaling可以根据CPU、内存使用率或应用自定义metrics自动扩展Pod数量 (支持replication controller、deployment和replica set)。
+ **Secret 密钥**：Sercert-密钥解决了密码、token、密钥等敏感数据的配置问题，而不需要把这些敏感数据暴露到镜像或者Pod Spec中。Secret可以以Volume或者环境变量的方式使用。有如下三种类型：kubernetes.io/dockerconfigjson: 用来存储私有docker registry的认证信息。
+ **ConfigMap 配置中心**：ConfigMap用于保存配置数据的键值对，可以用来保存单个属性，也可以用来保存配置文件。ConfigMap跟secret很类似，但它可以更方便地处理不包含敏感信息的字符串。ConfigMap可以通过三种方式在Pod中使用，三种分别方式为:设置环境变量、设置容器命令行参数以及在Volume中直接挂载文件或目录。
+ **Resource Quotas 资源配额**：资源配额(Resource Quotas)是用来限制用户资源用量的一种机制。资源配额有计算、存储、对象数限制。资源配额应用在Namespace上，并且每个Namespace最多只能有一个 ResourceQuota 对象。开启计算资源配额后，创建容器时必须配置计算资源请求或限制(也可以 用LimitRange设置默认值)。用户超额后禁止创建新的资源。

#### 2、controller和Service
+ K8s通过Controller来管理Pod，定义了Pod部署特性，使用时通过ReplicaSet来管理多个副本的。
+ service定义了外界访问一组特定Pod的方式，Service有自己的IP和端口，service为Pod提供了负载均衡。
+ 部署K8S集群：
  + 1-安装docker
  + 2-安装kubelet\kubeadm\kubectl
  + 3-用kubeadm创建Cluster
  + 4-初始化master、配置kubectl、安装Pod网络、添加节点

#### 3、K8s架构
* **master**：运行kube-apiserver、kube-scheduler、kube-controller-manager、etcd和pod网络。
  * API server提供RESTFul API，管理cluster的各种资源
  * Scheduler负责决定Pod在哪个node运行，考虑应用需求
  * control manager负责管理资源，有副本、部署、状态集、守护进程、命名空间等
  * etcd负责保存集群配置信息和各种资源状态信息。
  * Pod网络协助pod相互通信。
* **Node节点**：运行的组件有kubelet\kube-proxy\Pod网络
  * kubelet 根据调度信息创建和运行容器，向master报告状态
  * kube-proxy：负责将访问service的数据流转发到后端的容器，会实现负载均衡
  * Pod网络：保证Pod可以相互通信。
> kubelet是唯一没有以容器形式运行的K8s组件，通过systemd服务运行，其他组件以Pod形式运行。

#### 4、运行应用
总结过程为：
（1）用户通过kubectl 创建deployment
（2）Deployment创建Replicaset
（3）ReplicaSet创建Pod。
可以通过命令行创建，也可以通过kubectl apply -f nginx.yml创建。
kubectl可以通过label实现Pod指定到某个节点上。
* DaemonSet：每个节点上最多只能运行一个副本。
  (1) 配置方法与deployment类似，只是kind设置为DaemonSet。
  (2) hostNetwork可以指定Pod使用的网络
  (3) Containers定义运行服务的两个容器
* kube-proxy:指定类型资源，定义了容器，可以通过如下命令查看配置和运行状态。
```shell
kubectl edit deployment zyong-test-nginx
```
* Job：K8s的部署、ReplicaSet和DaemonSet用户管理服务类容器，对于工作烈日容器，使用Job。
  * job的重启策略会影响Pod的数量和Pod的重启次数。
  * 可以配置Cronjob，需要在kube-apiserver中进行修改，通过kubectl api-versions进行确认。
#### 5、通过service访问Pod
每个Pod有自己Ip地址，Pod不健壮，新Pod替代故障Pod时，是通过一组Pod对外提供服务，解决方案是service。
cluster内部通过kubernetes这个service访问Kubernetes API Server。
Cluster-ip通过iptables映射到Pod IP。访问时，基于配的规则按概率轮询访问，每一个节点都配置了相同的iptables规则。
cluster中的Pod可以通过<service_name>.<namespace_name> 来访问service。
* 外网如何访问service？
  * 可以通过节点的静态端口提供服务，Cluster外部可以通过<nodeIp>:<NodePort>访问Service。也可以利用LB对外提供服务，目前支持的有GCP、AWS、Azur。
#### 6.Rolling Update
滚动更新好处就是零停机，始终有副本在运行，保证业务的连续性。
每次更新仍然是创建ReplicaSet进行滚动更新，增加一个新Pod，减少一个旧的Pod。
* K8s提供maxSurge和maxUnavailable来精细控制Pod的替换数量。
```
# kubectl apply部署并更新应用
kubectl apply -f httpd.v3.yml --record
# 查看revision历史记录
kubectl rollout history deployment httpd
# 回滚到最初的版本
kubectl rollout undo deployment httpd --to-revision=1
```
* maxSurge：此参数控制更新过程中的副本总数超过期望的上限，向上取整。
* maxUnavailable：不可用的副本占期望的最大比例，向下取整。
* Liveness：判断进程退出码是否非0，非0认为容器故障，根据restartPolicy重启容器。
* Readiness探测：告诉K8S什么时候可以将容器加入service中对外提供服务。
  * 这两种探测均属于健康检查机制，liveness则是重启容器，readiness则是设置为不可用，二者之间没有依赖。
#### 7.数据管理
volume被mount到某个Pod时，pod中的所有容器都可以访问。
* emptyDir:不具备持久性，临时存储空间。
* hostPath：mount到host上，如果host崩溃，就无法访问了。
* 外部Storage provider：可以存到云厂商或者相应地方。
* PV：外部存储系统中的一块存储空间，PV由管理员创建和维护，PVC是对PV的申请，由普通用户创建和维护。PV提前创建，然后PVC申请PV在Pod中使用，叫做静态供给。动态供给可以通过StorageClass实现，AWS 的EBS有gp2和io1两种类型。
#### 8.secret和Configmap
- secret以密文形式存储数据，会以Volume的形式被mount到pod,容器可以通过文件或者环境变量的方式使用这些数据。
  - 文件方式支持动态更新，环境变量方式不支持。
- 应用的非敏感信息可以通过configmap配置。
#### 9.Helm-K8s包管理器
解决service\secrect\pvc\deployment集成部署的问题，定义应用与服务，以及服务与服务之间的依赖关系。支持参数化配置和多环境部署。
* helm架构：包含chart与release，chart是应用部署的自包含逻辑单元，release是chart运行实例。主要包含helm client和tiler server，server和kube api通信。
* chart：chart目录结构可以打包，包含chart yaml-描述概要信息，README.md-使用文档，requirements.yaml指定依赖关系，values.yaml-提供配置参数的默认值，templates目录-各类资源的配置模板。
  * 模版将chart参数化了，可以灵活定制应用。
#### 10、k8s网络
k8s采用扁平地址空间的网络模型，每个pod都有自己的Ip地址，Pod不需要配置NAT就能直接通信。service提供了高可用和负载均衡功能，稳定的前段对外提供服务，同时请求转发给正确的pod。
#### 11.监控与日志
* 基于web的Dashboard。用户可以部署容器化的应用、监控应用的状态、执行故障排查任务，以及管理K8s各种资源。
* weave scope和Heapster主要监控对象是Node和pod，这些数据对运维人员必需。
* Prometheus operator：可以监控K8S集群状态，API server，Scheduler， Controller manager，通过grafana展示监控数据。架构有：Prometheus Server，Exporter，可视化组件，Alertmanager等。
* 集群日志管理：K8s->fluentd->Elasticsearch->Kibana.

