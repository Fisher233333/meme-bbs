## 项目地址

- 

## 介绍

`bbs-go`是一个使用Go语言搭建的开源社区系统，采用前后端分离技术，Go语言提供api进行数据支撑，用户界面使用Nuxt.js进行渲染，后台界面基于element-ui。如果你正在学习Go语言，或者考虑转Go语言的Phper/Javaer...那么该项目对你有的学习会有很大的帮助，欢迎一起来交流。

![bbs-go功能简介](https://i.loli.net/2021/11/12/OxTBib2pGcV8jzU.png)

## 模块

### server

[![bbs-go-server](https://github.com/mlogclub/bbs-go/actions/workflows/bbs-go-server.yml/badge.svg)](https://github.com/mlogclub/bbs-go/actions/workflows/bbs-go-server.yml)

> 基于`Golang`搭建，提供接口数据支撑。

*技术栈*
- iris ([https://github.com/kataras/iris](https://github.com/kataras/iris)) Go语言 mvc 框架
- gorm ([http://gorm.io/](http://gorm.io/)) 最好用的Go语言数据库orm框架
- resty ([https://github.com/go-resty/resty](https://github.com/go-resty/resty)) Go语言好用的 http-client
- cron ([https://github.com/robfig/cron](https://github.com/robfig/cron)) 定时任务框架
- goquery ([https://github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)) html dom 元素解析

### site

[![bbs-go-site](https://github.com/mlogclub/bbs-go/actions/workflows/bbs-go-site.yml/badge.svg)](https://github.com/mlogclub/bbs-go/actions/workflows/bbs-go-site.yml)

> 前端页面渲染服务，基于`nuxt.js`搭建。

*技术栈*
- vue.js ([https://vuejs.org](https://vuejs.org)) 渐进式 JavaScript 框架
- nuxt.js ([https://nuxtjs.org](https://nuxtjs.org)) 基于Vue的服务端渲染框架，效率高到爆

### admin

> 管理后台系统，基于`vue.js + element-ui`搭建。

*技术栈*
- vue.js ([https://vuejs.org](https://vuejs.org)) 渐进式 JavaScript 框架
- element-ui ([https://element.eleme.cn](https://element.eleme.cn)) 饿了么开源的基于 vue.js 的前端库

## 功能预览

![首页.png](https://s2.loli.net/2022/04/12/DpvPwB9dlQ6Chef.png)
![发帖.png](https://s2.loli.net/2022/04/12/KC8eXfE6sDLq34V.png)
![发动态.png](https://s2.loli.net/2022/04/12/14pMPuGjEU6kiWV.png)
![个人中心.png](https://s2.loli.net/2022/04/12/1PVNjMh9nUAXsl8.png)
![手机版.png](https://s2.loli.net/2022/04/12/mowWb78CGIaH6T2.png)
![后台首页.png](https://s2.loli.net/2022/04/12/ErX2BLTnh7ldz8D.png)
![后台配置.png](https://s2.loli.net/2022/04/12/PwK6aC74XEZlIOL.png)

