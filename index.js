
const isPalindrome = (word) => {
    if (word)
    let middleCharacter = word.length/2;
    let isPlan = true;
    for (let i = 0; i < word.length; i++) {
        if ( i < middleCharacter ) {
            const doesWordMatches = word.charAt(i) === word.charAt(word.length - (i + 1));
            if (!doesWordMatches) {
                isPlan = false;
                break;
            }
        }

    }
    return isPlan;
}

console.log(isPalindrome("abba"))
console.log(isPalindrome("a"))