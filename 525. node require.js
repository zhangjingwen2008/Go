
//1.require可以导入包文件，可以是系统内置的，也可以使第三方的，也可以是自己实现的
//2.如果想require导入，那么一定要有对应的导出
//3.如果是自己写的模块，导入时要明确指定路径，文件名称不要加js

let ex=require('./06. node exports')

ex.printHello()
ex.testHello()
