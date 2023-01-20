var express = require('express');
var router = express.Router();
const firebaseAuth = require("../src/firebase/firebase");
const StatusCodes = require('http-status-codes');


/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.post('/signup', async function(req, res, next) {
    incomingUser = req.body.user;
    
    // Check if user is correctly filled
    if(!firebaseAuth.incomingUserCorrectlyFilled(incomingUser)) {
      res.status(StatusCodes.BAD_REQUEST).send({message:"User not correctly filled"})
    }
  
    try{
      const user = await firebaseAuth.register(incomingUser.username, incomingUser.email, incomingUser.password);
      res.status(StatusCodes.CREATED).send({createdUser:user, message:"User created"});
    }
    catch(error) {
      res.status(StatusCodes.INTERNAL_SERVER_ERROR).send({message:error.message.split("/")[1]});
      return
    }
});

router.post('/signin', async function(req, res, next) {
    incomingUser = req.body.user;

    // Check if user is correctly filled
    if(!firebaseAuth.incomingUserCorrectlyFilled(incomingUser)) {
      res.status(StatusCodes.BAD_REQUEST).send({message:"User not correctly filled"})
    }

    try{
      const user = await firebaseAuth.signInWithEmailAndPassword(incomingUser.email, incomingUser.password);
      res.status(StatusCodes.OK).send({signedInUser:user, message:"User signed in"});
    }
    catch(error) {
      res.status(StatusCodes.INTERNAL_SERVER_ERROR).send({message:error.message.split("/")[1]});
      return
    }
});

module.exports = router;
