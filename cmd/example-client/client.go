package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"

	private_game "github.com/dnovikoff/mahjong-api/genproto/private/game"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/dnovikoff/mahjong-api/pkg/bootstrap"
	"github.com/dnovikoff/mahjong-api/pkg/logs"
	"github.com/dnovikoff/mahjong-api/pkg/robots"
)

type Config struct {
	PublicGame  *bootstrap.ClientConfig `yaml:"public-game"`
	PrivateGame *bootstrap.ClientConfig `yaml:"private-game"`
	PublicLogs  *bootstrap.ClientConfig `yaml:"public-logs"`
	PrivateLogs *bootstrap.ClientConfig `yaml:"private-logs"`
	LogName     string                  `yaml:"log-name"`
}

var (
	config = flag.String("config", "", "path to config file")
	seed   = flag.Int64("seed", 0, "Redefine seed for game")
	rule   = flag.String("rule", "EMA", "Rule ID")
)

func main() {
	flag.Parse()
	var cfg Config
	check(bootstrap.DecodeConfigFile(*config, &cfg))
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	c, err := cfg.PrivateGame.Dial()
	check(err)
	defer c.Close()
	resp := create(ctx, c)
	token := resp.GetPlayers()[0].GetToken()
	robot := robots.NewSettingsDecorator(robots.NewEffective())
	robot.Settings.AutoWin = true
	rc, err := cfg.PublicGame.Dial()
	check(err)
	defer rc.Close()
	check(robots.Run(ctx, rc, token, robot))
	lc, err := cfg.PublicLogs.Dial()
	check(err)
	defer lc.Close()
	time.Sleep(time.Second)
	log := getLog(ctx, lc, resp.GameId)
	check(logs.SaveJSONLogs(cfg.LogName+"_", log, nil))
}

func getLog(ctx context.Context, gc *grpc.ClientConn, id string) *public_log.Log {
	pl := public_log.NewLogServiceClient(gc)
	resp, err := pl.GetLog(ctx, &public_log.GetLogRequest{LogId: id})
	check(err)
	return resp.GetLog()
}

func create(ctx context.Context, gc *grpc.ClientConn) *private_game.CreateResponse {
	pg := private_game.NewGameServiceClient(gc)
	req := &private_game.CreateRequest{
		OneofRules: &private_game.CreateRequest_RuleId{
			RuleId: *rule,
		},
		Players: []*private_game.Player{
			{Caption: "Player", PlayerType: private_game.PlayerType_CLIENT},
			{Caption: "robot1", PlayerType: private_game.PlayerType_ROBOT_EFFECTIVE},
			{Caption: "robot2", PlayerType: private_game.PlayerType_ROBOT_TSUMOGIRI},
			{Caption: "robot3", PlayerType: private_game.PlayerType_ROBOT_EFFECTIVE},
		},
		Timeouts: nil,
	}
	if *seed != 0 {
		req.Seed = &wrappers.Int64Value{
			Value: *seed,
		}
	}
	resp, err := pg.Create(ctx, req)
	check(err)
	return resp
}

func check(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}
