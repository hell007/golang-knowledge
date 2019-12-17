#!/bin/bash

# 以后台方式启动程序，并且将日志记录到 app.log
nohup ./$1 >> app.log &