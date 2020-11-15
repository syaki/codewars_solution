/**
 * Given positive integers `a` and `b` as strings,
 * evaluate `a / b` and return the quotient
 * and the remainder as strings in the
 * form `[quotient, remainder]`
 * @param {string} a
 * @param {string} b
 * @returns {string[]} [quotient, remainder]
 */
function divideStrings(a, b) {
    let [quo, rem] = ['0', ''],
        r

    for (let i = 0; i < a.length; i += 1) {
        rem = (rem + a[i]).replace(/^0+/, '')
        r = 0

        while (+rem >= +b && (rem = subtract(rem, b))) r++
        quo += r
    }

    return [!quo.match(/^0+$/) ? quo.replace(/^0+/, '') : '0', rem || '0']
}

/**
 * @param {string} a
 * @param {string} b
 * @returns {string[]} .
 */
function subtract(a, b, r = []) {
    a = a.split('').map((x) => +x)
    b = b.split('').map((x) => +x)
    for (let i = a.length - 1, bi = b.length - 1; i >= 0; i -= 1, bi -= 1) {
        let u = a[i],
            d = b[bi] | 0
        if (u >= d) {
            r.unshift(u - d < 10 ? u - d : (u - d) % 10)
            if (u - d >= 10) a[i - 1] += 1
        } else {
            for (var j = 1; j <= i; j += 1) (a[j - 1] -= 1), (a[j] += 10)
            ;(i += 1), (bi += 1)
        }
    }
    r = r.join('')
    return !r.match(/^0+$/) ? r.replace(/^0+/, '') : '0'
}

console.log(
    divideStrings(
        '8248301309760707934466546273758808203334766040337472693900000',
        '648445648688662163528194923667960069985077090082460000'
    )
)
