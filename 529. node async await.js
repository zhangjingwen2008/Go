let fs = require('fs')

//此地狱回调解决办法：使用promise封装回调函数
let checkStat1 = () => {
    fs.readFile('./1.txt', 'utf-8', function (err, data) {
        console.log('读取文件:', data)
        fs.writeFile('./2.txt', 'utf-8', function (err) {
            if(err){
                return
            }
            console.log("写入成功")
            fs.stat('./2.txt', 'utf-8', function (err, stat) {
                if(err){
                    return
                }
                return stat
            })
        })
    })
}
// checkStat1()

//解决：
let readFilePromise = () => {           //原回调函数1
    return new Promise((resolve, reject) => {
        try {
            fs.readFile('./1.txt', 'utf-8', function (err, data) {
                console.log('读取文件:', data)
                resolve(data)
            });
        } catch (e) {
            reject(e)
        }
    })
}
let writeFilePromise = (data) => {           //原回调函数2
    return new Promise((resolve, reject) => {
        fs.writeFile('./1.txt', data, 'utf-8', function (err) {
            if (err) {
                reject(err);
            }
            resolve('写入成功');
        });
    });
};
let statPromise = () => {           //原回调函数3
    return new Promise((resolve, reject) => {
        fs.stat('./1.txt', function (err,stat) {
            if (err){
                reject(err)
            }
            console.log('文件状态：',stat)
            resolve(stat);
        });
    });
};
//如果想使用async，await，promise，
//调用函数的外面要修饰为async
//promise函数前面加上await
let checkStat2 = async () => {
    let data = await readFilePromise()
    let res=await writeFilePromise(data)
    let stat = await statPromise()
    console.log('stat:', stat)
};
checkStat2()
