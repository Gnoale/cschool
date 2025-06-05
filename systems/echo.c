#include <stdio.h>

int main() {
    client_fd := accept(listen_fd, &client_addr, &client_addrlen);
    for (;;) {
        numRead = read(client_fd, buf, BUF_SIZE);
        if (numRead <= 0) {  // exit loop on EOF or error
            break;
        }
        if (write(client_fd, buf, numRead) != numRead) {
            // handle write error
        }
    }
    close(client_fd);
}
