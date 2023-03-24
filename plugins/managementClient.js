var request = require("request");

var options = { 
  method: 'POST',
  url: `https://${process.env.AUTH0BASEURL}/oauth/token`,
  headers: { 'content-type': 'application/json' },
  body: JSON.stringify({
    client_id: process.env.AUTH0_CLIENT_ID,
    client_secret: process.env.AUTH0_CLIENT_SECRET,
    audience: process.env.AUTH0_AUDIENCE,
    grant_type: 'client_credentials'
  })
};


// Export the request function result
module.exports = request(options, function (error, response, body) {
  if (error) throw new Error(error);

  return body.access_token;
});