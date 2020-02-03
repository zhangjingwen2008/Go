
//node内置读取文件模块
let fs=require('fs')
let filename='1.txt'

//同步读取文件
let data=fs.readFileSync(filename,'utf-8')
console.log('同步读取文件内容data：',data)

//异步读取文件
fs.readFile(filename, 'utf-8',/*回调函数*/function (err, data) {
    if (err){
        console.log('读取文件出错', err)
        return
    }
    console.log('异步读取文件数据：', data)
});

console.log('异步读取数据222')