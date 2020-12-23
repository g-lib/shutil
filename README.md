# shutil

[English](./README_en.md)

shutil——[Python`shutil`](https://github.com/python/cpython/blob/3.8/Lib/shutil.py)-[PyDoc](https://docs.python.org/zh-cn/3/library/shutil.html)的Golang克隆迁移，提供了文件以及文件集合的高阶操作。 特别是提供了支持文件复制和删除的函数。相对于原Python的`shutil`，更添加了些许符合名称的功能函数。（shutil是shell utilities的缩写，意即shell工具包）




## 功能函数列表:


**目录与文件操作**

+ `CopyFileObj(fSrc,fDst,length)`：拷贝/复制文件对象
+ `CopyFile(src,dst,followSymlinks bool)`：拷贝/复制文件
+ `CopyMode(src,dst,followSymlinks bool)`
+ `CopyStat(src,dst,followSymlinks bool)`
+ `Copy(src,dst,followSymlinks bool)`
+ `Copy2(src,dst,followSymlinks bool)`
+ `IgnorePatterns(...patterns)`
+ `CopyTree(src, dst, symlinks=False, ignore=None, copy_function=copy2, ignore_dangling_symlinks=False, dirs_exist_ok=False)`
+ `RmTree(path, ignore_errors=False, onerror=None)`：移除/删除 目录
+ `Move(src, dst, copy_function=copy2)`：移动/剪切
+ `DiskUsage(path)`：获取磁盘使用情况
+ `Chown(path, user=None, group=None)`：更改某个目录的主人
+ `Which(cmd, mode=os.F_OK | os.X_OK, path=None)`：查询某个命令的程序所在路径

**文件包操作**

+ `MakeArchive(base_name, format[, root_dir[, base_dir[, verbose[, dry_run[, owner[, group[, logger]]]]]]])`：打包文件
+ `GetArchiveFormats()`：获取文件包格式
+ `RegisterArchiveFormat(name, function[, extra_args[, description]])`：注册文件包格式
+ `UnregisterArchiveFormat(name)`：注销文件包格式
+ `UnpackArchive(filename[, extract_dir[, format]])`：解包文件包
+ `RegisterUnpackFormat(name, extensions, function[, extra_args[, description]])`
+ `UnregisterUnpackFormat(name)`
+ `GetUnpackFormats()`：获取解包格式

**查询终端输出尺寸**


+ `GetTerminalSize(fallback=(columns, lines))`：获取终端(terminal)尺寸
