package main

/*
协议：一组规则。要求使用协议的双方，必须严格遵守协议内容

网络分层架构：
    - OSI 七层模型：“物数网传会表应”
        （从底向上）
        1.物理层
        2.数据链路层
        3.网络层
        4.传输层
        5.会话层
        6.表示层
        7.应用层
    - TCP/IP 四层模型：“数网传应”
        （从底向上）
        1.数据链路层(1+2)
        2.网络层(3)
        3.传输层(4)
        4.应用层(5+6+7)

各层功能：
    链路层：ARP
        - 源mac地址————目标mac地址
        - ARP协议作用：借助IP获取mac地址
    网络层：IP
        - 源IP————目标IP
        - IP协议的作用：在网络环境中唯一标识一台主机
        - IP地址本质：2进制数。————点分十进制IP地址（string）
    传输层：TCP / UDP
        - port————在一台主机上唯一标识一个进程
    应用层：ftp、http、自定义
        - 对数据进行封装，解封装

数据通信过程：
    - 封装：数据 --> 应用层 --> 传输层 --> 网络层 --> 链路层
    - 解包：链路层 --> 网络层 --> 传输层 --> 应用层 --> 数据

通信过程所需元素：
    1.mac地址
        - 不需要用户指定，每个设备网卡都有一个唯一固定mac地址
        - 使用ARP协议，通过IP来获得mac地址
    2.IP地址
        - 需要用户指定，用来确认主机
    3.port端口号
        - 需要用户指定，用来确定程序
        - 不使用系统固定端口例如8080，尽量使用5000+的端口号，端口上限为65535
*/

func main() {

}
