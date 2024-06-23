from flask import Flask, send_from_directory
from gevent import pywsgi

app = Flask(__name__, static_folder='dist', static_url_path='')

@app.route('/')
def serve_index():
    return send_from_directory(app.static_folder, 'index.html')

if __name__ == '__main__':
    # app.run(debug=False, port=5000)
    server=pywsgi.WSGIServer(('0.0.0.0', 5000), app)
    server.serve_forever()
