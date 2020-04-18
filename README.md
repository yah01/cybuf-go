# CyBuf-go
a go package for marshal&unmarshal CyBuf format data

## Usage
CyBuf looks like JSON, there are three differences:
- Outermost braces are optional, expect in CyBuf stream
- CyBuf split attributes by space characters(space,line break,tab...), not comma.
- Attributes of CyBuf have no double quotes, expect for attribute names that contain control characters.

A standard CyBuf format data:
```javascript
{
	Name: "cybuf"
	Age: 1
	Weight: 100.2
	School: {
		Name: "Wuhan University"
		Age: 120
	}
	Friends: [
		{
			Name: "Zerone"
			Phone: 01010101
		}
		{
			Name: "ACM"
			Phone: 2333
		}
	]
}
```

## How to contribute
Just contact me, there are something need to do:
- Design error types and messages
- Marshal struct
- Unmarshal struct
- Zip/Unzip a cybuf data (go to [cybuf-formatter](https://github.com/yah01/cybuf-formatter) repo)
- Support attribute names containing control characters
- Support custom Marshal()/Unmarshal() methods
- Support unmarshal() from io.Reader(bytes stream)
- Optimize the algorithms

There's no CyBuf support for other languages, they need you:
- [cybuf-py](https://github.com/yah01/cybuf-py)
- [cybuf-cpp](https://github.com/yah01/cybuf-cpp)
- [cybuf-rust](https://github.com/yah01/cybuf-rust)
- [cybuf-java](https://github.com/yah01/cybuf-java/blob/master/Cybuf.java)
- ...