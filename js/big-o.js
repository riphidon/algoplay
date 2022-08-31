const sumArray = (numbers) => {
    let sum = 0;

    for (let i = 0; i < numbers.length; i++) {
        sum += numbers[i];
    }

    return sum;
}

const array = [7, 42, 8, 15, 23, 32, 100, 92, 64, 16, 21, 48, 52, 84, 170]

const startTime = new Date();
// console.log(sumArray(array));
const endTime = new Date();

const elapsed = endTime - startTime;
// console.log(elapsed);
// elapsed change almost every time it is executed

/********************** Space complexity exemples ********************************/

const calculateAverage = (array) => {
    let sum = 0; // created once

    for (let i = 0; i < array.length; i++) { // counter initialized once
        let number = array[i] // executed on every iteration but memory freed each time so count as 1
        sum += number;
        
    }

    return sum / array.length; // created once

} // calculateAverage O(3) = 0(1)

// console.log(calculateAverage([2, 6, 5, 1])) //3.5

const doubler = (items) => {
    let newArray = [];

    for (let i = 0; i < items.length; i++) {
        newArray.push(items[i]);
        newArray.push(items[i]);
    } 

    return newArray;
} // push 2 times item into array:  O(2n) = O(n)

doubler(['a', 'b', 'c']) // ['a', 'a', 'b', 'b', 'c", 'c']

/**************************** Time complexity  exemples ****************************/
const gallifrey = (n) => {
    for (let a = 0; a < n/2; a++) {
        console.log(a); 
    } // a loop O(n/2) = O(n)

    for (let b = 0; b < n; b++) {
        for (let c = 0; c < n; c++) {
            console.log(b + "," + c);
        }
    } // b loop O(n*n) = O(n²)

} // gallifrey On(n + n²) = O(n²)

// gallifrey(10);

const flux = (n) => {
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < n; j++) {
            console.log(j);
        }
    } // i loop O(3 * n) = O(n)

    for (let k = 0; k < 10000; k++) {
        console.log(k);
    } // k loop O(10000)

} //flux O(n + 10000) = O(n)

// flux(10);


/***********************************************************
 *                                                         *
 *                      RECURSIVE CODE                     *
 *                                                         *
 ***********************************************************/


const depth = (n) => {
    if (n === 0) {
        console.log("Snap out of it!");
        return
    }

    console.log(n)
    depth(n-1); // 10-9-8...-1-0 --> Time: O(n) | Space: O(n) nb of call to the stack
}

// depth(10);

const boom = (n) => {
    if (n < 1) {
        console.log("Explode!");
        return
    }

    console.log(n)
    boom(n-2);  // n/2 function calls. Time and Space : O(n/2) = O(n)
}

// boom(10);

/******************************* Time execution test*********************************** */

const slower = (array) => {
    const uniqueArr = [];

    for (let i = 0; i < array.length; i++) {
        let elt = array[i];
        if(!uniqueArr.includes(elt)) { // includes consume time by iterating through array
            uniqueArr.push(elt);
        } // i loop with includes method inside --> Time n * n : O(n²)
    }
    return uniqueArr;
} // Time O(n²) Space O(n), worst case scenario array contains only uniques.




const faster = (array) => {
    const uniques = new Set();

    for (let i = 0; i < array.length; i++) {
        const elt = array[i];
        uniques.add(elt) // set is hashed data structure = constant time operation O(n)
    }

    return Array.from(uniques); // space for conversion gives n + n = 2n = O(n)

} // Time O(n) Space O(n)

const randomArr = []
for (let i = 0; i < 100; i++) {
    const nb = Math.floor(Math.random() * 100);
    randomArr.push(nb);
}

const slowerStart = new Date();
console.log(slower(randomArr));
const slowerEnd = new Date();
console.log(`slower n^2 finished in ${slowerEnd - slowerStart} ms.`);

const fasterStart = new Date();
console.log(faster(randomArr));
const fasterEnd = new Date();
console.log(`faster n finished in ${fasterEnd - fasterStart} ms.`);


/******************************* Multiple arguments functions *********************************** */

const multi = (m, n) => {
    for (let i = 0; i < m; i++) {
        console.log("ohayou!");
    }

    for (let j = 0; j < n; j++) {
        console.log("matane!")
        
    }
} // O(m+n)

// multi(4, 7);

const okashi = (m, n) => {
    for (let i = 0; i < m.length; i++) {
        for (let j = 0; j < n.length; j++) {
            console.log(m[i], n[j]);
        }
        
    }
} // O(m*n)

// okashi(['strawberry', 'vanillia', 'nutella', 'peanut butter'], ['cookie', 'muffin', 'waffle'])

const undou = (str1, str2) => {
    if (str1.length > str2.length) {
        for (let i = 0; i < str1.length; i++) {
            console.log(str1[i]);   
        }
    } else {
        for (let j = 0; j < str2.length; j++) {
            console.log(str2[j]);
        }
    }
} // O(max(m, n)) || O(n) where n = length of longer str 

// undou("yoga", "weights")