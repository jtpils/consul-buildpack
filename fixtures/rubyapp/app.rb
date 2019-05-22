ENV['RACK_ENV'] ||= 'development'

require 'rubygems'
require 'sinatra/base'
require 'json'
require 'typhoeus'

Bundler.require :default, ENV['RACK_ENV'].to_sym

class AppServer < Sinatra::Base
  get '/' do
    "Hi, I'm an app with a consul agent sidecar!"
  end

  get '/config' do
    puts "Sending a request to the consul sidecar at localhost:#{ENV['CONFIG_SERVER_PORT']}/config/"
    response = Typhoeus.get("localhost:#{ENV['CONFIG_SERVER_PORT']}/config/")
    if response.body.size > 0
      puts "Received #{response.body} from the consul sidecar"
      STDOUT.flush
      response.body
    else
      {"error": "error connecting to local consul"}.to_json
    end
  end

  run! if app_file == $0
end