function sumExcept(a, i, n) {
    if (!Number.isInteger(i) || i < 0) {
        i = 0;
    }

    if (!Number.isInteger(n) || n < 0) {
        n = 0;
    }

    let sum = 0;
    const end = i + n;

    for (let index = 0; index < a.length; index++) {
        if (index >= i && index < end) {
            continue;
        }

        if (Number.isInteger(a[index])) {
            sum += a[index];
        }
    }

    return sum;
}

module.exports = sumExcept