const assert = require('chai').assert
const rotateMatrix = require('./1.6-rotateMatrixInPlace').rotateMatrix

describe('Testing 1.6-rotateMatrixInPlace', _ => {
    it(`Input \nanticlockwise\n[[1,   2,   3], \n[4,   5,   6], \n[7,   8,   9]]\n\n should return 
        \n[[3,   6,   9],    \n[2,   5,   8],    \n[1,   4,   7]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3], [4,   5,   6], [7,   8,   9]], 'anticlockwise'), 
                [[3,   6,   9], [2,   5,   8], [1,   4,   7]])
        })
    
        it(`Input \nclockwise\n[[1,   2,   3], \n[4,   5,   6], \n[7,   8,   9]]\n\n should return 
        \n[[3,   6,   9],    \n[2,   5,   8],    \n[1,   4,   7]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3], [4,   5,   6], [7,   8,   9]], 'anticlockwise'), 
                [[7,   4,   1], [8,   5,   2], [9,   6,   3]])
        })

        it(`Input \n anticlockwise \n[[1,   2,   3,   4], \n[5,   6,   7,   8], \n[9,  10,  11,  12], \n[13,  14,  15,  16]]\n\n should return 
        \n[[4,   8,  12,  16],    \n[3,   7,  11,  15],    \n[2,   6,  10,  14],    \n[1,   5,   9,  13]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3,   4], [5,   6,   7,   8], [9,  10,  11,  12], [13,  14,  15,  16]], 'anticlockwise'), 
            [[4,   8,  12,  16], [3,   7,  11,  15], [2,   6,  10,  14], [1,   5,   9,  13]])
        })
})