module.exports = async function(input) {
    const result = [];

    const size = folder =>
        new Promise(resolve => folder.size(resolve));

    const read = (folder, index) =>
        new Promise(resolve => folder.read(index, resolve));

    function isBroken(name) {
        if (typeof name !== "string" || name === "file") {
            return false;
        }

        let compressed = "";

        for (const ch of name) {
            if (compressed.length === 0 || compressed[compressed.length - 1] !== ch) {
                compressed += ch;
            }
        }

        return compressed === "file";
    }

    async function dfs(folder) {
        const len = await size(folder);

        for (let i = 0; i < len; i++) {
            const item = await read(folder, i);

            if (typeof item === "string") {
                if (isBroken(item)) {
                    result.push(item);
                }
            } else if (
                item &&
                typeof item.read === "function" &&
                typeof item.size === "function"
            ) {
                await dfs(item);
            }
        }
    }

    await dfs(input);

    result.sort();

    return result;
};