// go get 命令————一键获取代码，编译并安装
// c.biancheng.net/view/123.html
package main

//go get 命令可以借助代码管理工具通过远程拉取或更新代码包及其依赖包
//并自动完成编译和安装
//整个过程就像安装一个app一样简单
//
//这个命令可以动态获取远程代码包，目前支持的有bitbucket，githuub，Google code 和launchpad
//在使用 go get 命令前，需要安装与远程包匹配的代码管理工具，如Git，SVN，HG等
//参数中需要提供一个包名
//
//这个命令在内部实际上分成了两部操作
//第一步是下载源码包
//第二部是执行go install
//下载源码包的go工具回自动根据不同的域名调用不同的源码工具，对应关系如下：
//BitBucket （Mercurial Git）
//GitHub（Git）
//Google Code Project Hosting （Git， Mercurial， Subversion）
//Launchpad （Bazaar）
//所以为了go get命令能正常工作，
//你必须确保安装了合适的源码管理工具
//并同时把这些命令加入你的PATH中，其实go get 支持自定义域名的功能
//
//参数介绍 ：
//-d 只下载不安装
//-f 只有在你包含了 -u参数的时候才有效， 不让-u去验证import中的每一个都已经获取了
//这对于本地fork的包特别有用
//-fix在获取源码之后先运行fix， 然后再去做其他的事情
//-t 同时也下载需要为运行测试所需要的包
//-u 强制使用网络去更新包和它的依赖包，（不会更新已经存在的包？？）
//-v 显示执行的命令
//-insecure 允许使用不安全的HTTP方式进行下载操作

//
//远程包的路径格式
//go语言的代码被托管于github，该网站基于git代码管理工具
//类似的托管网站还有：code.google.com, bitbucket.org
//
//这些网站的项目包路径都有一个共同的标准，
// github.com/golang/go
//网站域名   作者或机构 项目名
//网站域名：代码托管的网站
//作者或机构：表明这个项目的归属，一般为网站的用户名，
//项目名： 每个网站下的作者或机构可能回同时拥有很多的项目

// go get + 远程包
//默认情况下，go get可以直接使用
//例如：想获取go的源码并编译，使用下面的命令行即可：
//go get github.com/davyxu/cellnet
//获取前，请确保GOPATH已经设置
//go 1.8 后，GOPATH默认在用户目录的go文件夹下
//cellnet只是一个网络库，并没有可执行文件，因此在go get 操作成功后GOPATH下的bin目录下不会有任何编译好的二进制文件
//需要测试获取并编译的二进制的，可以尝试下面的这个命令
//当获取完成皇后，就会自动在GOPATH的bin目录下生成编译好的二进制文件
// go get github.com/davyxu/tabtoy
//
//go get使用时的附加参数
///

func main() {

}
