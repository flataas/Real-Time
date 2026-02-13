// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;

// Note the return type: void*
void* incrementingThreadFunction(void*arg){
    for (int k = 0; k < 1000000; k++){
        pthread_mutex_lock(&mtx);
        i++;
        pthread_mutex_unlock(&mtx);
    }
    return NULL;
}

void* decrementingThreadFunction(void*arg){
    for (int k = 0; k < 1000000; k++){
        pthread_mutex_lock(&mtx);
        i--;
        pthread_mutex_unlock(&mtx);
    }
    return NULL;
}


int main(){
    pthread_t incThread, decThread;

    pthread_create(&incThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incThread, NULL);
    pthread_join(decThread, NULL);

    printf("The magic number is: %d\n", i);
    return 0;
}
