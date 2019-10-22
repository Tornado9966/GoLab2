// This file contains examples of scenarios implementation using
// the SDK for menu management.

const menu = require('./menu/client');

const client = menu.Client('http://localhost:8080');

// Scenario 1: Display available menu.
client.listDishes()
    .then((list) => {
        console.log('\033[36m============ Scenario 1 ============');
        console.log('Available menu:\033[32m');
        list.forEach((d) => console.log(d.name+'\t\t', d.price+'\t\t'));
    })
    .catch((e) => {
        console.log(`Problem listing available menu: ${e.message}`);
    });

// Scenario 2: Create order.
client.createOrder([1, 2, 10, 5], 12)
    .then((resp) => {
        console.log('\x1b[36m============ Scenario 2 ============');
        console.log('------------ Order ------------');
		console.log('\x1b[35mTable number: ', resp.Table);
		resp.Dishes.forEach((d) => console.log(d.name+'\t\t', d.price+'\t\t'));
		console.log('\x1B[31mGeneral Sum: ', resp.GenSum);
		console.log('General Sum without VAT: ', resp.GenSumWithoutVAT.toFixed(2));
		console.log('Tips: ', resp.Tips.toFixed(2));
    })
    .catch((e) => {
        console.log(`Problem creating a new menu: ${e.message}`);
    });
