package sync

import (
  "github.com/playmakers/ganalyse/lib/vendors"

  "github.com/rapito/go-shopify/shopify"
  simplejson "github.com/bitly/go-simplejson"
  "fmt"
  s "strings"
)

type ShopifyProduct struct {
  Id   int
  Name string
  Urls []string
  Variants []*ShopifyVariant
  VendorProducts []*vendors.Product
}

// func (p *ShopifyProduct) UpdateVariant(color string, size string, position string, availability int) {
//   for _, variant := range p.Variants {
//     // fmt.Printf("test: %s %s || %s %s\n", color, variant.Color, size, variant.Size)
//     if color == variant.Color && size == variant.Size {
//       // variant.Availability = availability
//       fmt.Printf(" --> match: %s %s || %s %s\n", color, variant.Color, size, variant.Size)
//       // fmt.Printf("match: %d\n", variant)
//     }
//   }
// }

type ShopifyVariant struct {
  Id   int
  Color string
  Size string
  Position string
  Quantity int
  Policy string
}

func (v *ShopifyVariant) String() string {
  return fmt.Sprintf("id: %d\topt1: %s\topt2: %s\topt3: %s\n", v.Id, v.Color, v.Size, v.Position)
}

func Store(store string, apiKey string, pass string) shopify.Shopify {
  return shopify.New(store, apiKey, pass)
}

func GetProductWithUrls(shopify shopify.Shopify, namespace string) []*ShopifyProduct {
  collection_id := fetchCollectionId(shopify, namespace)
  products := fetchProducts(shopify, collection_id)

  sem := make(chan bool, len(products))

  for _, product := range products {
    go func(product *ShopifyProduct, namespace string) {
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

func fetchProducts(shopify shopify.Shopify, collection_id int) []*ShopifyProduct {
  products := []*ShopifyProduct{}

  result, _ := shopify.Get(fmt.Sprintf("products.json?collection_id=%d&fields=id,vendor,title,variants,options&limit=20#", collection_id))
  jsonData, _ := simplejson.NewJson(result)

  rawProducts := jsonData.Get("products")
  for index, _ := range rawProducts.MustArray() {
    rawProduct := rawProducts.GetIndex(index)
    product := &ShopifyProduct{
      Id: rawProduct.Get("id").MustInt(),
      Name: fmt.Sprintf("%s %s", rawProduct.Get("vendor").MustString(), rawProduct.Get("title").MustString()),
    }

    option1Name := s.Split(s.Replace(rawProduct.Get("options").GetIndex(0).Get("name").MustString(), "ss", "ß", -1), " ")[0]
    option2Name := s.Replace(rawProduct.Get("options").GetIndex(1).Get("name").MustString(), "ss", "ß", -1)
    option3Name := rawProduct.Get("options").GetIndex(2).Get("name").MustString()

    rawVariants := rawProducts.GetIndex(index).Get("variants")
    for vIndex, _ := range rawVariants.MustArray() {
      rawVariant := rawVariants.GetIndex(vIndex)

      options := map[string]string{
        option1Name: rawVariant.Get("option1").MustString(),
        option2Name: rawVariant.Get("option2").MustString(),
        option3Name: rawVariant.Get("option3").MustString(),
      }

      variant := &ShopifyVariant{
        Id:       rawVariant.Get("id").MustInt(),
        Size:     options["Größe"],
        Color:    options["Farbe"],
        Position: options["Position"],
        Quantity: rawVariant.Get("inventory_quantity").MustInt(),
        Policy:   rawVariant.Get("inventory_policy").MustString(),
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
