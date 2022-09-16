const multipleNaive = (arr) => {
    let multiplied = [];
    for (let i = 0; i < arr.length; i++) {
        let j = 0;
        let product = 1;

        for (let j = 0; j < arr.length; j++) {
            if (j !== i) {
                product *= arr[j];
            }
        }
        multiplied.push(product)
    }

    console.log(multiplied)
}
// multipleNaive([1,2,3,4])

function greedyProductsOfAll(arr) {
    const productsOfPrevious = [];
    let previousProducts = 1;

    for (let i = 0; i < arr.length; i++) {
        productsOfPrevious[i] = previousProducts;
        previousProducts *= arr[i] 
    }

    const productsOfAfter = [];
    let afterProducts = 1;

    for (let j = arr.length - 1; j > -1; j--) {
        productsOfAfter[j] = afterProducts;
        afterProducts *= arr[j] 
    }

    const products = [];
    for (let k = 0; k < arr.length; k++) {
        products[k] = productsOfAfter[k] * productsOfPrevious[k];
    }

    console.log(products)
}

// greedyProductsOfAll([1,2,3,4])

// Optimized Time And Space Complexity
function productsOfAll(arr) {
    const products = [];
    let previousProducts = 1;
    let afterProducts = 1;

    for (let i = 0; i < arr.length; i++) {
        products[i] = previousProducts;
        previousProducts *= arr[i]
    }

    for (let j = arr.length -1; j > -1; j--) {
        products[j] = products[j] * afterProducts;
        afterProducts *= arr[j];
    }
    

    console.log(products)
}
productsOfAll([1,2,3,4])