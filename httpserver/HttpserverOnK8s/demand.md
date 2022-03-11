一、功能需求：
* 优雅启动
* 优雅中止
* 资源需求和QoS
* 配置和代码分离
* 探活
* service
* ingress

二、实现：
1、优雅启动 
   * 探针:readinessProbe
    
2、优雅中止
   * 原理：
      容器终止流程
      Pod 被删除，状态置为 Terminating。
      kube-proxy 更新转发规则，将 Pod 从 service 的 endpoint 列表中摘除掉，新的流量不再转发到该 Pod。
      如果 Pod 配置了 preStop Hook ，将会执行。
      kubelet 对 Pod 中各个 container 发送 SIGTERM 信号以通知容器进程开始优雅停止。
      等待容器进程完全停止，如果在 terminationGracePeriodSeconds 内 (默认 30s) 还未完全停止，就发送 SIGKILL 信号强制杀死进程。
      所有容器进程终止，清理 Pod 资源。
   * 实现：程序对kill -15(SIGTERM)进行捕获后，停止接受新请求 

3、资源需求和QoS
* QoS 服务质量：资源不足时，驱逐策略
    * Guaranteed
        * Pod 中的每个容器都必须指定内存限制和内存请求。
        * 对于 Pod 中的每个容器，内存限制必须等于内存请求。
        * Pod 中的每个容器都必须指定 CPU 限制和 CPU 请求。
        * 对于 Pod 中的每个容器，CPU 限制必须等于 CPU 请求。
    * Burstable
        * Pod 不符合 Guaranteed QoS 类的标准。
        * Pod 中至少一个容器具有内存或 CPU 请求。
    * BestEffort
        * 没有设置内存和 CPU 限制或请求
    
4、配置和代码分离
* configmap、env
    
5、 探活:探活、可读、startupprobe
* LivenessProbe容器是否活着，如果没活着，删除重建。
* readinessprobe：容器是否就绪，标识它no ready，不会接收来自service的流量。
* startupprobe：类似于readinessprobe，只是在启动时，频率很低。
* 3种方式 http get，tcp socket、exec。
     
6、 service
* https://kubernetes.io/zh/docs/concepts/services-networking/service/
四层负载均衡,iptable, 1.11之后变成ipvs，都是netfilter框架的产物
* IPVS 提供了更多选项来平衡后端 Pod 的流量。 这些是：
  * rr：轮替（Round-Robin）
  * lc：最少链接（Least Connection），即打开链接数量最少者优先
  * dh：目标地址哈希（Destination Hashing）
  * sh：源地址哈希（Source Hashing）
  * sed：最短预期延迟（Shortest Expected Delay）
  * nq：从不排队（Never Queue）
* 四种类型
  * Cluster IP 集群内访问
  * NodePort 暴露给外部访问
  * loadBalancer 暴露给外部访问
  * ExternalName 类似cname，可以实现跨namespace的ingress

7、 ingress
* Ingress 是对集群中服务的外部访问进行管理的 API 对象，典型的访问方式是 HTTP。
* 负载均衡转发规则单一，推荐使用istio