package main

import (
	"fmt"
	"time"
)

type Base struct {
	Id   int
	Name string
}

//博主类
type Blogger struct {
	Base
	WeiBos   []*PostContent
	Comments map[int][]*PostContent //key:针对哪一条微博进行评论的Id；value：具体评论内容
	Fans     []FansInterface
}

//定义一个函数，完成博主类对象的创建
func NewBlogger(name string) *Blogger {
	//1.创建博主对象
	blg := new(Blogger)
	//2.初始化博主
	blg.Name = name
	blg.WeiBos = make([]*PostContent, 0)
	blg.Comments = make(map[int][]*PostContent)
	return blg
}

//发布微博方法的实现
func (b *Blogger) PostWeiBo(content string, wbType int) {
	//1.创建PostContent对象
	weibo := new(PostContent)
	//2.成员进行初始化
	weibo.Id = b.GetId()
	weibo.Content = content
	weibo.Type = wbType
	weibo.CommentTime = time.Now()
	weibo.PostMan = b.Name
	weibo.To = "All"
	//3.存储微博数据
	b.WeiBos = append(b.WeiBos, weibo)
	b.Notify(weibo.Id)
}

//获取微博编号
func (b *Blogger) GetId() int {
	//将微博最后一条id进行加一操作即可
	if len(b.WeiBos) == 0 {
		return 0
	} else {
		return b.WeiBos[len(b.WeiBos)-1].Id + 1
	}
}

type BloggerInterface interface {
	//粉丝关注博主
	Attach(bFans FansInterface)
	//粉丝取消关注博主
	Detach(bFans FansInterface)
	//发布微博后，通知粉丝
	Notify(wbid int)
}

//发送通知
func (b *Blogger) Notify(wbid int) {
	//1.遍历获取每个粉丝数据
	for _, value := range b.Fans {
		value.Update(b, wbid)
	}
	//2.发送通知
}

//粉丝类
type Fans struct {
	Base
}

//粉丝操作的接口
type FansInterface interface {
	//接收博主发出的通知
	Update(bloggerI BloggerInterface, wbid int)
	//具体操作的方法，例如：发布评论
	Action(bloggerI BloggerInterface, wbid int)
}

//获取博主新发布的微博
func (b *Blogger) GetWeiBo(wbid int) (content *PostContent) {
	for _, value := range b.WeiBos {
		if value.Id == wbid {
			content = value
			return
		}
	}
	return nil
}

//友好的粉丝
type FriendFans struct {
	Fans
}

func (f *FriendFans) Update(bloggerI BloggerInterface, wbid int) {
	fmt.Printf("你好%s,你所关注的博主发布了新的微博\n", f.Name)
	f.Action(bloggerI, wbid)
}
func (f *FriendFans) Action(bloggerI BloggerInterface, wbid int) {
	//1.获取博主发布的微博
	blogger, ok := bloggerI.(*Blogger)
	if ok {
		weibo := blogger.GetWeiBo(wbid)
		//2.进行评论
		//2.1 构建评论内容
		cType := weibo.Type
		message := ""
		switch cType {
		case 1:
			message = "非常好呀"
		case 2:
			message = "加油！"
		}
		postComment := PostContent{0, message, time.Now(), cType, f.Name, blogger.Name}
		//2.2 发布评论
		blogger.AddComment(postComment, wbid)
		blogger.ShowComment(wbid)
	}
}

//发布评论
func (b *Blogger) AddComment(comment PostContent, wbid int) {
	b.Comments[wbid] = append(b.Comments[wbid], &comment)
}

//展示评论
func (b *Blogger) ShowComment(wbid int) {
	//1.根据微博id，获得微博数据
	blog := b.GetWeiBo(wbid)
	fmt.Println("博主名称：", blog.PostMan)
	fmt.Println("微博内容：", blog.Content)

	//2.展示微博对应评论内容
	for _, value := range b.Comments[wbid] {
		fmt.Println("评论人：", value.PostMan)
		fmt.Println("评论内容：", value.Content)
	}
}

//不友好的粉丝
type BadFans struct {
	Fans
}

func (f *BadFans) Update(bloggerI BloggerInterface, wbid int) {
	fmt.Printf("你好%s,你所关注的博主发布了新的微博\n", f.Name)
}
func (f *BadFans) Action(bloggerI BloggerInterface, wbid int) {
	//1.获取博主发布的微博
	blogger := new(Blogger)
	weibo := blogger.GetWeiBo(wbid)
	fmt.Println(weibo)
}

//粉丝关注博主
func (b *Blogger) Attach(bFans FansInterface) {
	b.Fans = append(b.Fans, bFans)
}

//粉丝取关博主
func (b *Blogger) Detach(bFans FansInterface) {
	for i := 0; i < len(b.WeiBos); i++ {
		if b.Fans[i] == bFans {
			//移除切片中的数据
			b.Fans = append(b.Fans[:i], b.Fans[i+1:]...)
		}
	}

}

//微博&评论
type PostContent struct {
	Id          int       //编号
	Content     string    //内容
	CommentTime time.Time //时间
	Type        int       //类型
	PostMan     string    //发布人
	To          string    //给谁发布的（博主姓名）
}

func main() {
	blg := NewBlogger("张三")
	friendFans := new(FriendFans)
	friendFans.Id = 1
	friendFans.Name = "李四"
	blg.Attach(friendFans)
	blg.Detach(friendFans)
	for _, value := range blg.Fans {
		fmt.Println(value)
	}
	blg.PostWeiBo("Good Day!", 1)
}
