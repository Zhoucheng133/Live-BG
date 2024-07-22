# 直播模板

用于兼容[netPlayer](https://github.com/Zhoucheng133/netPlayer-Next)

**注意，低于v3.1.2版本的netPlayer使用v1.0.1，v3.1.2或更新版本的netPlayer使用v1.0.2或更新的版本**

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