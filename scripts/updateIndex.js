const fs = require('fs');

const INDENT_CHAR = "\t";

var indexString = "# Index\n\n";

const getIndent = (n) => {
	let indent = "";
	for (var i = 0; i < n; i++) {
		indent += INDENT_CHAR;
	}
	return indent;
};

const walk = (dir, level = 0) => {
	let files = fs.readdirSync(dir).filter(file =>
		fs.statSync(dir + "/" + file).isDirectory()
	);
	files.forEach(file => {
		let filePath = `${dir}/${file}`;
		let link = filePath.replace("./", "");
		indexString += `${getIndent(level)}- [${file}](${link})\n`;
		walk(filePath, level + 1);
	});
};

const build = () => {
	walk("./ctci");
	fs.writeFile("./INDEX.md", indexString, (err) => {
		if (err) {
			console.log("couldn't write to INDEX.md. err:", err);
		} else {
			console.log("INDEX.md successfully updated!");
		}
	});
};

build();
