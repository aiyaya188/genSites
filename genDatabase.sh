#!/bin/bash
#批量创建数据库
user=root
password=Aa@20192019
#socket=/usr/local/mysql/mysql.sock
#mycmd="mysql -u$user -p$password -S $socket"
mycmd="mysql -u$user -p$password"
for((i=1;i<10;i++))
do
#创建64个数据库，1-64
#        $mycmd -e "create database sitesDb$i if not exists sitesDb1"  
        $mycmd -e "create database sitesDb$i DEFAULT CHARSET utf8 COLLATE utf8_general_ci"  
#删除64个数据库
#        $mycmd -e "drop database db$i" 
done
