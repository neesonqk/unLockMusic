#include <iostream>
#include <string>
#include <cstring>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
using namespace std;

class Socket {
    public:
        Socket(int _port);
        ~Socket();
        void Listen();
    private:
        int port;
};