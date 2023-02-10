var express = require('express');
var router = express.Router();
var AuthenticationClient = require('auth0').AuthenticationClient;

var auth0 = new AuthenticationClient({
  domain: '{dev-aa5pftowvhf5q3p2.us.auth0.com',
  clientId: 'SSYQ2dPBMNtkZHt3FQRuSCkGTJIavqhF',
});

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});



  
module.exports = router;
