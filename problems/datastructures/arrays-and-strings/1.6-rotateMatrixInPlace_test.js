const assert = require('chai').assert
const rotateMatrix = require('./1.6-rotateMatrixInPlace').rotateMatrix

describe('Testing 1.6-rotateMatrixInPlace', _ => {
    it(`Input \nanticlockwise\n[[1,   2,   3], \n[4,   5,   6], \n[7,   8,   9]]
        \n\n should return \n[[3,   6,   9],    \n[2,   5,   8],    \n[1,   4,   7]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3], [4,   5,   6], 
                [7,   8,   9]], 'anticlockwise'), [[3,   6,   9], [2,   5,   8], 
                [1,   4,   7]])
        })
    
        it(`Input \nclockwise\n[[1,   2,   3], \n[4,   5,   6], \n[7,   8,   9]]
        \n\n should return \n[[7,   4,   1],    \n[8,   5,   2],    \n[9,   6,   3]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3], [4,   5,   6], 
                [7,   8,   9]], 'clockwise'), [[7,   4,   1], [8,   5,   2], 
                [9,   6,   3]])
        })

        it(`Input \n anticlockwise \n[[1,   2,   3,   4], \n[5,   6,   7,   8], 
        \n[9,  10,  11,  12], \n[13,  14,  15,  16]]\n\n should return 
        \n[[4,   8,  12,  16],    \n[3,   7,  11,  15],    \n[2,   6,  10,  14],    
        \n[1,   5,   9,  13]]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[1,   2,   3,   4], [5,   6,   7,   8], 
                [9,  10,  11,  12], [13,  14,  15,  16]], 'anticlockwise'), 
            [[4,   8,  12,  16], [3,   7,  11,  15], [2,   6,  10,  14], 
            [1,   5,   9,  13]])
        })

        it(`Input \n anticlockwise \n[[1,   2,   3,   4,   5], \n[6,   7,   8,   9,  10], \n[11,  12,  13,  14,  15], \n[16,  17,  18,  19,  20], \n[ 21, 22, 23, 24, 25 ]] 
            should return\n[[5,  10,  15,  20,  25],\n[4,   9,  14,  19,  24],\n[3,   8,  13,  18,  23],\n[2,   7,  12,  17,  22],\n[1,   6,  11,  16,  21] ]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[ 1, 2, 3, 4, 5 ], [ 6, 7, 8, 9, 10 ],
                [ 11, 12, 13, 14, 15 ], [ 16, 17, 18, 19, 20 ], [ 21, 22, 23, 24, 25 ] ], 
                'anticlockwise'), 
                [[ 5, 10, 15, 20, 25 ], [ 4, 9, 14, 19, 24 ],
                [ 3, 8, 13, 18, 23 ], [ 2, 7, 12, 17, 22 ],[ 1, 6, 11, 16, 21 ] ])
        })

        it(`Input \n clockwise \n[[1,   2,   3,   4,   5], \n[6,   7,   8,   9,  10], \n[11,  12,  13,  14,  15], \n[16,  17,  18,  19,  20], \n[ 21, 22, 23, 24, 25 ]] 
            should return\n[[21,  16,  11,   6,   1],\n[ 22,  17,  12,   7,   2],\n[ 23,  18,  13,   8,   3],\n[ 24,  19,  14,   9,   4],\n[ 25,  20,  15,  10,   5] ]`,
        function () {
            assert.deepStrictEqual(rotateMatrix([[ 1, 2, 3, 4, 5 ], [ 6, 7, 8, 9, 10 ],
                [ 11, 12, 13, 14, 15 ], [ 16, 17, 18, 19, 20 ], [ 21, 22, 23, 24, 25 ] ], 
                'clockwise'), 
                [ [ 21, 16, 11, 6, 1 ], [ 22, 17, 12, 7, 2 ], [ 23, 18, 13, 8, 3 ],
                [ 24, 19, 14, 9, 4 ], [ 25, 20, 15, 10, 5 ] ])
        })
})