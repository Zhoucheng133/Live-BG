from flask import Flask, send_from_directory
from gevent import pywsgi

app = Flask(__name__, static_folder='dist', static_url_path='')

port=9098

@app.route('/')
def serve_index():
    return send_from_directory(app.static_folder, 'index.html')

@app.route('/api/port')
def serve_port():
    global port
    return port

if __name__ == '__main__':
    port=input("输入端口号 (默认为9098): ")
    if port=="":
        port="9098"
        print("使用默认端口: 9098")
    print("服务运行在: http://127.0.0.1:5000")
    server=pywsgi.WSGIServer(('0.0.0.0', 5000), app, log=None)
    server.serve_forever()
