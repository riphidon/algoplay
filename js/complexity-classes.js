/****** 1 - CONSTANT  O(1) *******/

const o1 = (n) => {
    let result = 0;
    for (let i = 0; i < 5; i++) { // fixed nb of iteration
        result += n; // run in constant time
    }
    return result;
} // O(5) = O(1)

o1(4) //20



///////////////////////////////////////////
/****** 2 - LOGARITHMIC  O(log(n)) *******/

const logn = (n) => {
    while (n > 1) {
        console.log(n);
        n /= 2;
    }
    console.log("Finished");
} // O(log(n))

// logn(32)

const recursiveLogn = (n) => {
    if (n < 1) {
        console.log("We reached the bottom of it")
        return
    }
    console.log(n)
    recursiveLogn(n / 2);
}

// recursiveLogn(42);



///////////////////////////////////
/****** 3 - LINEAR  O(n) *******/

const linear = (array) => {
    for (let i = 0; i < array.length; i++) {
        console.log(array[i]);
    }
}

// linear(["ecchus", "overlook", "noctis", "labyrinthus", "chasma"]);


////////////////////////////////////////////
/****** 4 - LOGLINEAR  O(n*log(n)) *******/

const loglinear = (str) => {
    console.log(str);
    if (str.length <= 1) {
        return;
    }
    const midIdx = Math.floor(str.length / 2);
    loglinear(str.slice(0, midIdx)) // nb of recursive call O(log(n)) BUT str.slice creates a copy , worst case scenario n/2
    // ==> n/2 * log(n) = O(n*log(n))
}

// loglinear("abcdefghijklmnopqrstuvwxyz");

const loglinear2 = (array) => {
    let str = "";
    for (let i = 0; i < array.length; i++) {
        str += array[i];
    } // loop : n

    console.log(str);
    console.log("----------------------------");

    if (array.length <= 1) return;

    const midIdx = Math.floor(array.length/2);
    const left = array.slice(0, midIdx);
    const right = array.slice(midIdx)
    loglinear2(left); // recursive calls log(n)
    loglinear2(right);// recursive calls log(n)

}

// loglinear2(['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h']);

//                                  n
//                    <------------------------------->
//                  ^          [abcdefgh]  8            = 8
//                  |             / \                  
//                  |            /   \                 
//         log(n)   |      [abcd] 4    [efgh] 4         = 8     
//                  |       / \           / \             
//                  |  [ab]2  [cd]2    [ef]2 [gh]2      = 8     
//                  |    /\    /\    /\    /\        
//                     a  b  c  d  e  f  g   h          = 8         


///////////////////////////////////////
/****** 5 - POLYNOMIAL  O(n^c) *******/

const poly = (array) => {
    for (let i = 0; i < array.length; i++) {
        for (let j = 0; j < array.length; j++) {
            console.log(array[i] + "/" + array[j])
        }
    }
} // O(n²) if another loop is nested --> O(n^3) etc.

// poly(['chocolate', 'banana', 'biscuit']);

const recursivePoly = (str) => {
    if (str.length === 0) return;

    const firstChar = str[0];
    const rest = str.slice(1);
    console.log(firstChar);

    recursivePoly(rest)
} // n recursive calls * n slice operation = n²

// recursivePoly("banoffee");


///////////////////////////////////////
/****** 6 - EXPONENTIAL  O(c^n) *******/

const expo = (n) => {
    if (n === 1) return;
    expo(n-1);
    expo(n-1);
}

expo(4)

// from one level to the next we double the nb of nodes n times = 2^n
//                4       
//               /  \                  
//              /    \                 
//             3      3      
//            / \    / \             
//           2   2  2   2          
//          /\  /\  /\  /\   
//          11  11  11  11     



///////////////////////////////////////
/****** 7 - FACTORIAL  O(n!) *******/

const facto = (n) => {
    if (n===1) { return 1};
    return n * facto(n-1); // recursive calls nb depends on input size  
}

console.log(facto(4))
