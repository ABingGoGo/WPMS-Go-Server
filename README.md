# WPMS-Go-Server

## 项目背景说明

Work Plan Management System, used to manage daily work plans.  

As a novice programmer in the industry, I plan to use go+c # to implement the coding work of this system. Before this, I had no experience with coding for Go or C #, but I love coding and hope to improve my coding skills through this system's coding work.

﻿工作计划管理系统,用于管理日常工作计划。  

作为一个初入行业的程序员,我计划使用go+c#来实现这一系统的编码工作。在此之前,我并没有接触过go或者c#的编码工作,不过我热爱代码,希望通过这个系统的编码工作来使我的编码能力得到提升。


## 项目技术栈
- go1.22.0 windows/amd64  
- xorm  
- gin

## 项目目录结构
由于历史原因，我并没有接受过系统的编码学习，现在我自身的所有编码能力都是通过网上自学获得的。  
所以我结合现有的知识以及想法，暂时将项目的目录结构这样划分：  
- 根目录  
    - 服务模块1  
        - cmd
           - xxxxApplication.go (此服务启动入口文件)
        - config 
           - xxxxApplication.go (配置文件处理代码)
        - internal
           - entity (数据库结构体代码包)
           - mapper (数据库自定义查询方法包)
              - DataBaseInit.go (数据库初始化代码，包含创建数据库连接)
           - service (业务层代码包)
           - router (路由代码包)
           - controller (控制器代码包)
           - model (对外暴露结构体包)
        - application.yaml (此服务配置文件)
    - 服务模块2  
    - 服务模块3