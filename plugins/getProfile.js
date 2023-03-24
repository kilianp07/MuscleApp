const request = require('request-promise-native');
const verifyToken = require('./decodeToken');
const access_token = require('./managementClient');

// Middleware function to authenticate requests
const authenticate = async (req, res, next) => {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Unauthorized' });
    }
  
    const token = authHeader.substring(7);
    try {
      const decoded = await verifyToken(token);
      console.log(access_token)
      const options = {
        uri: `https://${process.env.AUTH0BASEURL}/api/v2/users/${decoded.sub}`,
        headers: {
          'Authorization':  access_token.body.access_token,
          'Content-Type': 'application/json',
        },
      };
      const user = await request(options);
      req.user = JSON.parse(user);
      next();
    } catch (err) {
      console.error(err);
      return res.status(401).json({ error: 'Unauthorized' });
    }
    return user;
};

module.exports = authenticate;