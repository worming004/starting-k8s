const express = require("express");

process.on("SIGINT", function () {
  process.exit(1);
});

const app = express();
const port = 3000;

app.use(function (req, res, next) {
  console.log(`path reached: ${req.path}`);
  next();
});

app.get("/", (req, res) => {
  res.send("Hello World!\n");
});

app.get("/pod", (req, res) => {
  const podname = process.env.POD_NAME;
  res.send(podname + "\n");
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
