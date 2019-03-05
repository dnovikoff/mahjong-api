package logs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/golang/protobuf/proto"
)

func SaveJSON(name string, x interface{}) error {
	bytes, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name+".json", bytes, 0644)
}

func SaveReadableJSON(name string, x interface{}) error {
	bytes, err := MarshalReadableJSON(x)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name+".json", bytes, 0644)
}

func MarshalReadableJSON(x interface{}) ([]byte, error) {
	bytes, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	bytes = cleanOneof(bytes)
	bytes, err = formatJson(bytes)
	if err != nil {
		return nil, err
	}
	bytes = instancesR.ReplaceAllFunc(bytes, func(src []byte) []byte {
		src = spaceR.ReplaceAll(src, []byte(` `))
		src = []byte(strings.TrimSpace(string(src)))
		return src

	})
	return bytes, err
}

func SaveJSONLogs(name string, log *public_log.Log, debug *private_log.DebugLog) error {
	log, debug = proto.Clone(log).(*public_log.Log), proto.Clone(debug).(*private_log.DebugLog)
	save := func(name string, saver func(string, interface{}) error) error {
		if err := saver(name+"log", log); err != nil {
			return err
		}
		if debug == nil {
			return nil
		}
		if err := saver(name+"create", debug.Create); err != nil {
			return err
		}
		if err := saver(name+"debug", debug); err != nil {
			return err
		}
		for k, v := range debug.PlayerMessages {
			if err := saver(name+fmt.Sprintf("debug_%v", k+1), v); err != nil {
				return err
			}
		}
		return nil
	}
	if err := save(name, SaveJSON); err != nil {
		return err
	}
	cleanDebug(debug)
	cleanLog(log)
	if err := save(name+"x_", SaveReadableJSON); err != nil {
		return err
	}
	return nil
}

// Remove timestamps from logs for readability
func cleanDebug(sl *private_log.DebugLog) {
	for _, v := range sl.GetPlayerMessages() {
		for _, m := range v.Messages {
			m.Time = nil
		}
	}
}

// Remove timestamps from logs for readability
func cleanLog(log *public_log.Log) {
	log.Started = nil
	log.Ended = nil
	for _, r := range log.GetRounds() {
		for _, e := range r.GetEvents() {
			e.Time = nil
		}
	}
}

var (
	instancesR = regexp.MustCompile(`"(instances|results|changes|[a-z0-9_]*money)":\s*\[([^\]]+)\]`)
	spaceR     = regexp.MustCompile(`\s+`)
	oneofR     = regexp.MustCompile(`"Oneof[^"]+":\{`)
)

// Remove Oneofs for readability
func cleanOneof(src []byte) []byte {
	pairs := oneofR.FindAllIndex(src, -1)
	for _, pair := range pairs {
		x := 0
		for i := pair[0]; i < pair[1]; i++ {
			src[i] = ' '
		}
		for i := pair[1]; i < len(src); i++ {
			if src[i] == '{' {
				x++
			}
			if src[i] == '}' {
				if x == 0 {
					src[i] = ' '
					break
				}
				x--
			}
		}
	}
	return src
}

func formatJson(x []byte) ([]byte, error) {
	return json.MarshalIndent(json.RawMessage(x), "", " ")
}
