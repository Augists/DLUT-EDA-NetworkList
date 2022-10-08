# DLUT-EDA Network List

我得承认这个项目是我的愿景，有些不够切合实际

我的想法💡是，大家都有自己的熟人小圈子：

* 每个人每个月对校园网的需求都是波动的
* 上个月用了 20 G，下个月我可能又会需要用 120 G
* 5 个人的小圈子，可能总共只会用 3 个人的量

所以我就想，通过这种方式，让大家均摊校园网的成本，同时也能让你在超量使用的情况下收到来自小圈子里别人的“援手”

它可能有点乌托邦了，但是我相信这能帮助一些人在一定程度上解决校园网的窘境

我把我的想法💡实在在这里，并使用 [Beego](https://github.com/beego/beego) 搭建了前后端将它实现得尽可能的简单易用，希望路过的你能受到启发

## TODO

- [ ] 将 shell 脚本转化为 golang，不再依赖 [mega 的 DLUT-EDA-Login 项目](https://github.com/bboymega/dlut-eda-shell-login)
- [ ] 随机选择账户登陆失败后，继续随机下一个
- [ ] 打包为 Release 发布，多平台

## The MIT License

Copyright © 2022 <Augists>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
