require 'rubygems'
require 'bundler/setup'
require 'google/api_client'
require 'google_drive'

SCOPE = %w(
  https://www.googleapis.com/auth/drive
  https://spreadsheets.google.com/feeds
).join(' ')

ISSUER      = "103549132987-1cgpa5muukrlflcsabgn6ntduskl4ubi@developer.gserviceaccount.com"
SPREADSHEET = "1waJ9i_yY8K39zz3_urD4V8TPLefJeg_fiBY1KTjfuG0"
WORKSHEET   = 1
DATAFILE    = "../examples/data.csv"

client = Google::APIClient.new({
  :application_name    => "ganalyse",
  :application_version => '0.1'
})
client.authorization = Signet::OAuth2::Client.new({
  :token_credential_uri => 'https://accounts.google.com/o/oauth2/token',
  :audience             => 'https://accounts.google.com/o/oauth2/token',
  :scope                => SCOPE,
  :issuer               => ISSUER,
  :signing_key          => Google::APIClient::KeyUtils.load_from_pkcs12('../config/client.p12', 'notasecret')
})
client.authorization.fetch_access_token!
access_token = client.authorization.access_token

# Creates a session.
session = GoogleDrive.login_with_oauth(access_token)
spreadsheet = session.spreadsheet_by_key(SPREADSHEET)
spreadsheet.worksheets[WORKSHEET].export_as_file(DATAFILE)

puts "Exported to: #{DATAFILE}"
