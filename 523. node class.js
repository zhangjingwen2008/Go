//父类
class Person {
    constructor(name, age) {
        this.name = name
        this.age = age
    }

    say() {
        console.log(`Hello, I'm ${this.name}, I'm ${this.age}.`)
    }
}
let p1=new Person('Father',50)
p1.say()


//子类
class XiaoDi extends Person{
    constructor(name, age) {
        super(name,age)
        this.name = name
        this.age = age
    }
}
let X1=new XiaoDi('Sally',11)
X1.say()