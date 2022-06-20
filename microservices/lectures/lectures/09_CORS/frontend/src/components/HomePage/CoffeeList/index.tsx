import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Table } from 'react-bootstrap'
import { Product } from '../../../models/product.type'

export const CoffeeList = () => {
  const [products, setProducts] = useState<Product[]>([])

  const readData = () => {
    axios.get("http://localhost:9090/products")
      .then((response) => {
        console.log(response.data)
        setProducts(response.data)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  useEffect(() => {
    return () => {
      readData()
    }
  }, [])


  return (
    <div>
      <h1 style={{ marginBottom: "40px" }}>Menu</h1>
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
          {
            products.map((product) => (
              <tr key={product.id}>
                <td>{product.name}</td>
                <td>{product.price}</td>
                <td>{product.sku}</td>
              </tr>
            ))
          }
        </tbody>
      </Table>
    </div>
  )
}