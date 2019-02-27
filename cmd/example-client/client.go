package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	proto_log "github.com/dnovikoff/mahjong-api/genproto/log"
	"github.com/dnovikoff/mahjong-api/pkg/robots"
)

func main() {
	address := flag.String("address", ":9090", "address to connect")
	insecure := flag.Bool("insecure", true, "TLS disabled")
	logname := flag.String("log-name", "example", "Prefix for log files")
	seed := flag.Int64("seed", 0, "Redefine seed for game")
	rule := flag.String("rule", "EMA", "Rule ID")
	flag.Parse()
	var opts []grpc.DialOption
	if *insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	gc, err := grpc.Dial(*address, opts...)
	check(err)
	defer gc.Close()
	c := proto_api.NewGameClient(gc)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	req := &proto_api.CreateRequest{
		OneofRules: &proto_api.CreateRequest_RuleId{
			RuleId: *rule,
		},
		Creator: &proto_api.CreatorInfo{
			Name:  "Your Name",
			Email: "your@email",
		},
		Players: []*proto_api.Player{
			&proto_api.Player{Caption: "Player", PlayerType: proto_api.PlayerType_CLIENT},
			&proto_api.Player{Caption: "robot1", PlayerType: proto_api.PlayerType_ROBOT_EFFECTIVE},
			&proto_api.Player{Caption: "robot2", PlayerType: proto_api.PlayerType_ROBOT_TSUMOGIRI},
			&proto_api.Player{Caption: "robot3", PlayerType: proto_api.PlayerType_ROBOT_EFFECTIVE},
		},
		Timeouts: nil,
	}
	if *seed != 0 {
		req.Seed = &wrappers.Int64Value{
			Value: *seed,
		}
	}
	resp, err := c.Create(ctx, req)
	check(err)
	token := resp.GetPlayers()[0].GetToken()
	robot := robots.NewSettingsDecorator(robots.NewEffective())
	robot.Settings.AutoWin = true
	check(robots.Run(ctx, gc, token, robot))
	log, err := c.GetStorageLog(ctx, &proto_api.GetLogRequest{LogId: resp.GameId})
	check(err)
	check(saveStorageLog(*logname, log))
}

func saveStorageLog(name string, sl *proto_api.StorageLog) error {
	sl = proto.Clone(sl).(*proto_api.StorageLog)
	cleanStorage(sl)
	if err := saveJson(name+"_create", sl.Create); err != nil {
		return err
	}
	if err := saveJson(name+"_log", sl.Log); err != nil {
		return err
	}
	for k, v := range sl.PlayerMessages {
		if err := saveJson(fmt.Sprintf("%v_%v.debug", name, k+1), v); err != nil {
			return err
		}
	}
	return nil
}

// Remove timestamps from logs for readability
func cleanStorage(sl *proto_api.StorageLog) {
	cleanLog(sl.Log)
	for _, v := range sl.GetPlayerMessages() {
		for _, m := range v.Messages {
			m.Time = nil
		}
	}
}

// Remove timestamps from logs for readability
func cleanLog(log *proto_log.Log) {
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

func saveJson(name string, x interface{}) error {
	bytes, err := json.Marshal(x)
	if err != nil {
		return err
	}
	bytes = cleanOneof(bytes)
	bytes, err = formatJson(bytes)
	if err != nil {
		return err
	}
	bytes = instancesR.ReplaceAllFunc(bytes, func(src []byte) []byte {
		src = spaceR.ReplaceAll(src, []byte(` `))
		src = []byte(strings.TrimSpace(string(src)))
		return src

	})
	return ioutil.WriteFile(name+".json", bytes, 0644)
}

func formatJson(x []byte) ([]byte, error) {
	return json.MarshalIndent(json.RawMessage(x), "", " ")
}

func check(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}
