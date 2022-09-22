package main

github.com/agambondan/wedding-be/app/lib"
"fmt"
)

func main() {
	//url := "https://twelve-data1.p.rapidapi.com/ipo_calendar?country=IND"
	//url = "https://ht2.ajaib.co.id/api/v1/stock/detail/AGRO/price/"
	//url = "https://indonesia-stock-exchange.vercel.app/api/indices"
	//req, _ := http.NewRequest("GET", url, nil)
	//
	//req.Header.Add("Sec-Fetch-Site", "same-site")
	//req.Header.Add("X-RapidAPI-Key", "f5950e79c4msh5d7ba62bf2c4a95p1f7352jsn17e492bee49f")
	//req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")
	//req.Header.Add("Authorization", "jwt eyJ0eXAiOiJBQ0NFU1MiLCJhbGciOiJFUzUxMiJ9.eyJhamFpYiI6IjYzNjEyMS5XRUIuQUNDRVNTIiwicGxhdGZvcm0iOiJXRUIiLCJhdWQiOiJodHRwczovL2FwcC5hamFpYi5jby5pZCIsImV4cCI6MTY1OTA2NzExNCwianRpIjoiYzA3YTNhMDUtYWU4Yi00MGUyLWI1ZGItMjAyM2U3OGY5YTEwIiwiaXNzIjoiaHR0cHM6Ly9hcHAuYWphaWIuY28uaWQiLCJpYXQiOjE2NTkwNjYyMTQsInN1YiI6IjYzNjEyMSJ9.Ae9-qDQgKFKl-j95KZ1qbyHAFnlv-GhBHAWg4T4DEr0-A7PtmuUqFffuoIY_biyUUvvNC7AWS8xyfY6lgYvW7BTcAHj6PFpsqgGe3lzj3_WgDKzXdDisK5prrdF7fdEMtEN4bNG_ek7ciM1rc8ETETDgDa5o3PDr6qvl-cT5Pji8m2R2")
	//
	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))
	//fmt.Println(req.Referer())
	//rand.Seed(1)
	//fmt.Println(rand.Int())
	fmt.Println(lib.GeneratePassword(6, 0, 6, 0))

}
