```go
r := mux.NewRouter().StrictSlash(true)

sap := api.SPAHandler {
	StaticPath: "./web/dist",
	IndexPath:"index.html",
}

r.PathPrefix("/").Handler(spa)

http.ListenAndServer("localhost:8080",r)
```

### 设置路由
```go
r := mux.NewRouter().StrictSlash(true)
```
这里使用了gorilla/mux包来创建一个新的路由r，并且设置了`StrictSlash(true)`。这意味有路由匹配，会严格区分末尾的`/`,确保路径一致

### 配置SPA处理器
```go
sap := api.SPAHandler {
	StaticPath: "./web/dist",
	IndexPath:"index.html",
}

```
SPAHandler自定义了HTTP处理器，用来处理单页应用（SPA）的路由，它将
`StaticPath`设置为`./web/dist`,即打包好的前端文件路径，`IndexPath`设置为
index.html，即默认返回路径

### 默认路由
```go
r.PathPrefix("/").Handler(spa)
```
这行代码未经匹配到的APIl路由请求都会有SPAHandler处理，这对SPA非常重要，因为通常所有的路径都会返回 index.html 
然后由前端的路由逻辑决定显示什么内容

### SPAHandler原理
```go

func (h SPAHandler) ServerHTTP(w http.Response, r *http.Request){
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
	    http.Error(w,err.Error(),http.StatusBadRequest)
		return
    }
	
	path = filepath.Join(h.StaticPath,path)
	
	_,err = os.Stat(path)
	if os.IsNotExist(err) {
	    // file does not exist,serve index.html
		http.ServeFile(w,r,filepath.Join(h.StaticPath,h.IndexPath))
		return
    } else if err != nil{
	    http.Error(w,err.Error(),http.StatusInternalServerError)
		return
    }
	
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w,r)
}
```

`SPAHandler`是一个自定义的HTTP处理器，用来处理所有的静态文件的请求，当请求来临时，它会检查请求的路径是否指向了实际存在的文件

1. 路径解析和文件检查：它首先将请求转换为绝对路径，并尝试找响应的文件，如果文件不存在，说明这是一个SPA的路由请求，而不是静态文件，
因此返回`index.html`，让前端框架如Vue或React处理路由
2. 返回静态文件或`index.html`, 如果请求的文件存在，则直接返回该文件，否则返回index.html
3. 处理静态文件：`http.FileServer(http.Dir(h.StaticPath)).ServerHTTP(w,r)`负责处理静态文件的请求，将文件内容返回给客户端

# 为什么文件不存在说明是SPA路由，返回index.html为什么前端能处理

要理解为什么文件不存在就说明这是一个SPA的路由请求，并且能返回`index.html`后前端能处理，以下是一些具体的例子

### 假设的场景
你有一个基于React或Vue构建的单页应用（SPA）。这个应用的前端代码已经打包放在`web/dist`目录下，其中包含以下文件

* `index.html`: 单体应用的入口文件
* `assets/index.js`: 打包好的JavaScript代码，包含应用逻辑和路由配置
* `assets/style.css`: 样式文件

你希望用户访问不同url路径查看不同的页面和组件，比如

* '/': 显示首页组件
* '/about' 显示关于我们页面组件
* '/contact': 显示联系页面组件

### 传统Web应用与SPA的区别
在传统Web应用中，每一个URL路径（如`/about`或`contact`）都会对应一个物理文件，比如`about.html`或`contact.html`。
当用户请求这些路径的时候，服务器会返回对应的HTML文件。

但在SPA，所有页面实际上是由JavaScript在浏览器中动态生成的。整个应用只有一个`index.html`文件，无论用户的请求路径是什么，
服务器都只返回这个`index.html`。然后，由前端框架（如React或Vue）根据用户请求路径，动态的决定展示那个组件或页面。

### 文件不存在的例子
假设你在浏览器输入`http://localhost:8080/about` ,在传统的Web应用中，服务器会寻找一个名为about.html文件，如果这个文件
存在，服务器就返回它，在SPA中，这个`about.html` 文件并不存在，在服务器找不到这个文件，说明这是SPA的路由请求，而不是静态文件请求

### 为什么返回`index.html`
当服务器发现请求文件不存在时，它返回`index.html`文件，因为这是SPA的入口文件，返回这个文件以后，前端会通过JavaScript代码处理这个路径。

例如，在`index.js`中，React Router(或者Vue Router)会根据当前的URL路径决定展示那个组件。
```javascript
import React from 'react'
import {BrowserRouter as Router,Route,Switch} from 'react-router-dom';
import Home from './Home'
import About from './About'
import Contact from './Contact';

function App() {
    return (
        <Router>
            <Switch>
                <Route exact path="/" component={Home} />
                <Route path="/about" component={About} />
                <Route path="/contact" component={Contact}/>
            </Switch>
        </Router>
    );
}

export default App;
```

# 为什么选择SPA而不是传统Web应用

选择SPA有以下几个关键因素

### 1.用户体验
无缝导航：SPA允许用户在页面之间导航时不需要刷新整个页面，这大大的提升了用户体验。
页面的切换几乎是及时的，因为所有必须要的资源（HTML、CSS、JavaScript）只会加载一次。
后续切换页面仅仅涉及到内容的动态更新

### 2.更高的开发效率
* 组件化开发，使用React或Vue等框架进行SPA开发时，可以将页面划分为多个可复用的组件。这样不仅使
代码更容易管理和维护，也方便开发团队并发开发不同功能
* 前后端分离：SPA通常与后端API进行通信，而不是传统的服务器端渲染，这些前后端分离架构可以使的
前端和后端独立开发、测试和部署，提高开发效率

### 3.现代应用
* 更好支持移动设备，SPA天生支持移动设备，因为他能减少页面的加载时间，提供更流畅的用户体验



