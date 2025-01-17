package main

import "fmt"

// Existing implementation
type AlgoliaProducts struct {
	productVariants []string
}

// Define interface to support both existing and new structures
type ProductsInterface interface {
	getProductVariants()
}

// Implement interface for AlgoliaProducts
func (a *AlgoliaProducts) getProductVariants() {
	fmt.Println("Algolia StoreProducts")
}

// New schema
type AtlasProducts struct {
	productVariants []string
}

// Atlas-specific method to represent the data
func (a *AtlasProducts) getAtlasProductVariants() {
	fmt.Println("Atlas StoreProducts")
}

// Adapter to make AtlasProducts compatible with ProductsInterface
type AtlasProductsAdapter struct {
	atlasProducts *AtlasProducts
}

// Implement ProductsInterface for AtlasProductsAdapter
func (a *AtlasProductsAdapter) getProductVariants() {
	a.atlasProducts.getAtlasProductVariants()
}

func main() {
	algolia := &AlgoliaProducts{}
	algolia.getProductVariants()

	atlas := &AtlasProducts{}
	atlasAdapter := &AtlasProductsAdapter{
		atlasProducts: atlas,
	}
	atlasAdapter.getProductVariants()
}
