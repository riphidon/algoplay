class Node {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");
const e = new Node("e");
const f = new Node("f");

a.left = b;
a.right = c;
b.right = e;
b.left = d;
c.right = f;

const g = new Node(10);
const h = new Node(12);
const i = new Node(6);
const j = new Node(2);
const k = new Node(9);
const l = new Node(3);

g.left = h;
g.right = i;
h.left = j;
h.right = k;
i.right = l;


//#region  Breadth first traversal
const breadthFirstPrint = (root) => {
    const queue = [root];
    while (queue.length > 0) {
        let curr = queue.shift();
        console.log(curr.val);
        if (curr.left !== null) {
            queue.push(curr.left);
        }
        if (curr.right !== null) {
            queue.push(curr.right)
        }
    }
}

// breadthFirstPrint(a);

const breadthFirstSearch = (root, target) => {
    if (root.val === target) {
        return true;
    }

    const queue = [root];
    while (queue.length > 0) {
        let curr = queue.shift();
        if (curr.val === target) {
            return true;
        }
        if (curr.left !== null) {
            queue.push(curr.left);
        }
        if (curr.right !== null) {
            queue.push(curr.right)
        }
    }

    return false;
}

// console.log(breadthFirstSearch(a, "c"));
// console.log(breadthFirstSearch(a, "f"));
// console.log(breadthFirstSearch(a, "x"));


const sumTree = (root) => {
    let sum = 0;
    const queue = [root];
    while (queue.length > 0) {
        let curr = queue.shift();
        if (curr.left !== null) {
            queue.push(curr.left);
        }
        if (curr.right !== null) {
            queue.push(curr.right)
        }
        sum += curr.val;
    }

    return sum;
}

// console.log(sumTree(g));

//#endregion

//#region Depth first

// const depthFirstPrint = (root) => {
//     const stack = [root];
//     while (stack.length > 0) {
//         const curr = stack.pop();
//         console.log(curr.val);
//         if (curr.right !== null) stack.push(curr.right);
//         if (curr.left !== null) stack.push(curr.left);
//     }
// }



const preOrder = (curr) => {
    if (curr === null) {
        return;
    }
    console.log(curr.val);
    preOrder(curr.left);
    preOrder(curr.right);

}

const postOrder = (curr) => {
    if (curr === null) {
        return;
    }
    postOrder(curr.left);
    postOrder(curr.right);
    console.log(curr.val);
}
const inOrder = (curr) => {
    if (curr === null) {
        return;
    }
    inOrder(curr.left);
    console.log(curr.val);
    inOrder(curr.right);

}

// preOrder(a);
// postOrder(a);
// inOrder(a);

const sumDepth = (root) => {
    if (root === null) {
        return 0;
    }

    return root.val + sumDepth(root.left) + sumDepth(root.right);
}

console.log(sumDepth(g));


//#endregion