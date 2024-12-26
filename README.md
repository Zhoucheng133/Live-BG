# 直播模板

用于兼容[netPlayer](https://github.com/Zhoucheng133/netPlayer-Next)

**注意，低于v3.1.2版本的netPlayer使用v1.0.1，v3.1.2或更新版本的netPlayer使用v1.0.2或更新的版本**

|netPlayer版本号|适配程序版本|
|-|-|
|<3.0.0|不支持|
|<3.1.2|<1.0.2|
|>=3.1.2|>=1.0.2

## 截图

![截图](other/demo.png)

## 自行打包

1. 安装Node.js，版本>=18
2. 安装yarn:
```bash
npm install yarn -g
```
3. 打包
```bash
yarn build
pyinstaller --onefile --add-data "dist\*;dist/" --add-data "dist\assets\*;dist/assets/" app.py
```

## 更新内容

### 1.0.4 (2024/12/26)
- 添加手动输入服务地址的功能

<details>
<summary>过去的版本</summary>

### 1.0.3 (2024/12/2)
- 添加主动请求数据的功能

### 1.0.2 (2024/7/21)
- 注意 ⚠️ 这个版本开始不再兼容低于v3.1.2版本的netPlayer
- 适配新版本的netPlayer
- 可以自定义ws服务端口

### 1.0.1 (2024/6/24)
- 修复一个布局问题

### 1.0.0 (2024/6/24)
- 第一个版本

</details>