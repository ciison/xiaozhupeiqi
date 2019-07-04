# 安装 protobuf

~~~
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
~~~

# 报错
遇到错误, 如果遇到墙我也无能为力了.....
~~~ vim /etc/ld.so.conf
protoc: error while loading shared libraries: libprotoc.so.20: cannot open shared object file: No such file or directory
~~~

解决方法： [参考](https://blog.csdn.net/stevenluopan/article/details/44746193)

~~~
修改配置文件
vim /etc/ld.so.conf

添加
 /usr/local/lib
 
然后执行: 刷新配置文件, 是配置生效
sudo ldconfig 
~~~