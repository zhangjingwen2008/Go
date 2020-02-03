//1.导入合约实例

//2.读取数据

//3.写入数据

//4.读取数据

let instance=require('./532. 智能合约 instance')
const from='0xBc278edab4B690C7Aa594d4a8F4353D54796491C'

//异步调用，返回值是一个promise
//2.读取数据
/*
instance.methods.getValue().call().then(data => {
    console.log('data:', data)

    //3.写入数据
    instance.methods.setValue('Hello HangTou').send({
        from: from,
        value: 0,
    }).then(res => {
        console.log('res:', res)

        //4.读取数据
        instance.methods.getValue().call().then(data => {
            console.log('data2:', data)
        });
    });
})
 */

//web3与区块链交互的返回值都是promise，可以直接使用async/await
let test = async () => {
    try {
        let v1 = await instance.methods.getValue().call()
        console.log('v1:', v1)

        let res = await instance.methods.setValue('Hello Hangtou').send({
            from: from,
            value: 0,
        })
        console.log('v1:', v1)

        let v2 = await instance.methods.getValue().call()
        console.log('v2:', v2)
    } catch (e) {
        console.log(e)
    }
};
test()