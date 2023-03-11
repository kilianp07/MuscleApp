require('dotenv').config();
const axios = require('axios');
var express = require('express');
var router = express.Router();
jwtCheck = require('../plugins/checkJwt');


async function getAccessToken() {
  const url = `https://${process.env.AUTH0BASEURL}/oauth/token`;
  const data = {
    grant_type: 'client_credentials',
    client_id: process.env.AUTH0CLIENTID,
    client_secret: process.env.AUTH0CLIENTSECRET,
    audience: `muscleapp`
  };

  const response = await axios.post(url, data);
  return response.data.access_token;
}

router.get('/profile', jwtCheck, (req, res) => {
  const accessToken  = req.headers["authorization"]&& req.headers["authorization"].split(' ')[1]
  //const accessToken = getAccessToken();
  axios.get(`https://${process.env.AUTH0BASEURL}/userinfo`, {
    headers: {
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json'
    }
  })
  .then(response => {
    const userProfile = response.data;

    // Do something with the user profile
    console.log(userProfile);

    res.json(userProfile);
  })
  .catch(error => {
    console.error(error);
    res.status(500).send('Error retrieving user profile');
  });
});

  
module.exports = router;
