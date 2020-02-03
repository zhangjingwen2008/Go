
// fs.stat/fs.statSync：访问文件的元数据，比如文件大小，文件的修改时间
// fs.readFile/fs.readFileSync：异步/同步读取文件
// fs.writeFile/fs.writeFileSync：异步/同步写入文件
// fs.readdir/fs.readdirSync：读取文件夹内容
// fs.unlink/fs.unlinkSync：删除文件
// fs.rmdir/fs.rmdirSync：只能删除空文件夹，思考：如何删除非空文件夹？使用fs-extra 第三方模块来删除。
// fs.watchFile：监视文件的变化