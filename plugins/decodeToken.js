const jwt = require('jsonwebtoken');
const jwksClient = require('jwks-rsa');
require('dotenv').config();

const client = jwksClient({
  jwksUri: `https://${process.env.AUTH0BASEURL}/.well-known/jwks.json`,
});

const verifyToken = (token) => {
  return new Promise((resolve, reject) => {
    const kid = jwt.decode(token, { complete: true }).header.kid;
    client.getSigningKey(kid, (err, key) => {
      if (err) {
        reject(err);
      } else {
        const signingKey = key.publicKey || key.rsaPublicKey;
        jwt.verify(
          token,
          signingKey,
          {
            algorithms: ['RS256'],
          },
          (err, decoded) => {
            if (err) {
              reject(err);
            } else {
              resolve(decoded);
            }
          }
        );
      }
    });
  });
};

module.exports = verifyToken;