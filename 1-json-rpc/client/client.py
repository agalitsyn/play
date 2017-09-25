import requests
import json
import uuid
import random
import sys

import logging
LOG = logging.getLogger(name=__file__)


def rpc_call(url, method, args):
    data = {'id': random.randint(1, sys.maxsize), 'method': method, 'params': [args]}

    r = requests.post(url, json=data)
    return r.json()

args = {'Name': 'django', 'Path': '/tmp'}
resp = rpc_call("http://localhost:5435/rpc", "DatabaseOperation.Backup", args)
if resp['error'] is None:
    print(resp['result'])
