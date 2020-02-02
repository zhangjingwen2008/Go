
//remixd命令：remixd -s E:\Dapps\contracts --remix-ide http://remix.ethereum.org
//solidity网络教程文档：https://bwcs.gitee.io/02_eth_basic/


//指定编译器版本，版本标识符
pragma solidity ^0.4.17;

//关键字 contract 跟java的class⼀样 智能合约是Inbox
contract Inbox{
    //string 是数据类型， message是成员变量，在整个智能合约⽣命周期都可以访问
    //public 是访问修饰符，是storage类型的变量，成员变量和是全局变量
    string public message;

    //函数以function开头，构造函数
    function Inbox (string initMessage) public {
        //本地变量
        var tmp = initMessage;
        message = tmp;
    }

    //view是修饰符，表示该函数仅读取成员变量，不做修改
    function getMessage() public view returns(string) {
        return message;
    }
}