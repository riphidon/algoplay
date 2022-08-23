class Node {
    constructor(value) {
        this.value = value;
        this.next = null;
    }
}

class LinkedList {
    constructor() {
        this.head = null;
    }

    // append(val) {
    //     if (this.head === null) {
    //         this.head = new Node(val);
    //         return;
    //     }
    //     let curr = this.head;
    //     while (curr.next !== null) {
    //         curr = curr.next;
    //     }
    //     curr.next = new Node(val);
    // }

    append(val) {
        if (this.head === null) {
            this.head = new Node(val);
            return;
        }
       this._append(this.head, val);
    }

    _append(curr, val) {
        if(curr.next === null) {
            curr.next = new Node(val);
            return
        }
        this._append(curr.next, val)
    }

    // print() {
    //     let str = "";
    //     let curr = this.head;
    //     while (curr !== null) {
    //         str += curr.value + "-->";
    //         curr = curr.next
    //     }
    //     console.log(str);
    // }

    print() {
        const str = this._print(this.head);
        console.log(str);
    }

    _print(curr, str) {
        if (curr === null) {
            return '';
        }
        return curr.value + "-->" + this._print(curr.next)
    }

    // contains(target) {
    //     let curr = this.head;
    //     while (curr !== null) {
    //         if (curr.value === target) {
    //             return true;
    //         }
    //         curr = curr.next;
    //     }

    //     return false;
    // }

    contains(target) {
       return this._contains(this.head, target)
    }

    _contains(curr, target) {
        if (curr === null) {
            return false;
        }
        if (curr.value === target) {
            return true
        }

        return this._contains(curr.next, target)
    }

}

const list = new LinkedList();
list.append(5);
list.append(23);
list.append(2);
list.append(12);

// console.log(list.head)

// const sumList = (head) => {
//     let curr = head;
//     let sum = 0;
//     while (curr !== null) {
//         sum += curr.value;
//         curr = curr.next;
//     }

//     return sum;

// }

const sumList = (head) => {
    if (head === null) {
        return 0;
    }
    return head.value + sumList(head.next);
}
console.log(sumList(list.head));
list.print();


const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");

a.next = b;
b.next = c;
c.next = d;

const print = (head) => {
    let curr = head;
    while (curr !== null) {
        console.log(curr.value);
        curr = curr.next;
    }
}

// list.print();
// console.log(list.contains("c"));
// console.log(list.contains("a"));
// console.log(list.contains("b"));
// console.log(list.contains("z"));

//#region DELETE
// const deleteValue = (head, target) => {
//     if (head.value === target) {
//         return head.next;
//     }
//     let prev = null;
//     let curr = head;
//     while (curr !== null) {
//         if (curr.value === target) {
//             prev.next = curr.next;
//         }
//         prev = curr;
//         curr = curr.next;
//     }
//      return head
// }


const deleteValue = (head, target) => {
    if (head.value === target) {
        return head.next;
    }
    
    _deleteValue(head.next, target, head)

    return head;
}

const _deleteValue = (curr, target, previous) => {
    if (curr === null) {
        return;
    }
    if (curr.value === target) {
        previous.next = curr.next;
        return;
    }
    _deleteValue(curr.next, target, curr);
}


const newHead = deleteValue(a, "a");
// print(newHead);
//#endregion