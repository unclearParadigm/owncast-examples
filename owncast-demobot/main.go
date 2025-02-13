package main

import (
	"log"
	"net/http"
)

func main() {
	InitializeBotMemory()
	DemoBotConfiguration = CollectDemoBotConfiguration("config.yaml")

	httpMultiplexer := http.NewServeMux()
	httpMultiplexer.HandleFunc("/auto", RequireHttpPost(RequireJsonContentType(LogRequest(AutoRoute))))
	httpMultiplexer.HandleFunc("/stream/stop", RequireHttpPost(RequireJsonContentType(LogRequest(StreamStartStop))))
	httpMultiplexer.HandleFunc("/stream/start", RequireHttpPost(RequireJsonContentType(LogRequest(StreamStartStop))))
	httpMultiplexer.HandleFunc("/user/join", RequireHttpPost(RequireJsonContentType(LogRequest(UserJoin))))
	httpMultiplexer.HandleFunc("/user/message", RequireHttpPost(RequireJsonContentType(LogRequest(UserMessage))))

	log.Print("Bot is working for " + DemoBotConfiguration.OwncastAddress)
	log.Print("Listening for Webhooks on " + DemoBotConfiguration.ListenAddress)

	err := http.ListenAndServe(DemoBotConfiguration.ListenAddress, httpMultiplexer)
	if err != nil {
		log.Fatal(err)
	}
}
