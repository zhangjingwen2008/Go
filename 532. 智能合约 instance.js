//获取合约实例，导出去

let {bytecode,interface}=require('./530. 智能合约 compile')

//1.引入web3
let Web3=require('web3')
//2.new 一个web3实例
let web3=new Web3()
//3.设置网络
web3.setProvider("http://127.0.0.1:7545")               //Ganache中的ip

//获取ABI和合约地址
let abi=[{"constant":true,"inputs":[],"name":"getValue","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_str","type":"string"}],"name":"setValue","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_str","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]
let address = '0x0634403150507DFeC6384c9e81d5681Bea43b21d'

//此处abi已经是json对象，故不需要进行JSON.parse转换
let contractInstance=new web3.eth.Contract(abi,address)     //得到合约实例

console.log('address:',contractInstance.options.address)

module.exports=contractInstance