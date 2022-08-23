
//#region DATASETS
const vertices = ["A", "B", "C", "D", "E"];

const vertexIdxs = {
    "A": 0,
    "B": 1, 
    "C": 2, 
    "D": 3, 
    "E": 4
}
const edges = [
    ["A", "B"],
    ["A", "D"],
    ["B", "C"],
    ["C", "D"],
    ["C", "E"],
    ["D", "E"]
];

const adjacencyMatrix = [
    [0, 1, 0, 1, 0],
    [1, 0, 1, 0, 0],
    [0, 1, 0, 1, 1],
    [1, 0, 1, 0, 1],
    [0, 0, 1, 1, 0]
];

const adjencyList = [
    ["B", "D"],
    ["A", "C"],
    ["B", "D", "E"],
    ["C", "D"],
    ["A", "C", "E"],
    ["C", "D"]
]




//#endregion
class Node {
    constructor(value) {
        this.value = value;
        this.edgesList = []; 
    }

    connect(node) {
        this.edgesList.push(node);
        node.edgesList.push(this);
    }

    getAdjacentNodes() {
        return this.edgesList.map(n => n.value);
    }

    isConnected(node) {
        return this.edgesList.some((n) => {
            return node.value === n.value;
        })
    }
}

class Graph {
    constructor(nodes) {
        this.nodes = [...nodes]
    }
    addToGraph(node) {
        this.nodes.push(node);
    }

    breadthFirstTraversal(start, end) {
        const queue = [start];
    
        const visited = new Set();
        visited.add(start);
        while(queue.length > 0) {
            const node = queue.shift();

            if (node.value === end.value) {
                console.log("Found it!");
                return;
            }

            for (const edge of node.edgesList) {
                if (!visited.has(edge)) {
                    queue.push(edge);
                    visited.add(edge);
                }
            }

            console.log(node.value)
        }
    }

    findShortestPath(start, end) {
        const queue = [start];
        const visitedNodes = {};
        visitedNodes[start.value] = null;

        while (queue.length > 0) {
            const node = queue.shift();
            if (node.value === end.value){
               return this.reconstructPath(visitedNodes, start, end);
   
            }
            

            for (const adjacency of node.edgesList) {
                if (!visitedNodes.hasOwnProperty(adjacency.value)) {
                    visitedNodes[adjacency.value] = node;
                    queue.push(adjacency);
                    console.log(adjacency.value)
                }
            }
        }
       
    }

    reconstructPath(visitedNodes, startNode, endNode) {
        let currNode = endNode;
        let path = [];
        while(currNode !== null) {
            let prev = currNode.value;
            path.unshift(prev)
            currNode = visitedNodes[prev];
        }

        return path.join("-->");
        
    }

    depthFirstSearch(start, end, visited) {
        if (!visited) {
            visited = new Set();
        }
        if (start.value === end.value) {
            console.log("Destination reached at : ", end.value);
            return;
        }
        visited.add(start)
       
        console.log("Visiting -->", start.value)

        for (let adjacency of start.edgesList) {
            if (!visited.has(adjacency)) {
                visited.add(adjacency);
                this.depthFirstSearch(adjacency, end, visited);
            }
        }
        
            

    }

}

//#region EDGES

const findAdjacent = (node) => {
    let adjacents = [];
   for (let edge of edges) {
        const nodeIdx = edge.indexOf(node);
        if (nodeIdx > -1) {
            const adjNode= nodeIdx === 0 ? edge[1]: edge[0];
            adjacents.push(adjNode);
        }
   }

   return adjacents
}

// console.log(findAdjacent("A"));

const isConnected = (node1, node2) => {
    return edges.some((edge) => {
        return edge.indexOf(node1) > -1 && edge.indexOf(node2) > -1;
    })
}

// console.log(isConnected("C", "D"));
// console.log(isConnected("A", "E"));

//#endregion

//#region Matrix
const findAdjenciesMatrix = (node) => {
    const adjacents = [];
    const nodeIdx = vertexIdxs[node];

    adjacencyMatrix[nodeIdx].forEach((elt, index) => {
        if (elt === 1) {
            adjacents.push(vertices[index]);
        }
    })
    return adjacents;
}

// console.log(findAdjenciesMatrix("A"));
// console.log(findAdjenciesMatrix("C"));

const isConnectedMatrix = (node1, node2) => {
    const node1Idx = vertexIdxs[node1];
    const node2Idx = vertexIdxs[node2];
    return !!adjacencyMatrix[node1Idx][node2Idx];
   
}
// console.log(isConnectedMatrix("A", "B"));
// console.log(isConnectedMatrix("A", "E"));

//#endregion

//#region List

const nodeA = new Node('A');
const nodeB = new Node('B');
const nodeC = new Node('C');
const nodeD = new Node('D');
const nodeE = new Node('E');

nodeA.connect(nodeB);
nodeA.connect(nodeD);
nodeB.connect(nodeC);
nodeC.connect(nodeD);
nodeC.connect(nodeE);
nodeD.connect(nodeE);


const graph = new Graph([
    nodeA,
    nodeB,
    nodeC,
    nodeD,
    nodeE
]);

// for (let node of graph.nodes) {
//     console.log(node.value);
//     for (connectedNode of node.edgesList) {
//         console.log(`Node ${node.value} is connected to ${connectedNode.value}`);
//     }
// }

// console.log(nodeA.isConnected(nodeE));
//#endregion

//#region Breadth First and Depth first

const DFW = new Node('DFW');
const JFK = new Node('JFK');
const LAX = new Node('LAX');
const HNL = new Node('HNL');
const SAN = new Node('SAN');
const EWR = new Node('EWR');
const BOS = new Node('BOS');
const MIA = new Node('MIA');
const MCO = new Node('MCO');
const PBI = new Node('PBI');

const airportsGraph = new Graph([
    DFW,
    JFK,
    LAX,
    HNL,
    SAN,
    EWR,
    BOS,
    MIA,
    MCO,
    PBI
]);

DFW.connect(LAX);
DFW.connect(JFK);
LAX.connect(HNL);
LAX.connect(EWR);
LAX.connect(SAN);
JFK.connect(BOS);
JFK.connect(MIA);
MIA.connect(MCO);
MIA.connect(PBI);
MCO.connect(PBI);

// console.log(airportsGraph.findShortestPath(DFW, MIA));
// airportsGraph.depthFirstSearch(DFW, MCO)

// Topological order

//#endregion

