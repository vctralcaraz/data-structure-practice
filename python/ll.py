class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def add(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
            self.size += 1
        else:
            i = self.head
            while i.next is not None:
                i = i.next
            i.next = new_node
            self.size += 1

    def remove(self, data):
        if self.head is None:
            return

        i = self.head

        if i.data == data:
            self.head = i.next
            i = None
            self.size -= 1
            return
        else:
            while i.next is not None:
                if i.next.data == data:
                    j = i.next
                    i.next = j.next
                    j = None
                    self.size -= 1
                    return
                i = i.next

    def print_list(self):
        i = self.head
        while i is not None:
            print(i.data, end=" -> ")
            i = i.next
        print("None")

__all__ = ["LinkedList"]
