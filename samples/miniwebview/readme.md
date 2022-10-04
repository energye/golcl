```bigquery
linux webkit 库支持
    打开 
        sudo vim /etc/apt/sources.list
    最后一行添加 
        deb http://cz.archive.ubuntu.com/ubuntu bionic main universe
    执行 sudo apt-get update
    可选-（1）安装 1.0 版本
        sudo apt-get install libwebkitgtk-1.0-0
    可选-（2）安装 3.0 版本 - 可能有bug，不能正常使用
        sudo apt-get install libwebkitgtk-3.0-0
```