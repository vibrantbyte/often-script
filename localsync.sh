#!/bin/bash

## 同步到挂载的磁盘上
## -P 显示拷贝进度
## -l 如果文件是软链接文件，则拷贝软链接本身而非软链接所指向的对象。
## -z 传输时压缩提高效率
## -t 保持mtime属性。强烈建议任何时候都加上"-t"，否则目标文件mtime会设置为系统时间，导致下次更新
## --exclude 指定排除规则来排除不需要传输的文件。
## –daemon 以守护进行同步
rsync -a -t -P -l --exclude=.* /home/chao/ /media/chao/share/linux-chao-back
