package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"crypto/md5"
    "encoding/hex"
    "io"

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

type ModuleInfo struct {
	Module string `json:"module"`
}

func getMD5Hash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }

    hashInBytes := hash.Sum(nil)[:16]
    return hex.EncodeToString(hashInBytes), nil
}

func sendHTTPRequest() []Script {
	resp, err := http.Get("https://ryswick.net/scripts/")
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

func downloadFile(filepath string, url string) error {
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

func checkForModuleUpdate() {
	resp, err := http.Get("https://ryswick.net/module/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data ModuleInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := getMD5Hash("COCKBOT.dll")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the module is different from the current one
	if data.Module != hash {
		downloadFile("COCKBOT.dll", "https://ryswick.net/static/COCKBOT.dll")
	}
}

func downloadScript(scriptFile string) {
	resp, err := http.Get("https://ryswick.net/scripts/" + scriptFile + "/")
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

	w.Resize(fyne.NewSize(600, 300))
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
		checkForModuleUpdate();
	})
	updateButton.Resize(fyne.NewSize(updateButton.Size().Width, 150))

	w.SetContent(container.NewBorder(updateButton, downloadButton, nil, nil, container.NewVBox(selectContainer, authorLabel, revisionDateLabel, descriptionLabel)))

	w.ShowAndRun()
}
