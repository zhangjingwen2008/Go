

let fs = require("fs")

//把异步读取文件的过程封装成一个promise
let readFilePromise = new Promise(function (resolve/*成功时调用*/, reject/*失败时调用*/) {
    fs.readFile('./1.txt', 'utf-8',function (err, data) {
        if (err){
            reject(err)     //出错，调用reject
        }
        resolve(data)       //成功，调用resolve
    });
})

//第一次改写，使用then方式调用
readFilePromise.then(res => {
    console.log('data:', res)
}).catch(err => {
    console.log(err)
})




