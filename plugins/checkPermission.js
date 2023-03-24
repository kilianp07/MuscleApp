require('dotenv').config();
const axios = require('axios');
managementToken = require('./getManagementToken');

function checkPermission(permission) {
  return function(req, res, next) {
    // Get the user's ID from the decoded JWT
    const userId = req.user.sub;
    // Make a request to the Auth0 Management API to get the user's permissions
     const options = {
        method: 'GET',
        url: `https://${process.env.AUTH0BASEURL}/api/v2/users/${userId}/permissions`,
        headers: { 'Authorization': `Bearer ${managementToken()}` }
      };

    // Make a request to the Auth0 Management API to get the user's permissions
    axios(options)
      .then(response => {
        const permissions = response.data;

        // Check if the user has the required permission
        if (permissions.some(p => p.permission_name === permission)) {
          next();
        } else {
          res.status(403).json({ message: "Forbidden" });
        }
      })
      .catch(error => {
        console.error(error);
        res.status(500).json({ message: 'Error retrieving user permissions' });
      });
  }
}

module.exports = checkPermission;