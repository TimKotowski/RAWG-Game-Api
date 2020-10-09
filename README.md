Created a basic api that gets third part api from RAWG, with a list of all the games from the api. 

Used an idiomatic json/encode package called "json iterator" to parse my json (very fast JSON parsing package), use SQLBoiler for the ORM.

As you will see in the code below, ioutil.ReadAll reading all data from a io.Reader and wait for everything to finish downloading-
Then I get the data back from it and put the matching data from the api and unmarshal the data into a struct called Games, reason being is I want to get only certain data from the api, not everything, so I use a struct to put the data into it. The struct fields have to match data from the api!

```golang
type Results struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Background_Image string `json:"background_image"`
}


type Games struct {
	Next    string    `json:"next"`
	Results []Results `json:"results"`
}

```

You will notice when I decode the response body into NewDecoder I didt do aioutil.ReadAll. One insight that I got as I learned how to use Go is that ReadAll is often inefficient for large readers, in my case with all the data coming in from a api that is sending lots and lots of game titles, can be possibly leaking out memory. When I started out, I used to do JSON parsing like this:

 ```golang
data, err := ioutil.ReadAll(r)

if err != nil {
log.Fatalf("err %v", err)
    return err
}
var results Games
jsonData, err := json.Unmarshal(data, &results)
if err != nil {
log.Fatalf("err %v", errs)
    return err
}
w.Write(jsonData)
 ```
Then, I learned of a much more efficient way of parsing the JSON, using the Decoder type


  ```golang
	var results Games
		if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
			log.Fatalf("errs %v", err)
			return
		}
```
Not only is this more concise, it is much more efficient, both memory and time wise

The decoder doesn't have to allocate a huge byte slice to accommodate for the data read-
It can simply re-use a small buffer which will be used against the Read method to get all the data and parse it. This saves a lot of time in allocations

The JSON Decoder can start parsing data as soon as the first chunk of data comes in, it doesn't have to wait for everything to finish downloading.
