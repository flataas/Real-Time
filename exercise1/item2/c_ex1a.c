#include <stdio.h>
#include <pthread.h>

#define ITERATIONS 1000000

int i;  // Shared variable with NO synchronization - intentionally broken!

typedef struct {
    int id;
} thread_arg_t;

void* increaser(void *arg) {
    thread_arg_t *args = (thread_arg_t*)arg;
    
    for (int j = ITERATIONS; j > 0; j--) {
        i++;  // RACE CONDITION: Multiple threads accessing shared variable
    }
    
    free(args);
    return NULL;
}

void* decreaser(void *arg) {
    thread_arg_t *args = (thread_arg_t*)arg;
    
    for (int j = ITERATIONS; j > 0; j--) {
        i--;  // RACE CONDITION: Multiple threads accessing shared variable
    }
    
    free(args);
    return NULL;
}

int main(void) {
    pthread_t inc_thread, dec_thread;
    
    i = 0;
    printf("Counter starts at %d\n", i);
    
    thread_arg_t *inc_arg = malloc(sizeof(thread_arg_t));
    inc_arg->id = 1;
    pthread_create(&inc_thread, NULL, increaser, inc_arg);
    
    thread_arg_t *dec_arg = malloc(sizeof(thread_arg_t));
    dec_arg->id = 2;
    pthread_create(&dec_thread, NULL, decreaser, dec_arg);
    
    pthread_join(inc_thread, NULL);
    pthread_join(dec_thread, NULL);
    
    printf("All workers finished.\nCounter ends at %d\n", i);
    
    return 0;
}