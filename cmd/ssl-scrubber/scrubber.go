package main

import (
	"flag"
	"fmt"
	"github.com/rhololkeolke/ssl-go-tools/pkg/persistence"
	// "github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"log"
	"time"
)

var addressVisionLegacy = flag.String("vision-legacy-address", "224.5.23.2:10005", "Multicast address for vision 2010 (legacy)")
var addressVision = flag.String("vision-address", "224.5.23.2:10006", "Multicast address for vision 2014")
var addressReferee = flag.String("referee-address", "224.5.23.1:10003", "Multicast address for referee 2013")

var visionLegacyEnabled = flag.Bool("vision-legacy-enabled", true, "Record legacy vision packages")
var visionEnabled = flag.Bool("vision-enabled", true, "Record vision packages")
var refereeEnabled = flag.Bool("referee-enabled", true, "Record referee packages")

var VisionLegacyType = persistence.MessageType{Id: persistence.MessageSslVision2010, Name: "vision-legacy"}
var VisionType = persistence.MessageType{Id: persistence.MessageSslVision2014, Name: "vision"}
var RefereeType = persistence.MessageType{Id: persistence.MessageSslRefbox2013, Name: "referee"}

var logFile = flag.String("file", "", "The log file to play")
var skipNonRunningStages = flag.Bool("skip", false, "Skip frames while not in a running stage")

func main() {
	flag.Parse()

	if *logFile == "" {
		log.Fatal("Missing logfile")
	}

	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true)
	play_pause_button := tview.NewButton("▶")
	step_back_button := tview.NewButton("↶")
	step_forward_button := tview.NewButton("↷")
	rewind_button := tview.NewButton("⏪")
	fast_forward_button := tview.NewButton("⏩")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(step_back_button, 0, 1, false).
			AddItem(rewind_button, 0, 1, false).
			AddItem(play_pause_button, 0, 1, true).
			AddItem(fast_forward_button, 0, 1, false).
			AddItem(step_forward_button, 0, 1, false),
		0, 1, true)

	reader_chan := make(chan *persistence.Reader)

	go func() {
		reader, err := persistence.NewReader(*logFile)
		if err != nil {
			log.Print(err)
			close(reader_chan)
		}
		reader_chan <- reader
	}()

	go func() {
		throbber_pattern := []string{"▐⠂       ▌",
			"▐⠈       ▌",
			"▐ ⠂      ▌",
			"▐ ⠠      ▌",
			"▐  ⡀     ▌",
			"▐  ⠠     ▌",
			"▐   ⠂    ▌",
			"▐   ⠈    ▌",
			"▐    ⠂   ▌",
			"▐    ⠠   ▌",
			"▐     ⡀  ▌",
			"▐     ⠠  ▌",
			"▐      ⠂ ▌",
			"▐      ⠈ ▌",
			"▐       ⠂▌",
			"▐       ⠠▌",
			"▐       ⡀▌",
			"▐      ⠠ ▌",
			"▐      ⠂ ▌",
			"▐     ⠈  ▌",
			"▐     ⠂  ▌",
			"▐    ⠠   ▌",
			"▐    ⡀   ▌",
			"▐   ⠠    ▌",
			"▐   ⠂    ▌",
			"▐  ⠈     ▌",
			"▐  ⠂     ▌",
			"▐ ⠠      ▌",
			"▐ ⡀      ▌",
			"▐⠠       ▌"}
		throbber_index := 0

		textView.Clear()
		fmt.Fprintf(textView,
			"Counting messages in [blue]'%s'[white]...[purple]%s[white]\n\n",
			*logFile, throbber_pattern[throbber_index%len(throbber_pattern)])
		throbber_index += 1
		time.Sleep(1 * time.Millisecond)

		finished := false
		for !finished {
			textView.Clear()

			select {
			case reader, ok := <-reader_chan:
				if ok {
					fmt.Fprintf(textView, "[blue]'%s'[white] has [green]%d[white] messages.", *logFile, reader.NumMessages())
				} else {
					fmt.Fprintf(textView, "[blue]'%s'[white]. [red]Error opening file and creating index[white]", *logFile)
				}
				finished = true
			case <-time.After(80 * time.Millisecond):
				fmt.Fprintf(textView,
					"Opening log [blue]'%s'[white]. May take a few moments to build the index...[purple]%s[white]\n\n",
					*logFile, throbber_pattern[throbber_index%len(throbber_pattern)])
				throbber_index += 1

			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
