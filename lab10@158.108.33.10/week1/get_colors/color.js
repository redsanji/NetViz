const fs = require("fs");
const jsonData = JSON.parse(fs.readFileSync("./sample.json", "utf8"));
const colors = jsonData.nodes.map(o => o.color);
const uniqColor = [...new Set(colors)];
console.log(uniqColor);
