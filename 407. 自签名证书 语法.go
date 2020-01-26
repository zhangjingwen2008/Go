package main

/*
	1. 使用openssl生成自签名证书

	   1. 创建一个目录如Mytest, 进入该目录, 在该目录下打开命令行窗口

	   2. 启动openssl
		  - openssl    //执行该命令即可

	   3. 使用openssl工具生成一个RSA私钥, 注意：生成私钥，需要提供一个至少4位的密码。
		  - genrsa -des3 -out server.key 2048   //des3: 使用3des对私钥进行加密

	   4. 生成CSR（证书签名请求）
		  req -new -key server.key -out server.csr

	   5. 删除私钥中的密码, 第一步给私钥文件设置密码是必须要做的, 如果不想要可以删掉
		  rsa -in server.key -out server.key		//out 参数后的文件名可以随意起

	   6. 生成自签名证书
		  x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

*/
