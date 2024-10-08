#include <iostream>
using namespace std;

class Node {
public:
  int data;
  Node *next;

  Node() {
    data = 0;
    next = NULL;
  }

  Node(int data) {
    this->data = data;
    this->next = NULL;
  }
};

class LinkedList {
  Node *head;
  int size;

public:
  LinkedList() {
    head = NULL;
    size = 0;
  }

  // insert data a the head or beginning of the list
  void insertAtHead(int data) {
    Node *newNode = new Node(data);

    if (head == NULL) {
      head = newNode;
      return;
    }

    newNode->next = this->head;
    this->head = newNode;
  }

  // insert data at the end of the list
  void insertAtEnd(int data) {
    Node *newNode = new Node(data);

    if (head == NULL) {
      head = newNode;
      return;
    }

    Node *temp = head;

    while (temp->next != NULL) {
      temp = temp->next;
    }

    temp->next = newNode;
  }

  // get the link at the data
  Node *get(int data) {
    if (head == NULL) {
      cout << "This list is empty" << endl;
      return NULL;
    }

    Node *temp = head;

    while (temp != NULL) {

      if (temp->data == data) {
        return temp;
      }
      temp = temp->next;
    }

    cout << "The item is not in the list" << endl;
    return NULL;
  }

  // print entire list
  void print() {
    Node *temp = head;

    if (temp == NULL) {
      cout << "This list is empty" << endl;
      return;
    }

    while (temp != NULL) {
      cout << temp->data << " ";
      temp = temp->next;
    }
    cout << endl;
  }
};

int main() {
  LinkedList list;

  list.print();

  list.insertAtHead(1);
  list.print();

  cout << endl;

  list.insertAtEnd(2);
  list.insertAtEnd(5);
  list.insertAtEnd(3);
  list.print();

  Node *five = list.get(5);
  if ( five != NULL) {
    cout << "the data at " << five << " is " << five->data << endl;
  }

  Node *eight = list.get(8);
  if (eight != NULL) {
    cout << "the data at " << eight << " is " << eight->data << endl;
  }
  return 0;
}
