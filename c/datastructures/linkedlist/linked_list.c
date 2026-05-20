#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <stdbool.h>
#include "linked_list.h"

Node *head = NULL;

bool add_node(int val) {
    Node *new_node = malloc(sizeof(Node));
    if (new_node == NULL) return false;
    *new_node = (Node){ .value = val, .next = NULL}; 

    if (head == NULL) {
        head = new_node;
        return true;
    }

    Node *cur = head;
    while (cur->next != NULL) {
        cur = cur->next;
    }
    cur->next = new_node;
    return true;
}

bool remove_node(int idx) {
    return true;
}


void print_list(void) {
    if (head == NULL) return;

    Node *cur = head;
    printf("head->");
    while (cur->next != NULL) {
        printf("{%d}->", cur->value);
        cur = cur->next;
    }
    printf("{%d}->NULL\n", cur->value);
}

