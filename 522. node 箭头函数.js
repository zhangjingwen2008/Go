
function Add(a,b) {
    return a + b
}
let c = Add(1, 2)
console.log('c:',c)

//箭头函数
let add = (a, b) => a + b
console.log('d:', add(1, 2))

//默认值
//1.函数支持默认值
//2.如果有默认值，最好从参数最右边开始往左填
function print(name, address = '靓仔地方') {
    console.log(`name=${name},address=${address}`)
}
print('Hunter')
print('Sally', 'Shenzhen');