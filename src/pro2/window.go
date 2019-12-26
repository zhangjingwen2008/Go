package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type Window interface {
	ShowWindow() //展示窗体界面
}

//压缩解压缩界面类
type ComWindow struct {
	Window
	*walk.MainWindow //主窗体
}

//压缩解压缩成功失败提示信息的界面类
type LabWindow struct {
	Window
}

//创建界面类对象
func Show(Window_Type string) {
	var Win Window
	switch Window_Type {
	case "main_window":
		Win = &ComWindow{}
	case "lab_window":
		Win = &LabWindow{}
	default:
		fmt.Println("参数传递错误")
	}
	Win.ShowWindow()
}

//实现ShowWindow方法，展示出空白的窗口
func (comWindow *ComWindow) ShowWindow() {
	pathWindow := new(ComWindow)
	err := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow,     //关联主窗体
		Title:    "文件压缩",                     //窗口标题名称
		MinSize:  declarative.Size{480, 230}, //窗口最小宽度和高度
	}.Create() //创建窗口
	if err != nil {
		fmt.Println(err)
	}

	//窗口的展示，需要通过坐标来指定
	pathWindow.SetX(650) //X坐标
	pathWindow.SetY(300) //Y坐标
	pathWindow.Run()     //运行窗口
}
