package persistence

import (
	"fmt"
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/sslproto"
	"log"
	"time"
)

type PlaybackMode int

const (
	PlayMode        PlaybackMode = 0
	PauseMode       PlaybackMode = 1
	StepBackMode    PlaybackMode = 2
	StepForwardMode PlaybackMode = 3
	RewindMode      PlaybackMode = 4
	FastForwardMode PlaybackMode = 5
)

type ScrubbableBroadcaster struct {
	Broadcaster
	log *IndexedLog

	setMode chan PlaybackMode
	stop    chan bool
}

func NewScrubbableBroadcaster() ScrubbableBroadcaster {
	return ScrubbableBroadcaster{Broadcaster: NewBroadcaster(), setMode: make(chan PlaybackMode), stop: make(chan bool)}
}

func (b *ScrubbableBroadcaster) Start(log *IndexedLog) {
	b.log = log
	b.stop = make(chan bool)

	for _, slot := range b.Slots {
		conn := connectForPublishing(slot.address)
		b.conns[slot.MessageType.Id] = conn
	}

	go b.publish()
}

func (b *ScrubbableBroadcaster) Play() {
	b.setMode <- PlayMode
}

func (b *ScrubbableBroadcaster) Pause() {
	b.setMode <- PauseMode
}

func (b *ScrubbableBroadcaster) Stop() {
	close(b.stop)

	for _, conn := range b.conns {
		conn.Close()
	}
}

func (b *ScrubbableBroadcaster) publish() {
	startTime := time.Now()
	refTimestamp := int64(0)
	curStage := sslproto.SSL_Referee_Stage(-1)

	var msg *Message
	var err error
	var mode = PauseMode

	var isRunning = true
	for {
		select {
		case mode = <-b.setMode:
			// nothing
		case <-b.stop:
			isRunning = false
			break
		default:
			// nothing
		}
		if isRunning {
			switch mode {
			case PlayMode:
				_, msg, err = b.log.Next()
				if err != nil {
					// TODO(dschwab): better error handling?
					fmt.Println("Error:", err)
				}
				if isRunningStage(curStage) {

					if refTimestamp != 0 {
						realElapsed := time.Now().Sub(startTime).Nanoseconds()
						msgElapsed := msg.Timestamp - refTimestamp
						sleepTime := msgElapsed - realElapsed
						time.Sleep(time.Duration(sleepTime))
					} else {
						startTime = time.Now()
						refTimestamp = msg.Timestamp
					}

					conn := b.conns[msg.MessageType.Id]
					if conn != nil {
						conn.Write(msg.Message)
					}
				} else {
					refTimestamp = 0
				}

				if b.SkipNonRunningStages && msg.MessageType.Id == MessageSslRefbox2013 {
					refMsg, err := msg.ParseReferee()
					if err != nil {
						log.Println("Could not parse referee message:", err)
					} else {
						curStage = *refMsg.Stage
					}
				}

			case PauseMode:
				if msg != nil {
					conn := b.conns[msg.MessageType.Id]
					if conn != nil {
						conn.Write(msg.Message)
					}
				}
				time.Sleep(16 * time.Millisecond)
			case StepBackMode:
				// TODO(dschwab): go back one frame
				mode = PauseMode
			case StepForwardMode:
				// TODO(dschwab): go forward one frame
				mode = PauseMode
			case RewindMode:
				// TODO(dschwab): write me
			case FastForwardMode:
				// TODO(dschwab): write me
			}
		} else {
			break
		}
	}
}
