var format = require('date-format');
var http = require('http');
var express = require('express');
var app = express()
var cors = require('cors');
app.use(cors());
var msg = '';
var obj
app.get('/getBLockDetails', function(req,res1){
  console.log('listening on port 8087! and Block number -' + req.query.blocknumber)
  var options = {
  host: 'localhost',
  port: '7050',
  path: '/chain/blocks/'+ req.query.blocknumber,
  method: 'GET',
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  }

};
console.log(options);
msg ='';
var req1 = http.request(options, function(res) {
  res.setEncoding('utf8');
  res.on('data', function(chunk) {
    msg += chunk;
  });
  res.on('end', function() {
        console.log("Message - " + msg);
        processData(JSON.parse(msg) , res1);
		});  
	});  
	
 req1.end()
});


function processData(data, res1){

var string1;
var strReturn ;
var strValue ;
var stateHash ;
var strPreviousHash;
var timestamp = null ;
var localLedgerCommitTimestamp = null;

if(data.transactions) {
	console.log("Check Result - " + data.transactions[0].payload);
	string1 = Buffer(data.transactions[0].payload, 'base64').toString('ascii');
	substring = "updateAsset";
	strReturn = string1.substring(string1.lastIndexOf("updateAsset"),string1.lastIndexOf(""));
	strValue = strReturn.replace(/[^a-zA-Z0-9-]/g,',');
	console.log("payload value - " + strValue);
} else {
	strValue = "";
}
if(data.stateHash) {
	stateHash = data.stateHash;
	if(data.transactions[0].timestamp.seconds) {
		timestamp = data.transactions[0].timestamp.seconds	
	}
} else {
	stateHash = '';
}

if(data.previousBlockHash) {
	strPreviousHash = data.previousBlockHash;
} else {
	strPreviousHash = '';
}

if(data.nonHashData.localLedgerCommitTimestamp.seconds) {
	
	var t = new Date(1970, 0, 1); // Epoch
    localLedgerCommitTimestamp = t.setSeconds(data.nonHashData.localLedgerCommitTimestamp.seconds);	
} 

var obj = {
            "payload": strValue,
            "stateHash": stateHash,
            "previousBlockHash": strPreviousHash,
            "blocktime": format(new Date(localLedgerCommitTimestamp))
         };
		 
		console.log( "blockdetails output:"+obj);
		
		res1.send(JSON.stringify(obj));
} 

app.listen(8089);


