// src/server.js
import { assetsMiddleware, prerenderedMiddleware, kitMiddleware } from '../build/middlewares.js';
import express from 'express';

const app = express();

const myMiddleware = function (req, res, next) {
	console.log('Hello world!');
  console.log(req)
  
	next();
};

app.use(myMiddleware);

app.get('/no-svelte', (req, res) => {
  console.log(req)
	res.end('This is not Svelte!');
});

app.all('*', assetsMiddleware, prerenderedMiddleware, kitMiddleware);

// Express users can also write in a second way:
// app.use(assetsMiddleware, prerenderedMiddleware, kitMiddleware);

app.listen(3000);
