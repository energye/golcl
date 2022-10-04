@echo off

title GIT一键提交
color 3

git config --global http.postBuffer 1048576000
echo 设置git post buffer => 1048576000

echo 当前目录是：%cd%

echo 开始添加变更：git add .
git add .

set /p declation=输入提交的commit信息:
git commit -m "%declation%"

echo 将变更情况提交到远程自己分支：git push origin master
git push origin master

echo 执行完毕！
pause