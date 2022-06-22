# zero-mall

#### 介绍
用go-zero实现一个开箱即用的电商微服务,另有[go-kratos的实现版本](https://gitee.com/zhaobu/mfkratos-mall.git)
项目正在努力完善中...
项目的发起源自自己业余时间的兴趣爱好,github上搜索电商,mall等关键字,比如mall, litemall 等这些排名前十的上万starts的优秀开源项目都是java技术栈的.目前golang社区最流行的开源微服务框架要数go-zero, kratos, kitex, 其中以go-zero最为活跃和落地项目最多. 所以mfzero-mall的目标就是学习看起mall, 通过这种接近企业级的开源项目提升整个gopher社区的技术氛围,共同进步, 为新的gopher提供优秀的学习参考资料. 

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)

#### Commit 规范

###### 格式如下

* 例：`fead(type)`:本次提交概述
* `type`: 本次 commit 的类型，诸如 bugfix docs style 等，参考如下:   

    * `fead`：添加新功能
    * `fix`：修补缺陷
    * `docs`：修改文档
    * `style`：修改格式
    * `refactor`：重构
    * `perf`：优化
    * `test`：增加测试
    * `chore`：构建过程或辅助工具的变动
    * `revert`：回滚到上一个版本
    * `merge`：合并

* `scope`: 本次 `commit` 波及的范围
* `subject`: 简明扼要的阐述下本次 `commit` 的主旨，在原文中特意强调了几点：

    1. 使用祈使句
    2. 首字母不要大写
    3. 结尾无需添加标点

---
##### 分支记录
* `dev`：开发版 `prod`：发布版 `stage`：稳定版
* 命令记录：
```
git checkout -b dev   // 新建分支
git push origin dev   // 推送新分支到线上
git branch -d dev     // 删除本地分支
git push origin :dev  // 删除远程分支,需要设置权限

git checkout pro // 切换到发布分支
git  merge dev   // 再合并开发分支
```
