package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	"strings"

	"strconv"

	"fmt"

	"encoding/json"

	"math"
)

type user struct {
    ID     int  `json:"id"`
    Name  string  `json:"name"`
    Location int  `json:"location"`
    Gender  string `json:"gender"`
}

type match struct {
    ID     int  `json:"id"`
    who_likes  int  `json:"who_likes"`
    who_is_liked int  `json:"who_is_liked"`
}

var users []user

var matches []match


func main() {

	userJson := `[{"id":1,"name":"Auguste","location":-19,"gender":"Female"},
	{"id":2,"name":"Ned","location":-25,"gender":"Male"},
	{"id":3,"name":"Saree","location":11,"gender":"Female"},
	{"id":4,"name":"Katheryn","location":-5,"gender":"Female"},
	{"id":5,"name":"Brice","location":-14},
	{"id":6,"name":"Kristofer","location":-19,"gender":"Male"},
	{"id":7,"name":"Crawford","location":-15,"email":"ccrosbie6@psu.edu"},
	{"id":8,"name":"Parker","location":-7,"gender":"Male","email":"plisciandri7@unesco.org"},
	{"id":9,"name":"Bevin","location":28,"email":"bgalier8@google.ru"},
	{"id":10,"name":"Christoffer","location":-25,"gender":"Male"},
	{"id":11,"name":"Marje","location":-9,"gender":"Female","email":"mfoldera@bbc.co.uk"},
	{"id":12,"name":"Lorine","location":7,"gender":"Female","email":"lthomtsonb@studiopress.com"},
	{"id":13,"name":"Monica","location":-22,"gender":"Female","email":"mbarabischc@exblog.jp"},
	{"id":14,"name":"Kattie","location":0,"gender":"Female","email":"kshoppeed@uiuc.edu"},
	{"id":15,"name":"Lindy","location":10,"gender":"Female","email":"lfrymane@elegantthemes.com"},
	{"id":16,"name":"Amity","location":-1,"gender":"Female"},
	{"id":17,"name":"Dorie","location":-9,"gender":"Male"},
	{"id":18,"name":"Shelby","location":-7,"gender":"Male"},
	{"id":19,"name":"Mame","location":-18,"gender":"Female","email":"mcommingsi@wikispaces.com"},
	{"id":20,"name":"Arv","location":-4,"gender":"Male"},
	{"id":21,"name":"Berkley","location":-10,"gender":"Male"},
	{"id":22,"name":"Huntley","location":-5,"gender":"Male"},
	{"id":23,"name":"Iago","location":-1,"gender":"Non-binary","email":"igrishankovm@amazonaws.com"},
	{"id":24,"name":"Ollie","location":-25,"gender":"Female","email":"odeverelln@artisteer.com"},
	{"id":25,"name":"Aliza","location":24},
	{"id":26,"name":"Natassia","location":-19,"gender":"Female"},
	{"id":27,"name":"Harv","location":20},
	{"id":28,"name":"Kimberli","location":15,"gender":"Female"},
	{"id":29,"name":"Amaleta","location":28,"gender":"Female","email":"abambridges@google.ru"},
	{"id":30,"name":"Marita","location":12,"gender":"Female"},
	{"id":31,"name":"Gilbertina","location":-22},
	{"id":32,"name":"Claudian","location":9,"gender":"Male","email":"chammerbergv@japanpost.jp"},
	{"id":33,"name":"Anya","location":6,"email":"akelletw@rediff.com"},
	{"id":34,"name":"Mildrid","location":-5,"gender":"Female","email":"mpersehousex@symantec.com"},
	{"id":35,"name":"Tomasina","location":18,"gender":"Polygender","email":"twelshy@cocolog-nifty.com"},
	{"id":36,"name":"Augustus","location":-30,"gender":"Male"},
	{"id":37,"name":"Anya","location":5,"gender":"Genderfluid","email":"abath10@eventbrite.com"},
	{"id":38,"name":"Wilfrid","location":21,"gender":"Male","email":"wgoodwill11@twitpic.com"},
	{"id":39,"name":"Tess","location":-28,"gender":"Female","email":"tjambrozek12@businessinsider.com"},
	{"id":40,"name":"Dickie","location":-29,"gender":"Male","email":"dcorbet13@etsy.com"},
	{"id":41,"name":"Cirilo","location":26},
	{"id":42,"name":"Gilli","location":21,"gender":"Female"},
	{"id":43,"name":"Franz","location":-25,"gender":"Male"},
	{"id":44,"name":"Eddie","location":18,"gender":"Male"},
	{"id":45,"name":"Rex","location":-20,"gender":"Agender","email":"rmytton18@liveinternet.ru"},
	{"id":46,"name":"Yank","location":-5,"gender":"Male"},
	{"id":47,"name":"Catrina","location":-30,"email":"cpiner1a@networksolutions.com"},
	{"id":48,"name":"Luca","location":2,"gender":"Agender","email":"lmarland1b@smugmug.com"},
	{"id":49,"name":"Natka","location":-15,"email":"nantonetti1c@reverbnation.com"},
	{"id":50,"name":"Normy","location":-5,"gender":"Male","email":"ntruswell1d@google.ca"},
	{"id":51,"name":"Cookie","location":3,"gender":"Female"},
	{"id":52,"name":"Otha","location":-13,"gender":"Female","email":"odance1f@ning.com"},
	{"id":53,"name":"Gina","location":-2,"gender":"Female"},
	{"id":54,"name":"Denise","location":-1,"gender":"Female"},
	{"id":55,"name":"Karney","location":-2},
	{"id":56,"name":"Barty","location":4,"gender":"Agender","email":"bbenbow1j@behance.net"},
	{"id":57,"name":"Pietro","location":11,"gender":"Male","email":"pdesouza1k@rambler.ru"},
	{"id":58,"name":"Bogey","location":-3},
	{"id":59,"name":"Aldis","location":-3,"gender":"Male"},
	{"id":60,"name":"Lenora","location":26,"gender":"Female","email":"lsautter1n@creativecommons.org"},
	{"id":61,"name":"Ainslee","location":-16,"gender":"Female"},
	{"id":62,"name":"Sigvard","location":-4,"gender":"Male"},
	{"id":63,"name":"Nicholle","location":-3,"gender":"Female"},
	{"id":64,"name":"Kelly","location":-19,"gender":"Male","email":"kpretsell1r@jimdo.com"},
	{"id":65,"name":"Daniella","location":18,"gender":"Female"},
	{"id":66,"name":"Lorne","location":10,"gender":"Male","email":"lborrett1t@tuttocitta.it"},
	{"id":67,"name":"Lona","location":-7,"gender":"Female","email":"lgreetham1u@hostgator.com"},
	{"id":68,"name":"Selia","location":-12,"gender":"Female","email":"scradick1v@ebay.com"},
	{"id":69,"name":"Danella","location":25,"gender":"Female"},
	{"id":70,"name":"Sanderson","location":-4,"email":"smourant1x@csmonitor.com"},
	{"id":71,"name":"Idaline","location":-10,"email":"ishepherdson1y@sitemeter.com"},
	{"id":72,"name":"Jemmy","location":-28,"gender":"Female"},
	{"id":73,"name":"Max","location":-3,"email":"mholsey20@disqus.com"},
	{"id":74,"name":"Paulina","location":-15,"gender":"Non-binary"},
	{"id":75,"name":"Ilka","location":23,"gender":"Female"},
	{"id":76,"name":"Roy","location":13,"gender":"Male","email":"ramner23@alibaba.com"},
	{"id":77,"name":"Devy","location":28,"gender":"Male","email":"dbrearton24@cafepress.com"},
	{"id":78,"name":"Port","location":-30,"gender":"Male","email":"pgaskin25@illinois.edu"},
	{"id":79,"name":"Elene","location":16,"gender":"Female"},
	{"id":80,"name":"Hashim","location":-5,"gender":"Male","email":"hdunbabin27@foxnews.com"},
	{"id":81,"name":"Claresta","location":-11,"gender":"Female"},
	{"id":82,"name":"Isidora","location":-9,"gender":"Female"},
	{"id":83,"name":"Cornie","location":-27,"gender":"Female","email":"chousden2a@mashable.com"},
	{"id":84,"name":"Abdel","location":3,"gender":"Male","email":"agreenshields2b@netscape.com"},
	{"id":85,"name":"Gwenore","location":-30,"gender":"Female"},
	{"id":86,"name":"Nona","location":4,"email":"ntuckley2d@cam.ac.uk"},
	{"id":87,"name":"Lorri","location":-8,"gender":"Female","email":"ldeignan2e@stanford.edu"},
	{"id":88,"name":"Rayna","location":-7,"gender":"Female"},
	{"id":89,"name":"Hillyer","location":20,"gender":"Male","email":"hpalmer2g@cnbc.com"},
	{"id":90,"name":"Ruthe","location":-9},
	{"id":91,"name":"Sterne","location":22,"email":"sivanishin2i@theatlantic.com"},
	{"id":92,"name":"Carlie","location":-6,"email":"cbridgens2j@google.ca"},
	{"id":93,"name":"Avril","location":-26,"gender":"Female"},
	{"id":94,"name":"Izabel","location":16,"gender":"Female","email":"iheymann2l@china.com.cn"},
	{"id":95,"name":"Adolphus","location":6,"gender":"Male"},
	{"id":96,"name":"Bonni","location":26,"gender":"Female"},
	{"id":97,"name":"Tobi","location":0,"gender":"Female","email":"tchampney2o@va.gov"},
	{"id":98,"name":"Giacomo","location":2,"gender":"Male"},
	{"id":99,"name":"Federica","location":-24,"gender":"Female","email":"fbunworth2q@mayoclinic.com"},
	{"id":100,"name":"Jacques","location":-11,"gender":"Male","email":"jespie2r@cisco.com"}]`

	json.Unmarshal([]byte(userJson), &users)

	matchJson := `[{"id":1,"who_likes":11,"who_is_liked":81},
	{"id":2,"who_likes":36,"who_is_liked":33},
	{"id":3,"who_likes":20,"who_is_liked":32},
	{"id":4,"who_likes":26,"who_is_liked":61},
	{"id":5,"who_likes":78,"who_is_liked":33},
	{"id":6,"who_likes":8,"who_is_liked":34},
	{"id":7,"who_likes":91,"who_is_liked":88},
	{"id":8,"who_likes":68,"who_is_liked":72},
	{"id":9,"who_likes":15,"who_is_liked":28},
	{"id":10,"who_likes":86,"who_is_liked":62},
	{"id":11,"who_likes":67,"who_is_liked":2},
	{"id":12,"who_likes":20,"who_is_liked":83},
	{"id":13,"who_likes":37,"who_is_liked":44},
	{"id":14,"who_likes":67,"who_is_liked":35},
	{"id":15,"who_likes":36,"who_is_liked":66},
	{"id":16,"who_likes":42,"who_is_liked":32},
	{"id":17,"who_likes":9,"who_is_liked":18},
	{"id":18,"who_likes":96,"who_is_liked":26},
	{"id":19,"who_likes":66,"who_is_liked":34},
	{"id":20,"who_likes":81,"who_is_liked":11},
	{"id":21,"who_likes":20,"who_is_liked":81},
	{"id":22,"who_likes":55,"who_is_liked":11},
	{"id":23,"who_likes":75,"who_is_liked":78},
	{"id":24,"who_likes":16,"who_is_liked":92},
	{"id":25,"who_likes":75,"who_is_liked":85},
	{"id":26,"who_likes":66,"who_is_liked":36},
	{"id":27,"who_likes":51,"who_is_liked":64},
	{"id":28,"who_likes":92,"who_is_liked":23},
	{"id":29,"who_likes":2,"who_is_liked":67},
	{"id":30,"who_likes":34,"who_is_liked":8}]`

	json.Unmarshal([]byte(matchJson), &matches)

    router := gin.Default()
    router.GET("/users", getUsers)
	router.GET("/matches", getMatches)
	router.GET("/near-matches/:user/:k", getNearMatches)
	router.GET("/get-similar-name/:q", getUserWithName)

    router.Run("localhost:3000")
}

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}


func getMatches(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, matches)
}

func getNearMatches(c *gin.Context){
	user_id, err := strconv.Atoi(c.Param("user"))
	k, err := strconv.Atoi(c.Param("k"))

	if err != nil {
        fmt.Println(err)
        return
    }

	user_loc := 0

	for _, a := range users {
		if a.ID == user_id {
			user_loc = a.Location
		}
	}

	var near_users []user

	for _, a := range users {
		if math.Abs(float64(a.Location - user_loc)) <= float64(k){
			near_users = append(near_users, a)
		}
	}

	c.IndentedJSON(http.StatusOK, near_users)
	return

}

func getUserWithName(c *gin.Context){
	q := c.Param("q")

	var users_with_q []user

	for _, a := range users {
		if strings.Contains(strings.ToLower(a.Name), strings.ToLower(q)){
			users_with_q = append(users_with_q, a)
		}
	}

	c.IndentedJSON(http.StatusOK, users_with_q)
	return
}
