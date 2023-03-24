var express = require('express');
var router = express.Router();
const { Sequelize } = require("sequelize");
const sequelize = new Sequelize(process.env.DB_NAME, process.env.DB_USER, process.env.DB_PASSWORD,
  {
  dialect: 'mysql'
  }
)
const Weight = require('../models/weight')(sequelize, Sequelize.DataTypes,Sequelize.Model);

// Get the weight by the user id contained in the JWT
router.get('/latest', function(req, res, next) {
    // Get the user id from the JWT
    const userId = req.user.sub;
    // Get the latest weight for the user
    Weight.findOne({
        where: {
            userId: userId
        },
        order: [
            ['timestamp', 'DESC']
        ]
    }).then(weight => {
        // Return the weight
        res.json(weight);
    }
    ).catch(err => {
        // Return the error
        res.status(500).json(err.message);
    }
    );
});

// Get the weight by the user id contained in the JWT
router.get('/all', function(req, res, next) {
    // Get the user id from the JWT
    const userId = req.user.sub;
    // Get all weights for the user
    Weight.findAll({
        where: {
            UserId: userId
        },
        order: [
            ['timestamp', 'DESC']
        ]
    }).then(weights => {
        // Return the weight
        res.json(weights);
    }).catch(err => {
        // Return the error
        res.status(500).json(err.message);
    });
});

router.post('/add', function(req, res, next) {
    // Get the user id from the JWT
    const userId = req.user.sub;
    // Get the weight from the request body
    const weight = req.body.weight;

    if (!Weight.incomingCorrectlyFilled(weight)) {
        res.status(500).json("Weight not correctly filled");
    }
    
    Weight.incomingCorrectlyFilled(weight).then(() => {
        // Create the weight
        Weight.create({
            value: weight.value,
            timestamp: weight.timestamp,
            UserId: userId
        }).then(weight => {
            // Return the weight
            res.json(weight);
        }).catch(err => {
            // Return the error
            res.status(500).json(err.message);
        });
    }).catch(err => {
        // Return the error
        res.status(500).json(err.message);
    });
});

router.put('/update', function(req, res, next) {
    // Get the user id from the JWT
    const userId = req.user.sub;
    // Get the weight from the request body
    const weight = req.body.weight;
    // Get the weight id from the request body
    const id = weight.id;

    if (!Weight.incomingCorrectlyFilled(weight) && !id) {
        res.status(500).json("Weight not correctly filled");
    }

    Weight.update({
        value: weight.value,
        timestamp: weight.timestamp,
        UserId: weight.userId
    }, {
        where: {
            id: id
        }
    }).then(() => {
        // Return the weight
        res.status(200).json("Weight updated");
    }).catch(err => {
        // Return the error
        res.status(500).json(err.message);
    });
});

router.delete('/delete', function(req, res, next) {
    // Get the user id from the JWT
    const userId = req.user.sub;
    // Get the weight from the request body
    const weight = req.body.weight;
    // Get the weight id from the request body
    const id = weight.id;

    if (!id) {
        res.status(500).json("Weight not correctly filled");
    }

    Weight.update({
        isDeleted: true
    }, {
        where: {
            id: id
        }
    }).then(() => {
        // Return the weight
        res.status(200).json("Weight deleted");
    }).catch(err => {
        // Return the error
        res.status(500).json(err.message);
    });
});


module.exports = router;