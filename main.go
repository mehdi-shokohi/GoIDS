package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	ids "goIDS/idsSource"
	"net/url"
	"regexp"
)
var defaultRules *ids.Rules
func FHandler(ctx *fasthttp.RequestCtx){
	dt,_:=url.QueryUnescape(string(ctx.PostBody()))
	for _,d:=range defaultRules.NewRules().Filters.Filter{

		match, _ := regexp.MatchString(d.Rule, dt)
		if match{
			fmt.Println(d.Description)
			fmt.Println(d.ID)
			fmt.Println(d.Tags.Tag)
			fmt.Println(d.Impact)
			ctx.Write([]byte(d.Description))
			return
		}
	}
	ctx.Write([]byte("OK"))
}

func main(){

	defaultRules=new(ids.Rules)


	fmt.Println("Server Is Running On Port 8080")
	fasthttp.ListenAndServe(":8080", FHandler)

}
