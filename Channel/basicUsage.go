package main

import (
	"net/http"
	"fmt"
	"time"
)

//Following segment of code demonstrate how to use spawn a go channel
//to perform concurrent programming.
func main () {

	//Str array
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//Channel to communiate between Go routine
	c := make(chan string)


	//Spawn checkLink function to channel (Thread)?
	//Noted that in here the for loop only responsible to spawn out task to go channel but not responsible to
	//wait all result to come back before main thread termination. Hence we wouldn't have chance to see the result
	//cuz it terminate before the function return with a result?
	for _, link := range links {
		go checkLink(link, c)
	}


	//This case is kind of complicate, basically doing the same thing as above, but this time
	for l := range c {
		//While url got loaded into channel c, the loop would resume and the var got asign to
		//l then l pass into inline-func as link(String). Finally the func would first sleep f
		//for 2s then call checkLink, as all this operaiton is happend under the scope of inline
		//function and which is named with keyword go so it will execute in a seperate routine aside from
		//main thread.
		go func(link string) {
			//Sleep for 2 seconds
			time.Sleep(2*time.Second)
			checkLink(link, c)
		}(l)
	}

}

//It's a ordinary func which issue a get request to given URL and check for the response status,
//err mean non 200 status.
func checkLink(link string, c chan string ){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}