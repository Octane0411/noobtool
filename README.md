# noobtool

学完了go的基础语法后想做一些简单的项目，在《gopl》中做过du的例子，索性封装一下，再加上一些其他我自己平时常用的工具和爬虫，做一个命令行工具

做了大概四五个功能

- tree：同linux环境中常用的tree
- email：快速发邮件
- ipgw：抄的代码，基于neugo
- v：获取valorant当日的皮肤商店，包括武器名字和图片
- ...

## 实现思路

本着了解一下go社区的目的大概逛了逛，命令行工具的框架大概有如下两个

- cobra
- cli

作为小白看不出什么区别，直接选star数更多的那个了

在功能实现中有很多个性化的配置如邮箱地址和校园网账号等等，引入了viper，这样任何人使用只要配置一下yaml就好了

获取valorant皮肤商店的功能基于https://checkvalorant.gamzo.in/的api，但目前实现认证仍然是个问题

## 我的收获

1. 知道了如何用go开发一个命令行工具
2. 在项目中学习了良好的go独有的项目结构，命名规范
3. tree命令中大量使用了goroutine，复习了之前学习的部分



