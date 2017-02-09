GEN_SRC := helloworld_constants.cpp  HelloWorld.cpp  helloworld_types.cpp
GEN_OBJ := $(patsubst %.cpp,%.o, $(GEN_SRC))

THRIFT_DIR := /usr/local/include/thrift
BOOST_DIR := /usr/include/

INC := -I$(THRIFT_DIR) -I$(BOOST_DIR)

.PHONY: all clean
all: server client

%.o: %.cpp
	$(CXX) -Wall -DHAVE_INTTYPES_H -DHAVE_NETINET_IN_H $(INC) -std=c++11 -c $< -o $@

server: Server.o $(GEN_OBJ)
	$(CXX) $^ -o $@ -L/usr/local/lib -lthrift

client: Client.o $(GEN_OBJ)
	$(CXX) $^ -o $@ -L/usr/local/lib -lthrift

clean:
	$(RM) *.o server client

