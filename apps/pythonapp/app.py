from flask import Flask
import os
import socket

app = Flask(__name__)

@app.route('/')
def hello():
    hostname = socket.gethostname()
    return f"""
    <h1>Hello from Python Flask App!</h1>
    <p><strong>Hostname:</strong> {hostname}</p>
    <p><strong>Environment:</strong> Running in Kubernetes</p>
    <p>This app is deployed using Docker and Kubernetes (kind)</p>
    """

@app.route('/health')
def health():
    return {'status': 'healthy'}, 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
