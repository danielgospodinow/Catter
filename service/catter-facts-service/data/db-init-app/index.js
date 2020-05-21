const fetch = require('node-fetch')
const fs = require('fs')
const readline = require('readline')

const url = "http://localhost:7001/api/fact"

var lineReader = readline.createInterface({
    input: fs.createReadStream('../facts.json')
})

lineReader.on('line', line => {
    if (line.includes("text") || line.includes("fact")) {
        const data = line.split(":")[1].trim().slice(1, -2)

        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                content: data
            })
        })
            .then(response => response.json())
            .then(response => console.log(response))
            .catch(error => console.log(error))
    }
})
