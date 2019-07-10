/**
 * ** Modified 1.6 Image rotation problem
 * Given a n by n matrix, can you rotate it in-place 90 degrees ?
 */

function rotateMatrix (matrix, direction) {
    var ringCount = rings(matrix.length)
    for (var x = 0; x < ringCount; x++) {
        matrix = rotateRing(matrix, direction, x)
    }
    return matrix
}

function isOdd (x) {
    return !(x%2 === 0)
}

function rings (length) {
    if(isOdd(length)) return (length - 1) / 2
    return length / 2
}

function rotateRing (matrix, direction, i) { 
    for (var j = i; j < matrix.length - 1 - i; j++) {
        if (direction === 'clockwise') {
            matrix = rotateFourClockwise(matrix, [i, j])
            continue
        }
        matrix = rotateFourAnticlockwise(matrix, [i, j]) 
    }
    return matrix
}

function rotateFourAnticlockwise (matrix, [i, j]) {
    var n = matrix.length - 1;
    var temp1; var temp2;
    temp2 = matrix[n - i - j][i];
    matrix[n - j - i][i] = matrix[i][i + j]
    temp1 = temp2
    temp2 = matrix[n - i][n - i - j]
    matrix[n - i][n - i - j] = temp1
    temp1 = temp2
    temp2 = matrix[i + j][n - i]
    matrix[i + j][n - i] = temp1
    temp1 = temp2
    matrix[i][i + j] = temp1
    return matrix
}

function rotateFourClockwise (matrix, [i, j]) {
    var n = matrix.length - 1;
    var temp1; var temp2;
    temp2 = matrix[i][i + j];
    matrix[i][i + j] = matrix[n - j - i][i]
    temp1 = temp2
    temp2 = matrix[i + j][n - i]
    matrix[i + j][n - i] = temp1
    temp1 = temp2
    temp2 = matrix[n - i][n - i - j]
    matrix[n - i][n - i - j] = temp1
    temp1 = temp2
    matrix[n - i - j][i] = temp1;
    return matrix
}

exports.rotateMatrix = rotateMatrix