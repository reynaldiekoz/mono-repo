const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('<h1 style="color: #28A745;">Hello from Node.js Service! V7</h1>');
});

const PORT = process.env.PORT || 6012;
app.listen(PORT, () => {
  console.log(`Node.js Service is running on port ${PORT}`);
});
