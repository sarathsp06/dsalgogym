/**
 * Given two strings, find a way to decide whether one is the permutation of the other
 */

 function isPermutation (s1, s2) {
     if (s1.length !== s2.length) return false
     if (s1 === s2) return true
     var checker = {}
     for (var element of s1) {
         if (checker[element]) {
             checker[element] += 1
             continue
         }
         checker[element] = 1
     };

     for (var element of s2){
         if (checker[element]) {
             checker[element] -= 1
             continue
         }
         return false
     }
     return true
 }

exports.isPermutation = isPermutation
