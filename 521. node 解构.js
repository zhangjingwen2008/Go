
//1.对数组进行解构
let arr1=[0,1,2,3,4,5]
console.log('arr1[0]',arr1[0])

let[a,b,c,d]=arr1
console.log(a,b,c,d)

//2.对对象进行解构
const person={
    name:'Hunter',
    age:23,
    address:'深圳'
}

//自动推导，一定要写成同名，否则要自己指定
let {name:liangzai,age,address}=person
console.log(liangzai,age,address)

//函数参数解构对象
function printPerson({name,age}){
    console.log('姓名：',name,' 年龄：',age)
}
printPerson(perso)