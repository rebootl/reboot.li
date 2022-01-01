import fs from 'fs';
import path from 'path';

// files/paths
const inFile = 'entries-reflectapp.json';
const inFile2 = 'entries-rebootli.json';
const outFile = 'entries.json';

const PORT = 3005;
const BASEURL = `http://localhost:${PORT}/`

const IMAGEDIR = '../mediaserver/media/rebootl/'

// load data
let data = JSON.parse(fs.readFileSync(inFile, 'utf8'));
const data2 = JSON.parse(fs.readFileSync(inFile2, 'utf8'));

// write data
const writeData = (f) => {
  fs.writeFileSync(f, JSON.stringify(data));
};

for (const entry of data) {
  entry.user = 'rebootl';
  delete entry._id;
  /*if (entry.type === 'note')
    entry.type = 'article'*/
  if (entry.type === 'brokenlink') {
    entry.type = 'link';
    entry.id = 'link' + entry.id;
  }
  if (entry.type === 'image') {
    const c = entry.comment;
    for (const i of entry.images) {
      const s = i.filepath.split('/');
      let n;
      if (s[0] === '')
        n = path.join(s[1], 'rebootl', s[2], s[3]);
      else
        n = path.join(s[0], 'rebootl', s[1], s[2]);
      i.filepath = n;
      i.url = new URL(n, BASEURL);
      if (c) i.comment = c;
      else i.comment = '';
    }
  }
/*  if (entry.type === 'link' && entry.hasOwnProperty('url')) {
    entry.text = entry.url;
    delete entry.url;
  }*/
}

data = [ ...data, ...data2 ];

/*const imagedirs = fs.readdirSync(IMAGEDIR);

// files object contains all files names
// log them on console
for (const d of imagedirs) {
  const images = fs.readdirSync(path.join(IMAGEDIR, d));
  for (const i of images) {
    const p = path.join('media/rebootl/', d, i);
    let found = false;
    for (const e of data) {
      if (e.type !== 'image') continue;
      for (const image of e.images) {
        if (image.filepath === p) {
          found = true;
        }
      }
    }
    if (!found) console.log(p);
  }
}*/

writeData(outFile);
