# 直播模板

用于兼容[netPlayer](https://github.com/Zhoucheng133/netPlayer-Next)

如果你想要自行打包这个项目，使用命令
```bash
pyinstaller --onefile --add-data "dist\*;dist/" --add-data "dist\assets\*;dist/assets/" app.py
```