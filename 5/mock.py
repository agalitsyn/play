#!/usr/bin/env python2 -tt
from __future__ import print_function

import argparse
import json
import logging
import time
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from os import environ

logging.basicConfig(level=environ.get('LOG_LEVEL', logging.INFO))
LOG = logging.getLogger('{}:{}'.format(__file__, __name__))
VERSION = environ.get('VERSION', 'unknown')

class HttpProcessor(BaseHTTPRequestHandler):
    def handle_any(self):
        sec = self.headers.get('Mock-Wait-For')
        if sec:
            LOG.info('Wait for {} s.'.format(sec))
            time.sleep(float(sec))

        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.send_header('Connection', 'close')
        self.end_headers()

        resp = {'result': 'ok', 'version': VERSION}
        self.wfile.write(json.dumps(resp))

    def do_GET(self):
        return self.handle_any()

    def do_POST(self):
        return self.handle_any()


def serve(port, protocol='HTTP/1.1'):
    HttpProcessor.protocol_version = protocol

    httpd = HTTPServer(('', port), HttpProcessor)
    hostaddr, port = httpd.socket.getsockname()
    LOG.info('Serving HTTP on {} port {} ...'.format(hostaddr, port))
    httpd.serve_forever()


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Web service mock.')
    parser.add_argument('port', type=int, help='Port for listening')
    args = parser.parse_args()

    serve(args.port)