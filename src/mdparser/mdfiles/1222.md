This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br>
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (Webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

## 搭建过程
## 初始化
安装react-router react-router-dom
在app.js文件中引入公共模块以及PrimaryLayout
在app.js中引入react-router-dom 在模块最外层添加HashRouter以保证在所有子模块中可以使用react-router-dom所有标签 
## 定义router模块
引入react-router-dom中（Route，Switch）以及router定义的所有模块
其中Switch标签保证路由规则匹配不到时匹配到My404模块
## 定义公共变量
config.js定义公共变量并且在index.js中引入
例：global.constants = {
  Fragment: React.Fragment
};
在所有子模块中引用 
例：const Fragment = global.constants.Fragment

## 状态以及父子组件调用
例：home.js
定义 isShowList3初始状态
 constructor(props) {
    super(props);
    this.state = {
      isShowList3: false
    };
  }
  点击this.showConponent改变当前状态
 父子组件方法调用
 父组件定义hideConponent方法
传值
  <List3 hideConponent={this.hideConponent} />

  在子组件中通过this.props.hideConponent调用父组件中的方法
  

