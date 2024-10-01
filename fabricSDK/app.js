const express = require('express');
const { createProduct } = require('./app/createProduct');

const enrollProducerAdmin = require('./app/enrollProducerAdmin');

const registerProducerUser = require('./app/registerUser');

const app = express();
app.use(express.json());

app.post('/api/enrollAdmin', (req,res) => {
    

    enrollProducerAdmin()
        .then(() => {
            res.status(200).send('Admin enrolled successfully');
        })
        .catch((error) => {
            res.status(500).send
        });
    });

app.post('/api/registerUser', (req, res) => {
    // retrive userId, affiliation from request body
    userId = req.body.userId;
    affiliation = req.body.affiliation;

    registerProducerUser(userId, affiliation)
        .then(() => {
            res.status(200).send('User registered successfully');
        })
        .catch((error) => {
            res.status(500).send(error);
        });
});






// // rest post for create product
app.post('/api/product', (req, res) => {
    // retrive productid, name_1, description, manufacturingDate, batchNumber from request body
    productid = req.body.productid;
    name_1 = req.body.name;
    description = req.body.description;
    manufacturingDate = req.body.manufacturingDate;
    batchNumber = req.body.batchNumber;
    userId = req.body.userId;

    // call createProduct function from fabricSDK
    // pass productid, name_1, description, manufacturingDate, batchNumber

    createProduct(productid, name_1, description, manufacturingDate, batchNumber , userId)
        .then(() => {
            res.status(200).send('Product created successfully');
        })
        .catch((error) => {
            res.status(500).send(error);
        });


		
});


let port = process.env.PORT || 3000;
app.listen(port, () => console.log(`server listening on port ${port}....`)); 