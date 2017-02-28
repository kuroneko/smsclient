package main

import (
	"flag"
	"fmt"
	"os"
	"gopkg.in/kuroneko/transmitsms.v0"
)

var (
	apiKey = flag.String("api-key", "", "BurstSMS API Key")
	apiSecret = flag.String("api-secret", "", "BurstSMS API Secret")
	message = flag.String("message", "", "Message to Send")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s [<flags>] <number> [<number>...]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	api := &transmitsms.SMSApi{
		BaseURL: "http://api.transmitsms.com",
		APIKey: *apiKey,
		APISecret: *apiSecret,
	}
	req := &transmitsms.SendSMSRequest{
		Message: *message,
		To: flag.Args(),
	}
	_, err := api.Send(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %s\n", err)
		apierr, ok := err.(*transmitsms.ApiError)
		if ok {
			fmt.Fprintf(os.Stderr, "HTTP Resp: %d\n", apierr.HttpCode)
			fmt.Fprintf(os.Stderr, "Response:\n%s\n", string(apierr.ResponseBody))
		}
		os.Exit(1)
	}
	os.Exit(0)
}