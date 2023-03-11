require('dotenv').config();
const { ManagementClient } = require('auth0');

 // Create an Auth0 Management API client
 const auth0 = new ManagementClient({
    domain: process.env.AUTH0BASEURL,
    clientId: process.env.AUTH0CLIENTID,
    clientSecret: process.env.AUTH0CLIENTSECRET,
  });