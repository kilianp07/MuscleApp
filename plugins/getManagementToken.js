var axios = require("axios").default;

var options = {
  method: 'POST',
  url: `https://${process.env.AUTH0BASEURL}/oauth/token`,
  headers: {'content-type': 'application/x-www-form-urlencoded'},
  data: new URLSearchParams({
    grant_type: 'client_credentials',
    client_id: 'oJu41RoaYiZwYLQzCMlK1VX9sqpS2vua',
    client_secret: process.env.AUTH0CLIENTSECRET,
    audience: `https://${process.env.AUTH0BASEURL}/api/v2/`
  })
};

function getAccessToken() {
    axios.request(options).then(function (response) {
    console.log(response.data);
    }).catch(function (error) {
    console.error(error);
    });
    return response.data.access_token;
}

module.exports = getAccessToken;

