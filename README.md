# v0.1 版 
start_gui.bat:负责启动后端服务器和前端gui程序<br>
server:后端API服务程序目录，使用python编写<br>
hackclown：前端界面程序目录，使用react + typescript + electron编写<br>
# 增加功能：
1.android常见漏洞介绍及利用方法<br>
2.增加了服务器后端，功能包括：<br>
searchsploit查询<br>
端口扫描<br>
http服务探测<br>
目录扫描<br>
nuclei扫描<br>
其他插件式扫描<br>
如图所示：
![image](https://github.com/followboy1999/hackclown/blob/main/img/android.png)
![image](https://github.com/followboy1999/hackclown/blob/main/img/searchsploit.png)
![image](https://github.com/followboy1999/hackclown/blob/main/img/nuclei.png)
![image](https://github.com/followboy1999/hackclown/blob/main/img/portscan.png)
![image]()
# 安装
1.安装mongodb</br>
下载地址：https://www.mongodb.com/try/download/community<br>
双击运行mongodb-windows-x86_64-7.0.1-signed.msi<br>
2.配置mongodb数据库<br>
安装完成后，使用mongodb compass连接本地localhost:27017，创建数据库clown，在数据库内创建ffuf_tasks c2_tasks gofree_tasks nuclei_tasks portscan_tasks webprobe_tasks，如下图<br>
![image](https://github.com/followboy1999/hackclown/blob/main/img/clown-database.png)
3.运行start_gui.bat
# 引用
https://github.com/projectdiscovery/nuclei<br>
https://github.com/shadow1ng/fscan<br>
https://github.com/niudaii/zpscan<br>
https://github.com/ffuf/ffuf<br>
