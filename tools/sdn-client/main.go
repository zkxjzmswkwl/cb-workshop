package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Script struct {
	File         string `json:"File"`
	Author       string `json:"Author"`
	Description  string `json:"Description"`
	RevisionDate string `json:"RevisionDate"`
}

func sendHTTPRequest() []Script {
	resp, err := http.Get("https://ryswick.net/sdn/scripts/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data struct {
		Scripts []Script `json:"scripts"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data.Scripts
}

// Save the script with a filename of scriptFile to C:\ProgramData\Jagex\Launcher\cockbot\scripts
func saveScriptToFile(scriptFile string, scriptContent string) error {
	file, err := os.Create("C:\\ProgramData\\Jagex\\Launcher\\cockbot\\scripts\\" + scriptFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(scriptContent)
	if err != nil {
		return err
	}

	return nil
}

func downloadScript(scriptFile string) {
	resp, err := http.Get("https://ryswick.net/sdn/scripts/" + scriptFile + "/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data struct {
		Script string `json:"script"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	err = saveScriptToFile(scriptFile, data.Script)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Lennissa SDN Client")

	w.Resize(fyne.NewSize(800, 600))
	data := sendHTTPRequest()

	authorLabel := widget.NewLabel("")
	revisionDateLabel := widget.NewLabel("")
	descriptionLabel := widget.NewLabel("")

	selectList := widget.NewSelect(func() []string {
		var options []string
		for _, script := range data {
			options = append(options, script.File)
		}
		return options
	}(), func(selected string) {
		for _, script := range data {
			if script.File == selected {
				authorLabel.SetText("Author: " + script.Author)
				revisionDateLabel.SetText("Revision Date: " + script.RevisionDate)
				descriptionLabel.SetText("Description: " + script.Description)
				break
			}
		}
	})

	selectContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), selectList)
	selectContainer.Resize(fyne.NewSize(200, 200))

	downloadButton := widget.NewButton("Download", func() {
		selected := selectList.Selected
		if selected != "" {
			downloadScript(selected)
		}
	})

	updateButton := widget.NewButton("Update", func() {
		data = sendHTTPRequest()
		selectList.Refresh()
	})
	updateButton.Resize(fyne.NewSize(updateButton.Size().Width, 150))

	w.SetContent(container.NewBorder(updateButton, downloadButton, nil, nil, container.NewVBox(selectContainer, authorLabel, revisionDateLabel, descriptionLabel)))

	w.ShowAndRun()
}
