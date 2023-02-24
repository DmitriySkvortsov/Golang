package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type tableData struct {
	HttpCode, ShortDecript, Description string
}

func main() {
	fmt.Println("work")
	c := colly.NewCollector(
		colly.AllowedDomains("developer.roman.grinyov.name"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	var employeeData []tableData

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnHTML("table.normal-th > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tableData := tableData{
				HttpCode:     el.ChildText("td:nth-child(1)"),
				ShortDecript: el.ChildText("td:nth-child(2)"),
				Description:  el.ChildText("td:nth-child(3)"),
			}
			employeeData = append(employeeData, tableData)

		})
	})
	//fmt.Println(tableData.HttpCode)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})
	err := c.Visit("https://developer.roman.grinyov.name/blog/80")
	if err != nil {
		return
	}
	fmt.Println(employeeData[4][HttpCode])
	/*
		content, err := json.Marshal(employeeData)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = os.WriteFile("employees.json", content, 0644)
		if err != nil {
			return
		}
		fmt.Println("Total employees: ", len(employeeData))

		data, err := ioutil.ReadFile("client_secret_881351674676-29t7ikt7v2kmk5o0uefvs0eusvbhouk4.apps.googleusercontent.com.json")
		checkError(err)
		conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
		checkError(err)

		client := conf.Client(context.TODO())
		service := spreadsheet.NewServiceWithClient(client)
		table, err := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
			Properties: spreadsheet.Properties{
				Title: "Коды ответов сервера",
			},
		}
		checkError(err)
		sheet, err := table.SheetByIndex(0)
		checkError(err)

		lencolumns := len(employeeData)-1
		for _, row := range {
			for coloumn := lencolumns; coloumn > 0 ; coloumn--{
				sheet.Update(row,coloumn,employeeData[row][coloumn])
			}
		}

	*/
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
