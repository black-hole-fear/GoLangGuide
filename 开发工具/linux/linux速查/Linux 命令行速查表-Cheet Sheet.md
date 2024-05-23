# Linux 命令行速查表

 [bash](https://www.cheat-sheet.cn/tags/bash) [linux](https://www.cheat-sheet.cn/tags/linux)

Linux是一套免费使用和自由传播的类Unix操作系统，是一个基于POSIX和Unix的多用户、多任务、支持多线程和多CPU的操作系统。它能运行主要的Unix工具软件、应用程序和网络协议。它支持32位和64位硬件。Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统。

## 系统

|                 |                                    |
| :-------------- | :--------------------------------- |
| uname           | 显示linux系统信息                  |
| uname -r        | 显示内核版本信息                   |
| uptime          | 显示系统运行的时间（包括平均负载） |
| hostname        | 显示系统主机名                     |
| hostname -i     | 显示系统的IP地址                   |
| last reboot     | 显示系统重新启动历史记录           |
| date            | 显示当前系统日期和时间             |
| timedatectl     | 查询和更改系统时钟                 |
| cal             | 显示当前日历的月份和日期           |
| w               | 显示系统中当前登录的用户           |
| whoami          | 显示您的登录身份                   |
| finger username | 显示有关用户的信息                 |

## 硬件

|                             |                                                           |
| :-------------------------- | :-------------------------------------------------------- |
| dmesg                       | 显示启动消息                                              |
| cat /proc/cpuinfo           | 显示有关CPU的更多信息，例如型号、型号名称、核心、厂商标识 |
| cat /proc/meminfo           | 显示有关硬件内存的更多信息，例如总内存和可用内存          |
| lshw                        | 显示有关系统硬件配置的信息                                |
| lsblk                       | 显示块设备相关信息                                        |
| free -m                     | 显示系统中空闲和使用的内存(-m标志表示内存(MB))            |
| lspci -tv                   | 在树状图中显示PCI设备                                     |
| lsusb -tv                   | 以树状图的形式显示USB设备                                 |
| dmidecode                   | 显示BIOS中的硬件信息                                      |
| hdparm -i /dev/xda          | 显示有关磁盘数据的信息                                    |
| hdparm -tT /dev/xda <:code> | 在设备xda上进行读速度测试                                 |
| badblocks -s /dev/xda       | 测试磁盘上不可读的块                                      |

## 用户

|                  |                                        |
| :--------------- | :------------------------------------- |
| id               | 显示活动用户的详细信息，如uid、gid和组 |
| last             | 显示系统中的最后一次登录               |
| who              | 显示谁已登录到系统                     |
| groupadd “admin” | 添加组“admin”                          |
| adduser “Sam”    | 添加用户 Sam                           |
| userdel “Sam”    | 删除用户 Sam                           |
| usermod          | 用于更改/修改用户信息                  |

## 文件

|                                    |                                                  |
| :--------------------------------- | :----------------------------------------------- |
| ls -al                             | 列出文件-包括常规文件和隐藏文件以及它们的权限    |
| pwd                                | 显示当前目录文件路径                             |
| mkdir ‘directory_name’             | 创建一个新目录                                   |
| rm file_name                       | 删除一个文件                                     |
| rm -f filename                     | 强制删除文件                                     |
| rm -r directory_name               | 递归地删除一个目录                               |
| rm -rf directory_name              | 强制并递归地删除一个目录                         |
| cp file1 file2                     | 将file1的内容复制到file2                         |
| cp -r dir1 dir2                    | 递归地将dir1复制到dir2。如果dir2不存在，则创建它 |
| mv file1 file2                     | 将file1重命名为file2                             |
| ln -s /path/to/file_name link_name | 创建到file_name的软链接                          |
| touch file_name                    | 创建一个新文件                                   |
| cat > file_name                    | 从键盘创建一个文件                               |
| more file_name                     | 输出文件的内容                                   |
| head file_name                     | 显示文件的前10行                                 |
| tail file_name                     | 显示文件的最后10行                               |
| gpg -c file_name                   | 加密一个文件                                     |
| gpg file_name.gpg                  | 解密文件                                         |
| wc                                 | 打印文件中的字节、单词和行数                     |
| xargs                              | 从标准输入执行命令                               |

## 进程

|                         |                                      |
| :---------------------- | :----------------------------------- |
| ps                      | 显示当前活动的进程                   |
| ps aux \| grep ‘telnet’ | 搜索进程’telnet’的id                 |
| pmap                    | 显示进程的内存映射                   |
| top                     | 显示所有正在运行的进程               |
| kill pid                | 使用给定的pid终止进程                |
| killall proc            | 杀死/终止所有名为proc的进程          |
| pkill process-name      | 向具有其名称的进程发送信号           |
| bg                      | 将一个在后台暂停的命令，变成继续执行 |
| fg                      | 将后台中的命令调至前台继续运行       |
| fg n                    | job n to the foreground              |
| lsof                    | 列出进程打开的文件                   |
| renice 19 PID           | 使进程以非常低的优先级运行           |
| pgrep firefox           | 查找Firefox进程ID                    |
| pstree                  | 在树模型中可视化过程                 |

## 文件权限

|                                        |                                                              |
| :------------------------------------- | :----------------------------------------------------------- |
| chmod octal filename                   | 将文件权限更改为八进制                                       |
| chmod 777 /data/test.c                 | 将rwx权限设置为owner、group和everyone(其他可以访问服务器的人) |
| chmod 755 /data/test.c                 | 将rwx设置为所有者，将r_x设置为组和所有人                     |
| chmod 766 /data/test.c                 | 为所有者设置rwx，为组和每个人设置rw                          |
| chown owner user-file                  | 更改文件的所有权                                             |
| chown owner-user:owner-group file_name | 更改文件的所有者和组所有者                                   |
| chown owner-user:owner-group directory | 更改目录的所有者和组所有者                                   |

## 网络

|                                        |                                                 |
| :------------------------------------- | :---------------------------------------------- |
| ip addr show                           | 显示IP地址和所有网络接口                        |
| ip address add 192.168.0.1/24 dev eth0 | 将IP地址192.168.0.1分配给接口eth0               |
| ifconfig                               | 显示所有网络接口的IP地址                        |
| ping host                              | ping命令发送ICMP回送请求以建立到服务器/PC的连接 |
| whois domain                           | 检索有关域名的更多信息                          |
| dig domain                             | 检索关于域的DNS信息                             |
| dig -x host                            | 对域执行反向查找                                |
| host google.com                        | 执行域名的IP查找                                |
| hostname -i                            | 显示本地IP地址                                  |
| wget file_name                         | 从在线资源下载文件                              |
| netstat -pnltu                         | 显示所有活动监听端口                            |

## 压缩/打包

|                                     |                                         |
| :---------------------------------- | :-------------------------------------- |
| tar -cf home.tar home<:code>        | 创建名为“home”的存档文件。tar文件’home' |
| tar -xf files.tar                   | 解压档案文件“files.tar”                 |
| tar -zcvf home.tar.gz source-folder | 从源文件夹创建压缩的tar存档文件         |
| gzip file                           | 压缩扩展名为.gz的文件                   |

## 安装包

|                      |                       |
| :------------------- | :-------------------- |
| rpm -i pkg_name.rpm  | 安装rpm包             |
| rpm -e pkg_name      | 删除rpm包             |
| dnf install pkg_name | 使用dnf工具安装软件包 |

## 安装源(编译)

|              |                                                              |
| :----------- | :----------------------------------------------------------- |
| ./configure  | 检查系统，以获得构建程序所需的软件。它将构建包含有效构建项目所需的指令的Makefile |
| make         |                                                              |
| make install | 编译后，该命令将二进制文件安装在默认/修改的路径中            |

## 搜索

|                           |                                                       |
| :------------------------ | :---------------------------------------------------- |
| grep ‘pattern’ files      | 在文件中搜索给定的模式                                |
| grep -r pattern dir       | Search recursively for a pattern in a given directory |
| locate file               | 查找文件的所有实例                                    |
| find /home/ -name “index” | 在/home文件夹中查找以’index’开头的文件名              |
| find /home -size +10000k  | 在主文件夹中查找大于10000k的文件                      |

## 登陆

|                              |                                 |
| :--------------------------- | :------------------------------ |
| ssh user@host                | 使用指定用户安全连接到主机      |
| ssh -p port_number user@host | 使用指定端口安全地连接到主机    |
| ssh host                     | 通过SSH默认端口22安全连接到系统 |
| telnet host                  | 通过telnet默认端口23连接到主机  |

## 文件传输

|                              |                                               |
| :--------------------------- | :-------------------------------------------- |
| scp file1.txt server2/tmp    | 安全地将file1.txt复制到/tmp目录中的server2    |
| rsync -a /home/apps /backup/ | 将/home/apps目录中的内容与/backup目录进行同步 |

## 磁盘使用情况

|                               |                                              |
| :---------------------------- | :------------------------------------------- |
| df -h                         | 显示安装系统上的空闲空间                     |
| df -i                         | 显示文件系统上的空闲inode                    |
| fdisk -l                      | 显示磁盘分区、大小和类型                     |
| du -sh                        | 以人类可读的格式显示当前目录中的磁盘使用情况 |
| findmnt                       | 显示所有文件系统的目标挂载点                 |
| mount device-path mount-point | 挂载设备                                     |

## 目录遍历

|          |                            |
| :------- | :------------------------- |
| cd ..    | 在目录树结构中向上移动一层 |
| cd       | 将目录更改为$HOME目录      |
| cd /test | 将目录更改为/test目录      |