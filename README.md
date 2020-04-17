# CyBuf-go
a go package for marshal&unmarshal CyBuf format data

# Usage
CyBuf looks like JSON, there are three differences:
- CyBuf need not outermost braces (expect in CyBuf stream)
- CyBuf split attributes by space characters(space,line break,tab...)
- Attributes of CyBuf have no double quotes

A standard CyBuf format data:
```yaml
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
			Name: "Acm"
			Phone: 2333
		}
	]
}
```
