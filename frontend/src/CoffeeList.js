import React, { useEffect, useState } from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';

const CoffeeList = () => {

    const [products, setProducts] = useState([])

    const readData = () => {
        axios.get('http://localhost:9090/products').then(response => {
            console.log(response.data);

            setProducts(response.data);
        }).catch(function (error){
            console.log(error);
        });
    }

    const getProducts = () => {
        let table = []

        for (let i=0; i < products.length; i++) {

            table.push(
            <tr key={i}>
                <td>{products[i].name}</td>
                <td>{products[i].price}</td>
                <td>{products[i].sku}</td>
            </tr>
            );
        }

        return table
    }

    useEffect(() => {
        readData()
    }, [])
   
    return (
      <div>
          <h1 style={{marginBottom: "40px"}}>Menu</h1>
          <Table>
              <thead>
                  <tr>
                      <th>
                          Name
                      </th>
                      <th>
                          Price
                      </th>
                      <th>
                          SKU
                      </th>
                  </tr>
              </thead>
              <tbody>
                  {getProducts()}
              </tbody>
          </Table>
      </div>
    ) 
}

export default CoffeeList;