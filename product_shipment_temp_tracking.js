/*
 Copyright IBM Corp 2016 All Rights Reserved.


 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
process.env['GOPATH'] = __dirname;

// Include the package from npm:
var util = require('util');
var hfc = require("hfc");
var fs = require('fs');
var express = require('express');
var bodyParser = require('body-parser');
var cors = require('cors');
var config = require('./config');

var userObj;
var app = express();
app.use(bodyParser.json());
app.use(cors());
var chaincodeIDPath = __dirname + "/chaincodeID";
var chain, chaincodeID;
var newUserName;
// Create application/x-www-form-urlencoded parser
var urlencodedParser = bodyParser.urlencoded({
    extended: false
});

console.log(" **** starting HFC sample ****");

// get the addresses from the docker-compose environment
var PEER_ADDRESS = config.CORE_PEER_ADDRESS;
var MEMBERSRVC_ADDRESS = config.MEMBERSRVC_ADDRESS;

// Create a chain object used to interact with the chain.
// You can name it anything you want as it is only used by client.
chain = hfc.newChain(config.chainName);
// Initialize the place to store sensitive private key information
chain.setKeyValStore(hfc.newFileKeyValStore('/tmp/keyValStore'));
// Set the URL to membership services and to the peer
console.log("member services address =" + MEMBERSRVC_ADDRESS);
console.log("peer address =" + PEER_ADDRESS);
chain.setMemberServicesUrl("grpc://" + MEMBERSRVC_ADDRESS);
chain.addPeer("grpc://" + PEER_ADDRESS);


// The following is required when the peer is started in dev mode
// (i.e. with the '--peer-chaincodedev' option)
var mode = config.DEPLOY_MODE;
console.log("DEPLOY_MODE=" + mode);
if (mode === 'dev') {
    chain.setDevMode(true);
    //Deploy will not take long as the chain should already be running
    chain.setDeployWaitTime(10);
} else {
    chain.setDevMode(false);
    //Deploy will take much longer in network mode
    chain.setDeployWaitTime(80);
}

chain.setInvokeWaitTime(5);
//chain.setInvokeWaitTime(10);

// Begin by enrolling the user
//enroll();

// Enroll a user.
function enroll(cb) {
    console.log("enrolling user admin ...");
    // Enroll "admin" which is preregistered in the membersrvc.yaml
    chain.enroll("admin", "Xurw3yU9zI0l", function(err, admin) {
        if (err) {
            console.log("ERROR: failed to register admin: %s", err);
            process.exit(1);
        }
        // Set this user as the chain's registrar which is authorized to register other users.
        chain.setRegistrar(admin);
        var userName = "JohnDoe";
        // registrationRequest
        var registrationRequest = {
            enrollmentID: userName,
            account: "bank_a",
            affiliation: "bank_a"
        };
        chain.registerAndEnroll(registrationRequest, function(error, user) {
            if (error) throw Error(" Failed to register and enroll " + userName + ": " + error);
            console.log("Enrolled %s successfully\n", userName);
            //deploy(user);
            //deploy(user,username1,initbal1,username2,initbal2);
              userObj = user;
                 cb(null, user);
        });
    });
}

init();

function init() {
    // Create a client chain.
//    chain = hfc.newChain(config.chainName);

    // Configure the KeyValStore which is used to store sensitive keys
    // as so it is important to secure this storage.
 //   keyValStorePath = __dirname + "/" + config.KeyValStore;
  //  chain.setKeyValStore(hfc.newFileKeyValStore(keyValStorePath));

//    chain.setMemberServicesUrl(config.caserver.ca_url);
newUserName = config.users[1].username;
    //Check if chaincode is already deployed
    //TODO: Deploy failures aswell returns chaincodeID, How to address such issue?
    if (fileExists(chaincodeIDPath)) {
        // Read chaincodeID and use this for sub sequent Invokes/Queries
        chaincodeID = fs.readFileSync(chaincodeIDPath, 'utf8');
    } else {
        registerAndEnrollUsers();
    }
}

function registerAndEnrollUsers() {
    // Enroll "admin" which is already registered because it is
    // listed in fabric/membersrvc/membersrvc.yaml with it's one time password.
    chain.enroll(config.users[0].username, config.users[0].secret, function(err, admin) {
        if (err) return console.log(util.format("ERROR: failed to register admin, Error : %j \n", err));
        // Set this user as the chain's registrar which is authorized to register other users.
        chain.setRegistrar(admin);

        console.log("\nEnrolled admin successfully\n");

        // registrationRequest
        var registrationRequest = {
            enrollmentID: newUserName,
            affiliation: config.users[1].affiliation
        };
        chain.registerAndEnroll(registrationRequest, function(err, user) {
            if (err) throw Error(" Failed to register and enroll " + newUserName + ": " + err);
            userObj = user;
            console.log("Enrolled %s successfully\n", newUserName);

            //chain.setDeployWaitTime(config.deployWaitTime);
            deployChaincode();

        });
    });
}

function deployChaincode() {
    console.log(util.format("Deploying chaincode ... It will take about %j seconds to deploy \n", chain.getDeployWaitTime()))
    //var args = getArgs(config.deployRequest);
    // Construct the deploy request
    var deployRequest = {
        //chaincodeName: config.CORE_CHAINCODE_ID_NAME,
        chaincodePath: config.deployRequest.chaincodePath,
        fcn: "init",
        args: []
    };

    // Trigger the deploy transaction
    var deployTx = userObj.deploy(deployRequest);

    // Print the deploy results
    deployTx.on('complete', function(results) {
        // Deploy request completed successfully
        chaincodeID = results.chaincodeID;
        console.log(util.format("[ Chaincode ID : ", chaincodeID + " ]\n"));
        console.log(util.format("Successfully deployed chaincode: request=%j, response=%j \n", deployRequest, results));
        // Store chaincode ID to a file        
        fs.writeFileSync(chaincodeIDPath, chaincodeID);

//        invoke();
    });
    deployTx.on('error', function(err) {
        // Deploy request failed
        console.log(util.format("Failed to deploy chaincode: request=%j, error=%j \n", deployRequest, err));
        process.exit(0);
    });
}

// Query chaincode
function query(user, order_id, cb) {
    console.log("querying chaincode ...");
    // Construct a query request
    var order_id1 = order_id;
    var function_name;
    if (order_id1 === '') {
        function_name = 'getAllCompleteOrderDetails';
        var argument = 'uniqueOrder';
    } else {
        function_name = 'getCompleteOrderDetails';
        argument = order_id;
    }
    var queryRequest = {
        chaincodeID: chaincodeID,
        fcn: function_name,
        args: [argument]
    };
    // Issue the query request and listen for events
    var tx = user.query(queryRequest);
    tx.on('complete', function(results) {
        console.log("query completed successfully; results=%j", results);
        cb(null, results);
        //process.exit(0);
    });
    tx.on('error', function(error) {
        console.log("Failed to query chaincode: request=%j, error=%k", queryRequest, error);
        process.exit(1);
    });
}

//Invoke chaincode
function invoke(user, order_id, package_id, current_temperature, location, max_temperature, time, carrier, min_temperature, shipping_address, order_date, events, cb) {
    console.log("invoke chaincode ...");
    // Construct a query request
    var invokeRequest = {
        chaincodeID: chaincodeID,
        fcn: "updateAsset",
        //args: ["updateAsset", "123456", "p123", "10","Test","20","11AM","FedEx"]
        args: [order_id, package_id, current_temperature, location, max_temperature, time, carrier, min_temperature, shipping_address, order_date, events]
    };
    // Issue the invoke request and listen for events
    var tx = user.invoke(invokeRequest);
    tx.on('submitted', function(results) {
        console.log("invoke submitted successfully; results=%j", results);
    });
    tx.on('complete', function(results) {
        console.log("invoke completed successfully; results=%j", results);
        cb(null, results);
        //query(user);
    });
    tx.on('error', function(error) {
        console.log("Failed to invoke chaincode: request=%j, error=%k", invokeRequest, error);
        process.exit(1);
    });
}


//For invoke
app.use(bodyParser.urlencoded({
    extended: false
}));
app.use(bodyParser.json());
app.post('/process_invoke', function(req, res) {
    console.log("Got a POST request for /process_invoke which invokes CC");
    var body = req.body;
    console.log("Body VaLUE :: " + body);
    var order_id = body.order_id;
    var package_id = body.package_id;
    var current_temperature = body.current_temperature;
    var location = body.location;
    var max_temperature = body.max_temperature;
    var time = body.time;
    var carrier = body.carrier;
    var shipping_address = body.shipping_address;
    var order_date = body.order_date;
    var events = body.events;
    var min_temperature = body.min_temperature;
    console.log("********* invoke input values: order_id:%s,package_id:%s,current_temperature:%s", order_id, package_id, current_temperature);
    enroll(function(err, enrolleduser) {
        if (err) {
            console.log('*********** Error from enroll');
            return;
        }
        console.log("******* reponse from enroll %s", enrolleduser.name);
        invoke(enrolleduser, order_id, package_id, current_temperature, location, max_temperature, time, carrier, min_temperature, shipping_address, order_date, events, function(err, results) {
            if (err) {
                console.log('***********Unknown Error');
                return;
            }
            console.log("********* return value from invoke : %s", results.result);
            // Prepare output in JSON format
            response = {
                invokeresult: results.result
            };
            res.end(JSON.stringify(response));
        });
    });
});

//For query
app.get('/process_query', function(req, res) {
    console.log("Got a get request for /process_query which quries CC");
    var order_id = req.query.order_id;
    var string1 = "{\"queryresult\"" + ":[";
    var string2 = "]}";
    console.log("********* query input values: username:%s", order_id);
    enroll(function(err, enrolleduser) {
        if (err) {
            console.log('*********** Error from enroll');
            return;
        }
        console.log("******* reponse from enroll %s", enrolleduser.name);
        query(enrolleduser, order_id, function(err, results) {
            if (err) {
                console.log('***********Unknown Error');
                return;
            }
            console.log("********* return value from query : %s", results.result);
            res.end(string1 + results.result + string2);
        });
    });
});

app.listen(8085);
console.log("Submit GET or POST to http://localhost:8085/data");
// Trigger the deploy transaction

function fileExists(filePath) {
    try {
        return fs.statSync(filePath).isFile();
    } catch (err) {
        return false;
    }
}

