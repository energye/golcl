@echo off

title GITһ���ύ
color 3

git config --global http.postBuffer 1048576000
echo ����git post buffer => 1048576000

echo ��ǰĿ¼�ǣ�%cd%

echo ��ʼ��ӱ����git add .
git add .

set /p declation=�����ύ��commit��Ϣ:
git commit -m "%declation%"

echo ���������ύ��Զ���Լ���֧��git push origin master
git push origin master

echo ִ����ϣ�
pause