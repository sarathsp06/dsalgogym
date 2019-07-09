const assert = require('chai').assert;
const isPermutation = require('./1.2-check_permutation').isPermutation;

describe('Testing 1.2-check_permutation', _ => {
    it('Input "a" and "" must return false', function () {
        assert.isFalse(isPermutation('a', ''));
    });

    it ('Input "a" and "a" must return true', function (){
        assert.isTrue(isPermutation('a', 'a'));
    });

    it ('Input "aabbbc" and "abc" must return false', function () {
        assert.isFalse(isPermutation('aabbbc', 'abc'));
    });

    it ('Input "abc" and "bca" must return true', function (){
        assert.isTrue(isPermutation('abc', 'bca'));
    });
});