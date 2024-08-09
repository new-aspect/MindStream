### 创建和嵌套组件

React的应用程序都是由组件组成的，组件是UI(用户界面)的一部分，具有自己的逻辑和外观。组件可以小到一个按钮

也可以大到整个界面。

React组件返回标记的JavaScript函数

```javascript
function MyButton() {
    return (
        <button>I'm a button</button>
    );
}
```

现在你已经声明了MyButton，你可以嵌套到另一个组件的:

```javascript
export default function MyApp() {
    return (
        <div>
            <h1>Welcome to my app</h1>
            <MyButton />
        </div>
    )
}
```

请注意，</Mybutton> 以大写字母开头，你就知道它是一个React组件的方式。React组件必须始终以大写字母开头，而HTML标签必须小写

`export default`关键子制定了文件中的主要组件。如果你不熟悉某些JavaScript语法，[MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/export) 和[javascript.info](https://javascript.info/import-export)由很好的参考资料


### 使用JSX标记
你看到上面的标记语法成为JSX。它是可选的，但大多数React项目为了方便使用JSX。我们推荐用于本地开发的所有工具都支持开箱即用的JSX.

JSX比HTML更严格。您必须关闭诸如`<br >`之类的标签。您的组件不能返回多个JSX标记。您必须将他们包装到共享父级中，例如<div></div>
或空的<></>包装器

```javascript
function AboutPage() {
    return (
        <>
            <h1>About</h1>
            <p>Hello there. <br/> How do you do?</p>
        </>
    )
}
```
如果您有大量的HTML需要移植到JSX，可以使用[在线转换器](https://transform.tools/html-to-jsx)

### 添加样式
在React中，您可以通过`className` 指定CSS类，它的工作方式与HMTL的`class`属性相同

```html
<img className="avatar"/>
```

然后在单独的CSS里面为其编写CSS规则

```css
.avatar {
    border-radius: 50%;
}
```

React没有规定如恶化添加CSS文件，在最简单的情况下，你将向HTML添加`link`标记，如果您使用构建工具或框架，
请查阅其文档以了解CSS如何添加到你的项目中。

### 显示数据
JSX允许你将标记放入到JavaScript中，大括号可以让你转回JavaScript，这样你就可以从代码中嵌入一些变量并将其展示给用户。
例如，这里将显示user.name

```javascript
return (
    <h1>
        {user.name}
    </h1>
)
```

你还可以从JSX属性"转义"到JavaScript。但必须使用花括号而不是引号，例如，`className="avatar"`将"avatar"转换为CSS类传递属性，

但src={user.imageUrl}读取JavaScript的user.imageUrl变量值，然后将该值作为src属性传递

```javascript
return (
    <img
        className="avatar"
        src={user.imageUrl}
        style={{
            width: user.imageSize,
            height: user.imageSize
        }}
    />
)
```

上面的示例中，`style={{}}`并不是特殊语法，而是`style={}`JSX大括号内常规{}对象，这样您的样式依赖于JavaScript变量时，您可以使用style属性

### 渲染列表

您将依赖JavaScript功能，（例如for循环和数组map()函数来呈现组件列表）

例如，假设您有一系列产品:

```javascript
const products = [
    {title:'Cabbage',id:1},
    {title:'Garlic',id:2},
    {title:'Apple',id:3},
];
```

在组件内部，使用map()将函数转换为<li>项目数组

```javascript
const listItems = products.map(product =>
    <li key={product.id}>
        {product.title}
    </li>    
 )
return (
    <ul>{listItems}</ul>
)
```

请注意<li>如何具有key属性，对于列表中的每个项目，您应该传递一个字符串或数字，以在其同级项目中唯一表示该项目，
通常，秘钥应该来自您的数据，例如数据库的ID。React使用您的键来了解您的稍后插入、删除或重新排序项目时发生的情况