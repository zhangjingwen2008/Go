
//导入solc编译器
let solc = require('solc')

//读取合约
let fs=require('fs')
let sourceCode=fs.readFileSync('./contracts/SimpleStorage.sol','utf-8')

//编译合约
// Setting 1 as second paramateractivates the optimiser
let output = solc.compile(sourceCode, 1)

console.log('abi:',output['contracts'][':SimpleStorage']['interface'])

/*
abi: [{"constant":true,"inputs":[],"name":"getValue","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_str","type":"string
"}],"name":"setValue","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_str","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructo
r"}]
*/

module.exports=output['contracts'][':SimpleStorage']

