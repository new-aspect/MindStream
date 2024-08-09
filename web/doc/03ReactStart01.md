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

### 响应事件
您可以通过在组件内声明事件来处理函数响应事件

```javascript
function MyButton() {
    function handleClick() {
        alert('You clicked me!')
    }

    return (
        <button onClick={handleClick}>
            Click me
        </button>
    )
}
```

请注意`onClick={handleClick}` 末尾没有括号! 不要调用事件处理函数，您只需要将其传递下去。当用户点击按钮的时候，React将调用你的事件处理函数

### 更新屏幕

通常，你会希望组件“记住”一些信息并显示它，例如，也许您想计算按钮被点击的次数。为此，请向您的组件添加状态

首先，从React 导入 useState

```javascript
import {useState} from 'react';
```

现在，你可以在组件内声明一个状态变量

```javascript
function MyButton() {
    const [count, setCount] = useState(0);
}
```

你将从useState中获得两件事，当前状态（count）和允许您更新它的函数（setCount）。您可以给他们起任何名字，但是约定编写[something, setSomething]。

第一次显示按钮的时候，count将为0，因为您将0传递给了useState()。当您想要改变状态的时候，请调用setCount()并将新值传递给它，单击次按钮将增加计数器

```javascript
function MyButton() {
    const [count, setCount] = useState(0);

    function handleClick() {
        setCount(count + 1);
    }

    return (
        <button onClick={handleClick}>
            Clicked {count} times
        </button>
    );
}
```

React将再次调用你的组件函数，这次，count将为1，那么它讲为2，等等。

如果多次渲染一个组件，每个组件都会获得自己单独的状态，分别点击每个按钮。

请记住每个按钮记住自己的count状态并不会影响其他按钮。
