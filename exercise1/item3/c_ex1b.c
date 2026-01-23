#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define ITERATIONS 1000000

// Shared state
typedef struct {
    int counter;
    int active_workers;
    pthread_mutex_t mutex;  // Mutex for mutual exclusion (critical section protection).
                            // Semaphores are for signaling/counting, not ownership-based locking.
    pthread_cond_t workers_done;
} shared_state_t;

shared_state_t state;

void* incrementer(void *arg) {
    for (int i = 0; i < ITERATIONS+1; i++) {
        pthread_mutex_lock(&state.mutex);
        state.counter++;
        pthread_mutex_unlock(&state.mutex);
    }
    
    // Signal we're done
    pthread_mutex_lock(&state.mutex);
    state.active_workers--;
    if (state.active_workers == 0) {
        pthread_cond_signal(&state.workers_done);
    }
    pthread_mutex_unlock(&state.mutex);
    
    return NULL;
}

void* decrementer(void *arg) {
    for (int i = 0; i < ITERATIONS; i++) {
        pthread_mutex_lock(&state.mutex);
        state.counter--;
        pthread_mutex_unlock(&state.mutex);
    }
    
    // Signal we're done
    pthread_mutex_lock(&state.mutex);
    state.active_workers--;
    if (state.active_workers == 0) {
        pthread_cond_signal(&state.workers_done);
    }
    pthread_mutex_unlock(&state.mutex);
    
    return NULL;
}

int main(void) {
    pthread_t inc_thread, dec_thread;
    
    // Initialize shared state
    state.counter = 0;
    state.active_workers = 2;
    pthread_mutex_init(&state.mutex, NULL);
    pthread_cond_init(&state.workers_done, NULL);
    
    // Start threads
    pthread_create(&inc_thread, NULL, incrementer, NULL);
    pthread_create(&dec_thread, NULL, decrementer, NULL);
    
    // Wait for workers to finish
    pthread_mutex_lock(&state.mutex);
    while (state.active_workers > 0) {
        pthread_cond_wait(&state.workers_done, &state.mutex);
    }
    printf("Routines done! counter at %d\n", state.counter);
    pthread_mutex_unlock(&state.mutex);
    
    // Cleanup
    pthread_join(inc_thread, NULL);
    pthread_join(dec_thread, NULL);
    pthread_mutex_destroy(&state.mutex);
    pthread_cond_destroy(&state.workers_done);
    
    return 0;
}