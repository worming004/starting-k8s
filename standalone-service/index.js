const express = require('express')

const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World!\n')
})

app.get('/pod', (req, res) => {
    const podname = process.env.POD_NAME
    res.send(podname + '\n')
  })

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})