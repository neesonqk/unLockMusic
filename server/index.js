var express = require('express');

var app = express();

app.get('/', function (req, res) {
	res.sendFile(__dirname + '/pac/pac.js');
})

app.listen(3000);