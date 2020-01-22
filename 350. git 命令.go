package main

/*
	安装与配置
		- sudo apt-get install git
		- git

	创建版本库
		- mkdir Liangzai
		- cd Liangzai
		- git init

	第一次提交文件
		- git add Liangzai.txt
		- git commit -m "第一次提交，这里写的是备注信息，用-m触发"
		- git log		//可以查看到提交信息

	第二次提交文件，即更新
		- git add Liangzai.txt			//此时txt文件是已修改过的
		- git commit -m "第二次提交"
		- git log

	回退版本
		- git reset --hard HEAD^		//一个尖号^即回退到 当前 的上一个版本，两个尖头就是两个版本
		- git reset --hard HEAD~100		//若要回退100个版本，不可能写100个尖号，所以还可以用波浪号来指定
		- git reset --hard HEAD~100		//若要回退100个版本，不可能写100个尖号，所以还可以用波浪号来指定
		例子：回退到第1版本
			1.git log					//获得第2版本的版本号（复制一长串版本号前面的随便几个字符即可）
			2.git reset --hard [刚刚复制的版本号]
			3.git reset --hard HEAD^ 	//完成回退
		- git reflog					//显示所有git操作日志

	查看状态
		- git status		//可查看分支位置，查看文件的变更等信息
	查看git操作日志
		- git log

	文件撤销（把已添加的文件撤回）
		- git checkout -- Liangzai.txt			//工作区的撤销
		- git reset HEAD Liangzai.txt			//暂存区的撤销

	文件版本比较
		- 方式一：git diff HEAD -- Liangzai.txt				//仓库与当前版本比较
		- 方式二：git diff HEAD HEAD^ -- Liangzai.txt		//当前与上一版本比较

	查看分支
		- git branch
	创建分支
		- git branch LiangzaiFenzhi
	切换分支
		- gie checkout LiangzaiFenzhi
	创建并跳转分支
		- git checkout -b Liangzaifenzhi02
	删除分支
		- git branch -d Liangzaifenzhi
	合并分支
		- （在master分支里） git merge LiangzaiFenzhi 		//快速合并，靠移动指针HEAD完成。若出现冲突，则要手动修正
	图形化显示分支记录
		- git log --pretty=oneline --graph

	解决bug的步骤
		- git stash		//保存当前工作目录和索引状态
		- git checkout -b bug01			//新建并跳转一个分支用来进行bug处理
		- 解决bug ing....
		- git add 修改的文件
		- git commit -m "bug的解决"
		- git checkout master
		- git stash pop 		//恢复到之前的工作状态

	新环境拉取代码编写后，再上传到github
		- git clone [git项目地址.git]
		- git checkout -b newLiangzaiBranch		//构建新分支并跳转
		- 编写内容.....
		- git add newLiangzai.txt
		- git commit -m "添加新内容到github"
		- git push origin newLiangzaiBranch		//将分支上传

	将本地分支与远程服务器分支关联
		- git branch --set-upstream-to=origin/[服务器分支名] [本地分支名]
		- git add liangzai.txt
		- git commit -m "关联后的git提交"

	从远程分支上拉取代码到本地所属分支
		- git pull origin [分支名称]
*/
