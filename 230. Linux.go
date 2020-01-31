package main

/*
	Ubuntu 18.04 下载：https://www.ubuntu.com/download/desktop

	图形界面->命令行界面：Ctrl+Alt+F3
	命令行界面->图形界面：Ctrl+Alt+F1
	开启命令行端口：Ctrl+Alt+T
	关闭命令行端口：Ctrl+D
	在同一个终端内开启多个标签：Ctrl+Shift+T

	进入某个目录（change directory）：cd			// cd - ：切换到上一个工作的目录
	清空命令行：Ctrl+L
	切换成root用户：sudo su		//退出root用户：exit
	切换其他用户：su [用户名]
	历史命令：history
	显示当前所在目录：pwd
	查看文件信息：file [文件名]
	查看命令所在位置：which [命令例如ls]
	查看当前登录用户名：whoami
	更改当前用户密码：passwd
	Linux下的任务管理器：top
	重启操作系统：reboot
	关机：init 0
	重启：init 6

	查看文件内容：cat [文件名]
	查看文件内容（分页）：more [文件名]
	查看文件内容（可输入命令）：less [文件名]
	查看文件内容（编译器访问）：vi [文件名]
	查看文件内容（弹出编辑器可编辑）：gedit [文件名]

	管道：ls -l | more 		【一个竖线就是管道，即将管道 左边的命令执行结果写给右边，再执行右边的命令】
	筛选：ls -l | grep if	【这里的话，筛选字符串if将会在输出中高亮显示】

	列举系统中所有进程：ps
	列举系统中所有进程（最常用）：ps -aux | grep xxx
		- a：所有进程
		- u：当前用户使用的进程
		- x：系统执行的进程
		- 这里的命令意思是，查找名字为xxx的进程

	查看详细信息（-a查看以“.”开头的隐藏的文件）：ls -la
	查看详细信息（-d查看文件夹）：ls -ld
	查看详细信息：ls -l [可指定对象名]
		【结果共10个字符】-rwxr-xr-x 1 lbmdlc120 root 174216 4月  28  2017 zipinfo
		1.【-】代表文件类型，Linux下共7种类型
			1.普通文件：-
			2.目录文件：d
			3.软链接文件：l
			4.字符设备文件：c
			5.块设备文件：b
			6.套接字文件：s
			7.管道文件：p
			8.unknown文件
		2.【rwxr-xr-x】9个字符，分为3组
			1.三个组分别为
				- 文件所有者
				- 文件所属组
				- 其他人
			2.每个组内的三个字母分别表示
				- r：读取。计数为4
				- w：写入。计数为2
				- x：执行。计数为1
				- 根据计数可以得到当前文件的权限是：755（7：rwx；5：r-x，5：r-x）
		3.【1】硬链接计数
			- 即有多少个文件共用同一个inode编号
			- 创建硬链接文件可以做到两个文件的同步
			- 特征：文件和硬链接文件之间，除文件名不一样之外，其他信息完全一致，并能实时同步
			- 创建硬链接：ln [旧文件] [新硬链接文件]
		4.【lbmdlc120】文件所属用户，谁创建默认属于谁
		5.【root】文件所属用户组
		6.【174216】文件所占空间大小。目录文件：4k整数倍
		7.【4月  28  2017】创建时间/最后一次修改的时间
		8.【zipinfo】当前查看的文件名

	将打印输出到指定文件里（覆盖）：ls -l test > out
	将打印输出到指定文件里（追加）：ls -l test >> out

	硬链接：ln 	 旧文件 新文件	【不能给目录创建硬链接】
	软链接：ln -s 旧文件 新文件 	【相当于windows快捷方式】

	移动文件：mv 源文件 目的位置
	改名文件：mv 源文件 不存在的文件/目录名

	压缩文件（tar）：tar -zcvf xxx.tar.gz			//z：gzip格式；c：创建压缩文件；v：输出压缩详细；f：指定压缩后的文件名
	压缩文件（tar）：tar -jcvf xxx.tar.bz2		//j：bzip2格式
	解压文件（tar）：tar -zxvf xxx.tar.gz
	解压文件（tar）：tar -jxvf xxx.tar.bz2
	- 压缩文件（zip）：zip -r 压缩包名(无.zip后缀)
	- 解压文件（zip）：unzip -d 解压缩位置 xxx.zip(待解压缩文件)
	压缩文件（rar）：rar a -r 压缩包名.rar(后缀可不加) 打包压缩材料
	解压文件（rar）：rar x 待解压缩包名.rar

	添加用户：  sudo adduser [新用户名]
	删除用户：  sudo deluser [用户名]
	添加用户组：sudo addgroup [新用户组名]
	删除用户组：sudo delgroup [用户组名]

	修改文件所拥有者(change owner)：sudo chown [新持有者] [要修改的文件名]
	修改文件所拥有组(change group)：sudo chgrp [新持有组] [要修改的文件名]
	修改文件权限(change mode)：	  sudo chmod [权限计数] [文件名]			【权限计数：例如536，代表r-x-wxrw-】

	杀死终止进程：		 kill -9 [要终止的进程号]
	查看终止进程的所有命令：kill -l

	放置在后台运行（&）：firefox www.baidu.com &		【防止在命令行卡死】
	查看后台运行的进城：jobs
	恢复后台一个进程到前台执行：fg [使用job查看的编号]

	启动/关闭网卡：sudo ifconfig [网卡名] up/down
	查看ip：ifconfig
	修改ip（重启后失效）：sudo ifconfig [网卡名] [新ip地址]			【网卡名：例如ens33或eth0.】
	修改ip（永久-图形界面）：设置——网络——有线连接 + ——ipv4——手动——填入地址、子网掩码、网关——“添加”
	修改ip（永久—命令界面）：修改 /etc/network/interfaces文件，添加网址、子网掩码、网关，指定DNS服务器

	测试某软件是否安装：sudo aptitude show [软件名称]
		   - 若没安装：sudo aptitude install [软件名称]

	Linux远程登录Linux：
			- 方式1：ssh -l 目标主机用户名 目标主机IP地址
			- 方式2：ssh [目标主机用户名]@[目标主机IP地址]		【推荐】
			- 退出连接：exit
	Linux远程发送文件给Linux：
			- 超级拷贝 super copy
			- 命令：scp -r [目标用户名]@[目标主机名]：[目标存储绝对路径]
			- 例子：scp -r ./test123 Liangzai@127.0.0.1:/home/Liangzai/test
*/
