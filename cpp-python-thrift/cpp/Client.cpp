#include <iostream>

#include <transport/TSocket.h>
#include <transport/TBufferTransports.h>
#include <protocol/TBinaryProtocol.h>

#include "HelloWorld.h"

using namespace std;

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

int main(int argc, char **argv) {
    boost::shared_ptr<TSocket> socket(new TSocket("localhost", 9090));
    boost::shared_ptr<TTransport> transport(new TBufferedTransport(socket));
    boost::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));
    HelloWorldClient client(protocol);

    try {
        transport->open();

        client.ping();
        cout << "ping()" << endl;

        transport->close();
    } catch (TException& tx) {
        cout << "ERROR: " << tx.what() << endl;
    }

    return 0;
}