#include <unistd.h>
#include <sys/mman.h>
#include <sys/select.h>

int main() {
    void *addr = mmap(NULL, 300000, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0);
    for (size_t i = 0; i < 300000; i++)
    {
        ((char*)addr)[i] = 'a';

    }

    addr = mmap(NULL, 300000, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0);
    for (size_t i = 0; i < 300000; i++)
    {
        ((char*)addr)[i] = 'a';


    }

    addr = mmap(NULL, 300000, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0);
    for (size_t i = 0; i < 300000; i++)
    {
              ((char*)addr)[i] = 'a';
    }
    
    select(0, NULL, NULL, NULL, NULL);
}