/**
 * Необходимо написать функцию, которая разделит каждую строку
 * в массиве `words` по строке `separator`.
 * Необходимо вернуть массив получившихся после разделения строк,
 * исключая пустые строки
 */
 export const splitWordsBySeparator = (words, separator) => {
    const result = [];
    for (const str of words) {
        let currentWord = '';
        let i = 0;

        while (i < str.length) {
            let match = true;

            for (let j = 0; j < separator.length; j++) {
                if (i + j >= str.length || str[i + j] !== separator[j]) {
                    match = false;
                    break;
                }
            }

            if (match) {
                if (currentWord.length > 0) {
                    result.push(currentWord);
                }
                currentWord = '';
                i += separator.length;
            } else {
                currentWord += str[i];
                i++;
            }
        }

        if (currentWord.length > 0) {
            result.push(currentWord);
        }
    }

    return result;
};