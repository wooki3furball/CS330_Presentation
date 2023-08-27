# Author: Michael Bopp
# Filename: server.py
# Description: Python program to spin up a local test server & launch browser.
# Date Created: 8/22/23
# Last Edited: 8/22/23
# Dependency: Python Interpreter

import http.server
import socketserver
import webbrowser
import threading

PORT = 8080

Handler = http.server.SimpleHTTPRequestHandler
Handler.extensions_map = {
    '.html': 'text/html',
    '.wasm': 'application/wasm',
    '': 'application/octet-stream',
}

def open_browser():
    webbrowser.open('http://localhost:8080')

# Start Server
with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print("serving at port", PORT)
    threading.Timer(0.5, open_browser).start() # Spin up browser
    httpd.serve_forever()