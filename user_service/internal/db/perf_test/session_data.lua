local wrk = require "wrk"

wrk.method = "GET"
wrk.headers["Accept"] = "application/json"
