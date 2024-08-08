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