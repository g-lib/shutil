# shutil

shutil——port of [Python's `shutil`](https://github.com/python/cpython/blob/3.8/Lib/shutil.py)-[Py文档](https://docs.python.org/zh-cn/3/library/shutil.html)




## Fuction List:


**Directory and files operations**

+ `CopyFileObj(fSrc,fDst,length)`
+ `CopyFile(src,dst,followSymlinks bool)`
+ `CopyMode(src,dst,followSymlinks bool)`
+ `CopyStat(src,dst,followSymlinks bool)`
+ `Copy(src,dst,followSymlinks bool)`
+ `Copy2(src,dst,followSymlinks bool)`
+ `IgnorePatterns(...patterns)`
+ `CopyTree(src, dst, symlinks=False, ignore=None, copy_function=copy2, ignore_dangling_symlinks=False, dirs_exist_ok=False)`
+ `RmTree(path, ignore_errors=False, onerror=None)`
+ `Move(src, dst, copy_function=copy2)`
+ `DiskUsage(path)`
+ `Chown(path, user=None, group=None)`
+ `Which(cmd, mode=os.F_OK | os.X_OK, path=None)`

**Archiving operations**

+ `MakeArchive(base_name, format[, root_dir[, base_dir[, verbose[, dry_run[, owner[, group[, logger]]]]]]])`
+ `GetArchiveFormats()`
+ `RegisterArchiveFormat(name, function[, extra_args[, description]])`
+ `UnregisterArchiveFormat(name)`
+ `UnpackArchive(filename[, extract_dir[, format]])`
+ `RegisterUnpackFormat(name, extensions, function[, extra_args[, description]])`
+ `UnregisterUnpackFormat(name)`
+ `GetUnpackFormats()`

**Querying the size of the output terminal**


+ `GetTerminalSize(fallback=(columns, lines))`
