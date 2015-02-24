#!/usr/bin/env ruby

require 'rubygems'
require 'bundler/setup'

require 'google/api_client'
require 'google_drive'

SCOPE = %w(
  https://www.googleapis.com/auth/drive
  https://spreadsheets.google.com/feeds
).join(' ')

require 'json'
JSON.parse(File.read(ENV['ENV'] || '../config/default.json').gsub("\n",'')).each do |key, value|
  ENV[key] ||= value
end

datafile = ENV['DATAFILE']

client = Google::APIClient.new({
  :application_name    => "ganalyse",
  :application_version => '0.1'
})
client.authorization = Signet::OAuth2::Client.new({
  :token_credential_uri => 'https://accounts.google.com/o/oauth2/token',
  :audience             => 'https://accounts.google.com/o/oauth2/token',
  :scope                => SCOPE,
  :issuer               => ENV['ISSUER'],
  :signing_key          => Google::APIClient::KeyUtils.load_from_pkcs12('../config/client.p12', 'notasecret')
})
client.authorization.fetch_access_token!
access_token = client.authorization.access_token

# Creates a session.
session = GoogleDrive.login_with_oauth(access_token)
spreadsheet = session.spreadsheet_by_key(ENV['SPREADSHEET'])
spreadsheet.worksheets[ENV['WORKSHEET'].to_i].export_as_file(datafile)

puts "Exported to: #{datafile}"
