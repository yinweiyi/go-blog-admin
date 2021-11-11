用go重写了一遍博客的前端页面，重新学习了一遍golang，演示地址[点击此处](http://go.blog.ewayee.com)。记录一下在开发的过程中遇到的问题

###框架的选择
网上看了一遍市面上常用的go web框架，主要包括：Beego、Echo、Iris、Gin等等，都非常出色，因为只是写个比较简单的业务，最终选择了gin作为本博客的开发框架

###多模板的支持
因为博客内容有很多个版块，像 随机标签云、最近热门、友情链接、评论表单以及评论列表等，这些模块的模板都是可以复用的，而Gin 默认允许只使用一个 html 模板,我是导入了gin推荐的
[多模板渲染](https://github.com/gin-contrib/multitemplate),以使用``` laravel blade```模板引擎类似go 1.6的 ```block template```功能

#### 使用多模板渲染时遇到的问题：
模板的导入顺序会影响到渲染的内容：我有一个``` master.html``` 基础模板，其它所有的模板都是继承自此模板，我是打算把此模板和其它的公共模板放到同一个名为``` layouts```的录下，
而我写了一个公共的方法遍历文件夹并导入模板的。 在开发的过程我当我新建``` 评论``` 表单模板的时候，因为是 ```comment.html```,而首字母c 在 m的前面，导致先渲染评论的模板，最终导致渲染出来只有评论表单的内容，找了好久问题才发现是模板导入顺序的问题。

模板中使用自定义方法：我们需要在模板中使用一些自定义的方法，比如首页列表文章内容要去掉html标签、日期格式化、随机标签云要随机文字大小以及颜色等，[多模板渲染](https://github.com/gin-contrib/multitemplate)工具提供了一个方法 ``` AddFromFilesFuncs```,可以传入一些自定义的函数

代码如下：
``` go
func main() {
    engine := gin.Default()

    engine.Run(":3000") // 监听并在 0.0.0.0:8080 上启动服务

    //加载公共模板
    engine.HTMLRender = loadTemplates("./resources/views")
}

//加载公共模板
func loadTemplates(templatesDir string) multitemplate.Renderer {
    r := multitemplate.NewRenderer()
    //Master template must render at first
    masterFile := []string{templatesDir + "/master.html"}
    
    layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
    if err != nil {
        panic(err.Error())
    }
    
    includes, err := filepath.Glob(templatesDir + "/*[^layouts]/*.html")
    if err != nil {
        panic(err.Error())
    }
    // Generate our templates map from our layouts/ and includes/ directories
    for _, include := range includes {
        layoutCopy := make([]string, len(layouts))
        copy(layoutCopy, layouts)
        files := append(masterFile, append(layoutCopy, include)...)
        r.AddFromFilesFuncs(fileName(include), template.FuncMap{
            "randomColor":      helpers.RandomColor,
            "randomInt":        helpers.RandomInt,
            "formatAsDate":     helpers.FormatAsDate,
            "formatAsDateTime": helpers.FormatAsDateTime,
            "stripTags":        strip.StripTags,
            "unescaped":        helpers.Unescaped,
            }, files...)
    }
    return r
}

//获取文件名
func fileName(filePath string) string {
    pathArr := strings.Split(filePath, string(filepath.Separator))
    pathArr = pathArr[len(pathArr)-2:]
    return strings.Join(pathArr, "/")
}

```
```master.html```放到```resources/views```下,其它公共的模板放到```resources/views/layouts```下,控制器里就可以这样调用了

``` go
 ctx.HTML(200, "article/show.html", gin.H{
    
 })
```
```article```是```resources/views``` 下的文件夹， ```show.html```是模板文件

###图片验证码

网上找了一下golang图片验证码的第三方库，最终选择了[github.com/dchest/captcha](https://github.com/dchest/captcha) 大概步骤就是生成图片的时候返回一个id，服务器写一个根据这个id来返回图片内容的服务，验证的时候也是根据这个id来验证。但就是渲染出来的图片看不太清楚，也没找到哪里可以配置图片清晰度，还没找到解决的办法。
