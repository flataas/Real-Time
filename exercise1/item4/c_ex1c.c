#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>
#include <unistd.h>

#define BUFFER_SIZE 5
#define NUM_ITEMS 20

// Bounded buffer structure
typedef struct {
    int *data;
    int capacity;
    int in;   // Write index
    int out;  // Read index
    
    pthread_mutex_t mutex;  // Protects buffer operations
    sem_t empty_slots;      // Counts empty slots (producers wait on this)
    sem_t filled_slots;     // Counts filled slots (consumers wait on this)
} bounded_buffer_t;

bounded_buffer_t buffer;

void buffer_init(bounded_buffer_t *bb, int capacity) {
    bb->data = malloc(sizeof(int) * capacity);
    bb->capacity = capacity;
    bb->in = 0;
    bb->out = 0;
    
    pthread_mutex_init(&bb->mutex, NULL);
    sem_init(&bb->empty_slots, 0, capacity);  // Initially all empty
    sem_init(&bb->filled_slots, 0, 0);        // Initially none filled
}

void buffer_push(bounded_buffer_t *bb, int item) {
    sem_wait(&bb->empty_slots);  // Wait if buffer is full
    
    pthread_mutex_lock(&bb->mutex);
    bb->data[bb->in] = item;
    bb->in = (bb->in + 1) % bb->capacity;
    printf("Produced: %d (buffer indices: in=%d, out=%d)\n", item, bb->in, bb->out);
    pthread_mutex_unlock(&bb->mutex);
    
    sem_post(&bb->filled_slots);  // Signal one more item available
}

int buffer_pop(bounded_buffer_t *bb) {
    sem_wait(&bb->filled_slots);  // Wait if buffer is empty
    
    pthread_mutex_lock(&bb->mutex);
    int item = bb->data[bb->out];
    bb->out = (bb->out + 1) % bb->capacity;
    printf("Consumed: %d (buffer indices: in=%d, out=%d)\n", item, bb->in, bb->out);
    pthread_mutex_unlock(&bb->mutex);
    
    sem_post(&bb->empty_slots);  // Signal one more slot available
    
    return item;
}

void buffer_destroy(bounded_buffer_t *bb) {
    free(bb->data);
    pthread_mutex_destroy(&bb->mutex);
    sem_destroy(&bb->empty_slots);
    sem_destroy(&bb->filled_slots);
}

void* producer(void *arg) {
    int id = *(int*)arg;
    
    for (int i = 0; i < NUM_ITEMS; i++) {
        int item = id * 100 + i;  // Create unique items
        buffer_push(&buffer, item);
        usleep(100000);  // Simulate work (100ms)
    }
    
    return NULL;
}

void* consumer(void *arg) {
    int id = *(int*)arg;
    
    for (int i = 0; i < NUM_ITEMS; i++) {
        int item = buffer_pop(&buffer);
        usleep(150000);  // Simulate slower consumption (150ms)
    }
    
    return NULL;
}

int main(void) {
    pthread_t prod_thread, cons_thread;
    int prod_id = 1, cons_id = 1;
    
    buffer_init(&buffer, BUFFER_SIZE);
    
    pthread_create(&prod_thread, NULL, producer, &prod_id);
    pthread_create(&cons_thread, NULL, consumer, &cons_id);
    
    pthread_join(prod_thread, NULL);
    pthread_join(cons_thread, NULL);
    
    buffer_destroy(&buffer);
    
    printf("\nAll items produced and consumed successfully!\n");
    
    return 0;
}