require('dotenv').config();
const { expressjwt: expressJwt } = require('express-jwt');
const jwks = require('jwks-rsa');

var jwtCheck = expressJwt({
    secret: jwks.expressJwtSecret({
      cache : true,
      ratelimit : true,
      jwksRequestsPerMinute : 5,
      jwksUri : "https://"+process.env.AUTH0BASEURL+"/.well-known/jwks.json"
    }),
    audience : process.env.AUTH0AUDIENCE,
    issuer : `https://${process.env.AUTH0BASEURL}/`,
    algorithms : ['RS256']
  })

module.exports = jwtCheck;