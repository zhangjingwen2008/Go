
function printHello(){
    console.log('helloworld')
}
function testHello(){
    console.log('helloworld')
}

// module.exports=ex={          //export方式1
//     printHello,
//     testHello,
// }

let ex = {                        //export方式2
    printHello,
    testHello,
}
module.exports = ex