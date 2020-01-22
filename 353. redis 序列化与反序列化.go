package main

/*
	序列化与反序列化
		序列化（字节化）
			- var buffer bytes.Buffer			//容器
			- enc := gob.NewEncoder(buffer)		//编码器
			- err := enc.Encode(dest)			//编码
		反序列化（反字节化）
			- dec := gob.NewDecoder(bytes.NewReader(buffer.bytes()))		//解码器
			- dec.Decode(src)					//解码
*/
