#!/usr/bin/env python

import socket

from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer

from helloworld import HelloWorld


class HelloWorldHandler:
    def __init__(self):
        self.log = {}

    def ping(self):
        print "ping()"

    def sayHello(self):
        print "sayHello()"
        return "say hello from " + socket.gethostname()

    def sayMsg(self, msg):
        print "sayMsg(" + msg + ")"
        return "say " + msg + " from " + socket.gethostname()

handler = HelloWorldHandler()
processor = HelloWorld.Processor(handler)
transport = TSocket.TServerSocket()
tfactory = TTransport.TBufferedTransportFactory()
pfactory = TBinaryProtocol.TBinaryProtocolFactory()

server = TServer.TForkingServer(processor, transport, tfactory, pfactory)

print "Starting python server..."
server.serve()
print "done!"
