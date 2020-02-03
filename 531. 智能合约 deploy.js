let {bytecode,interface}=require('./530. 智能合约 compile')

//1.引入web3
let Web3 = require('web3')
//2.new 一个web3实例
let web3=new Web3()
//3.设置网络
web3.setProvider("http://127.0.0.1:7545")               //Ganache中的ip

const account='0x8003C07343524a4b760F1D3b9dBa5Ea8a92fA5Bf'      //Ganache中的账户

console.log('version:',web3.version)

//1.拼接合约数据 interface
let contract=new web3.eth.Contract(JSON.parse(interface))

//2.拼接bytecode
contract.deploy({       //3.合约部署
    data: bytecode,                      //合约的bytecode
    arguments: ['HelloWorld']            //给构造函数传递参数，使用数组
}).send({
    from: account,
    gas: '3000000',             //不要用默认值，一定要写大一些，要使用单引号
    //gasPrice: '1',
}).then(instance => {           //返回合约实例
    console.log('address:',instance.options.address)
});

