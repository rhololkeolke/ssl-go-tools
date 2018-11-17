package main

import (
	"flag"
	"fmt"
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
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

	reader, err := persistence.NewReader(*logFile)
	if err != nil {
		log.Fatal(err)
	}

	msg_count_chan := make(chan int)

	// count messages
	go func() {
		var count int = 0
		for reader.HasMessage() {
			if _, err := reader.ReadMessage(); err != nil {
				log.Print(err)
				break
			}
			count += 1
		}

		msg_count_chan <- count
		close(msg_count_chan)
	}()

	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
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
			case res := <-msg_count_chan:
				fmt.Fprintf(textView, "[blue]'%s'[white] has [green]%d[white] messages.", *logFile, res)
				finished = true
			case <-time.After(80 * time.Millisecond):
				fmt.Fprintf(textView,
					"Counting messages in [blue]'%s'[white]...[purple]%s[white]\n\n",
					*logFile, throbber_pattern[throbber_index%len(throbber_pattern)])
				throbber_index += 1

			}
			time.Sleep(1 * time.Millisecond)			
		}
	}()
	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}

}
