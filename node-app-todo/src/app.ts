import express from 'express';

const app = express();
const port = 3000;

app.get('/', (req, res) => {
    res.send(`
    <html>
      <body>
        <h1>Todo App</h1>
        <ul>
          <li>Todo 1</li>
          <li>Todo 2</li>
        </ul>
      </body>
    </html>
  `);
});

app.listen(port, () => {
    console.log(`Server started on port ${port}`);
});