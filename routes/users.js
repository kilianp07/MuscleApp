require('dotenv').config();
const axios = require('axios');
var express = require('express');
const authenticate = require('../plugins/getProfile');
var router = express.Router();
jwtCheck = require('../plugins/checkJwt');
const getProfile = require('../plugins/getProfile');


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

// Create /profile route with the getProfile middleware
router.get('/', getProfile, async (req, res) => {
  res.json(req.user);
});



module.exports = router;
