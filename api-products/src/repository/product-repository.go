package repository

import (
	"fmt"
	"os"

	"github.com/danubiobwm/curso-grpc/api-products/src/pb/products"
	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const filename string = "products.txt"

func (pr *ProductRepository) LoadData() (products.ProductList, error) {
	productList := products.ProductList{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("error unmarshalling data: %v", err)
	}

	return productList, nil
}

func (pr *ProductRepository) SaveData(productList products.ProductList) error {
	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filename, err)
	}
	fmt.Printf("Data saved to %s successfully\n", filename)
	return nil
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := pr.LoadData()
	if err != nil {
		return products.Product{}, fmt.Errorf("error loading data: %v", err)
	}
	product.Id = int32(len(productList.Products) + 1) // Simple ID generation
	productList.Products = append(productList.Products, &product)
	err = pr.SaveData(productList)

	return product, err
}

func (pr *ProductRepository) FildAll() (products.ProductList, error) {
	return pr.LoadData()
}
