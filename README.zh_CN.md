# CyBuf-go
一个用于序列化&反序列化CyBuf格式数据的golang包。
- [English](README.md)
- [中文](README.zh_CN.md)

## 用法
CyBuf和JSON类似，但它们两者之间主要有三个不同：

- 除了在CyBuf流内部，最外层的花括号是可选的；
- CyBuf属性之间的分隔通过空白字符进行构造（例如空格，换行符，制表符等），而非逗号；
- 除了含有控制字符的属性名外，CyBuf的属性没有双引号。

一份标准的CyBuf格式数据示例如下：
```javascript
{
    Name: "yah01"
    Age: 21
    Weight: 100.2
    Live: true
    Friends: [
        {
            Name: "wmx"
            Age: 100
            Weight: 200.5
            Live: false
            Friends: nil
            School: {
                Name: "SHU"
                Age: 114514
            }
        }
    ]
    School: {
        Name: "Wuhan University"
        Age: 120
    }
}
```

## 如何贡献你的代码
你只需要联系作者征得同意即可！

想要贡献你的代码，你可以从以下几个方面着手：

- 设计错误类型及其信息；

- 序列化结构；

- 反序列化结构；

- 将CyBuf格式数据压缩/解压缩（详情请移步 [cybuf-formatter](https://github.com/yah01/cybuf-formatter) 仓库）；

- 支持含有控制符的构造名

- 支持自定义的Marshal()和Unmarshal()方法；

- 支持基于io.reader的unmarshal()方法（字节流）；

- 对各类算法进行优化。

  

目前CyBuf不支持golang之外的其他语言。让CyBuf支持更多的语言，需要你的一臂之力：
- [cybuf-py](https://github.com/yah01/cybuf-py)
- [cybuf-cpp](https://github.com/yah01/cybuf-cpp)
- [cybuf-rust](https://github.com/yah01/cybuf-rust)
- [cybuf-java](https://github.com/yah01/cybuf-java/blob/master/Cybuf.java)
- ...