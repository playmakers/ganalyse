#!/usr/bin/env ruby
# encoding: UTF-8

require 'rubygems'
require 'bundler/setup'

require 'pry'
require 'shopify_api'
require 'csv'

require 'json'
JSON.parse(File.read(ENV['ENV'] || '../config/default.json').gsub("\n",'')).each do |key, value|
  ENV[key] ||= value
end

ShopifyAPI::Session.setup({:api_key => ENV['SHOPIFY_APP_API_KEY'], :secret => ENV['SHOPIFY_APP_SECRET']})
session = ShopifyAPI::Session.new(ENV['SHOP'].dup, ENV['TOKEN'].dup)
ShopifyAPI::Base.activate_session(session)

def find(options, name)
  options.each do |option|
    return option.position if option.name == name
  end
  nil
end

def resolveProduct(product_id)
  product = ShopifyAPI::Product.find(product_id)
  colorPos = find(product.options, "Farbe")
  sizePos  = find(product.options, "Grösse") || find(product.options, "Größe")
  posPos   = find(product.options, "Position")
  variants = {}

  product.variants.each do |variant|
    variants[variant.id] = {
      :id           => variant.id,
      :color        => colorPos ? variant.send("option#{colorPos}") : nil ,
      :size         => sizePos  ? variant.send("option#{sizePos}")  : nil,
      :pos          => posPos   ? variant.send("option#{posPos}")   : nil,
      :price        => variant.price,
      :availability => variant.inventory_quantity,
    }
  end

  {
    :id       => product.id,
    :type     => product.product_type,
    :name     => "#{product.vendor} #{product.title}",
    :variants => variants,
  }
rescue ActiveResource::ResourceNotFound
  nil
end

def resolveProducts(datafile)
  products = {}
  CSV.foreach(datafile, :headers => true) do |row|
    _, _, product_id, _ = *row.to_hash.values
    puts product_id

    if product = resolveProduct(product_id)
      products[product_id] = product
    end
  end
  return products
end

def exportForGASetup(products)
  CSV.generate do |csv|
    csv << %w(ga:dimension7 ga:dimension8 ga:dimension2 ga:dimension3 ga:dimension5 ga:dimension4 ga:dimension6)

    products.each do |productId, product|
      product[:variants].each do |variantId, variant|
        csv << [
          productId,
          variantId,
          product[:type],
          product[:name],
          variant[:color],
          variant[:size],
          variant[:pos],
        ]
      end
    end
  end
end

def exportForGAImport(products)
  CSV.generate do |csv|
    products.each do |productId, product|
      product[:variants].each do |variantId, variant|
        csv << [
          productId,
          variantId,
          variant[:color],
          variant[:size],
          variant[:pos],
          variant[:price],
          variant[:availability],
        ]
      end
    end
  end
end

def exportForDefault(products)
  CSV.generate do |csv|
    products.each do |productId, product|
      product[:variants].each do |variantId, variant|
        if variant[:color] == "Schwarz" && variant[:pos] == "FB"
          csv << [
            productId,
            variantId,
            variant[:color],
            variant[:size],
            variant[:pos],
            variant[:price],
            variant[:availability],
          ]
        end
      end
    end
  end
end

products = resolveProducts(ENV['DATAFILE'])

# File.write("../examples/setup.csv", exportForGASetup(products))
# File.write("../examples/import.csv", exportForGAImport(products))
File.write("../examples/default.csv", exportForDefault(products))

# start a REPL session
# binding.pry


