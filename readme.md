# md-tidy

清除目录下没有被引用的图片、静态资源，节约存储空间，降低管理成本，图片检查范围和引用的文本范围均以当前目录为准。

# how to use

```bash
git clone 
go install .
md-tidy -check # 仅执行检查，不执行删除操作
md-tidy # 删除未使用的图片数据
```
