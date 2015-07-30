package sync

import (
  "github.com/rapito/go-shopify/shopify"
  simplejson "github.com/bitly/go-simplejson"
  "fmt"
)

type Product struct {
  Id   int
  Urls []string
  Variants []*Variant
}

type Variant struct {
  Id   int
  Options [3]string
}

func Store(store string, apiKey string, pass string) shopify.Shopify {
  return shopify.New(store, apiKey, pass)
}

func GetProductUrls(shopify shopify.Shopify, namespace string) []*Product {
  collection_id := fetchCollectionId(shopify, namespace)
  products := fetchProducts(shopify, collection_id)

  sem := make(chan bool, len(products))

  for _, product := range products {
    go func(product *Product, namespace string) {
      product.Urls = fetchMetafieldValues(shopify, product.Id, namespace)
      sem <- true
    }(product, namespace)
  }

  for _, _ = range products {
    <-sem
  }

  return products
}

func fetchCollectionId(shopify shopify.Shopify, namespace string) int {
  result, _ := shopify.Get(fmt.Sprintf("smart_collections.json?limit=1&handle=%s&fields=id#", namespace))
  jsonData, _ := simplejson.NewJson(result)
  return jsonData.Get("smart_collections").GetIndex(0).Get("id").MustInt()
}

func fetchProducts(shopify shopify.Shopify, collection_id int) []*Product {
  products := []*Product{}

  result, _ := shopify.Get(fmt.Sprintf("products.json?collection_id=%d&fields=id,variants#", collection_id))
  jsonData, _ := simplejson.NewJson(result)

  rawProducts := jsonData.Get("products")
  for index, _ := range rawProducts.MustArray() {
    product := &Product{
      Id: rawProducts.GetIndex(index).Get("id").MustInt(),
    }

    rawVariants := rawProducts.GetIndex(index).Get("variants")
    for vIndex, _ := range rawVariants.MustArray() {
      rawVariant := rawVariants.GetIndex(vIndex)
      variant := &Variant{
        Id: rawVariant.Get("id").MustInt(),
        Options: [...]string{
          rawVariant.Get("option1").MustString(),
          rawVariant.Get("option2").MustString(),
          rawVariant.Get("option3").MustString(),
        },
      }
      product.Variants = append(product.Variants, variant)
    }

    products = append(products, product)
  }

  return products
}

func fetchMetafieldValues(shopify shopify.Shopify, product_id int, namespace string) []string {
  values := []string{}

  result, _ := shopify.Get(fmt.Sprintf("products/%d/metafields.json?namespace=%s&fields=value#", product_id, namespace))
  jsonData, _ := simplejson.NewJson(result)
  metafields := jsonData.Get("metafields")

  for index, _ := range metafields.MustArray() {
    value :=  metafields.GetIndex(index).Get("value").MustString()
    values = append(values, value)
  }

  return values
}
